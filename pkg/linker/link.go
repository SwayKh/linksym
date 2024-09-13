package linker

import (
	"fmt"
	"io"
	"os"

	"github.com/SwayKh/linksym/pkg/config"
)

func MoveAndLink(sourcePath, destinationPath string, isDirectory bool) error {
	// If path is a directory, Rename it
	if isDirectory {
		err := os.Rename(sourcePath, destinationPath)
		if err != nil {
			return fmt.Errorf("Couldn't rename directory %s to %s: %w", sourcePath, destinationPath, err)
		}
	} else {
		// If path is a file, create a file at new location, copy it over, and
		// delete original file. This method allows better handling when linking
		// across file system than just renaming files
		err := moveFile(sourcePath, destinationPath)
		if err != nil {
			return err
		}
	}

	err := Link(sourcePath, destinationPath)
	if err != nil {
		return err
	}
	return nil
}

func Link(sourcePath, destinationPath string) error {
	err := os.Symlink(destinationPath, sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't create symlink %s: %w", destinationPath, err)
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
		return fmt.Errorf("Failed to source file: %s: %w", source, err)
	}
	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("Failed to create file %s: %w", destination, err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("Failed to copy file %s to %s: %w", source, destination, err)
	}
	err = dst.Sync()
	if err != nil {
		return fmt.Errorf("Failed to write file %s to disk: %w", destination, err)
	}

	err = os.Remove(source)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("Failed to Remove file %s\n Please run with elevated privileges", source)
		} else {
			return fmt.Errorf("Failed to Remove file %s: %w", source, err)
		}
	}
	return nil
}
