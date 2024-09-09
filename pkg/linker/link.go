package linker

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/SwayKh/linksym/pkg/config"
)

func Link(sourcePath, linkPath, configPath string) error {
	// Get File info, to check if it exists, and if it's a directory or not
	fileInfo, err := os.Stat(sourcePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("File path provided doesn't exist")
		} else {
			return err
		}
	}

	// If path is a directory, Rename ii
	if fileInfo.IsDir() {
		err = os.Rename(sourcePath, linkPath)
		if err != nil {
			return fmt.Errorf("Couldn't rename directory %s to %s\n %w", sourcePath, linkPath, err)
		}
	} else {
		// If Path is a file, copy it to new path, and delete original
		err = moveFile(sourcePath, linkPath)
		if err != nil {
			return err
		}
	}
	err = os.Symlink(linkPath, sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't create symlink %s\n %w", linkPath, err)
	}

	err = config.AddRecord(sourcePath, linkPath, configPath)
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
