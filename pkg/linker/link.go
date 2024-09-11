package linker

import (
	"fmt"
	"io"
	"os"

	"github.com/SwayKh/linksym/pkg/config"
)

func Link(sourcePath string, destinationPath string, isDirectory bool, toMove bool) error {
	// If path is a directory, Rename it
	if isDirectory {
		err := os.Rename(sourcePath, destinationPath)
		if err != nil {
			return fmt.Errorf("Couldn't rename directory %s to %s\n %w", sourcePath, destinationPath, err)
		}
	} else if toMove {
		// If Path is a file, copy it to new path, and delete original
		err := moveFile(sourcePath, destinationPath)
		if err != nil {
			return err
		}
	}
	err := os.Symlink(destinationPath, sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't create symlink %s\n %w", destinationPath, err)
	}

	err = config.AddRecord(sourcePath, destinationPath)
	if err != nil {
		return err
	}
	return nil
}

func moveFile(source, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("Failed to source file: %s\n %w", source, err)
	}
	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("Failed to create file %s\n %w", destination, err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("Failed to copy file %s to %s\n %w", source, destination, err)
	}
	err = dst.Sync()
	if err != nil {
		return fmt.Errorf("Failed to write file %s to disk: %w", destination, err)
	}

	err = os.Remove(source)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("Failed to Remove file %s\n Please run with elevated privileges\n", source)
		} else {
			return fmt.Errorf("Failed to Remove file %s\n %w", source, err)
		}
	}
	return nil
}
