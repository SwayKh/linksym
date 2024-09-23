package linker

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Move the source file to destination and creates a symlink at the source
// pointing towards the destination path
func MoveAndLink(sourcePath, destinationPath string, isDirectory bool) error {
	// If path is a directory, Rename it
	if isDirectory {
		err := os.Rename(sourcePath, destinationPath)
		if err != nil {
			return fmt.Errorf("Couldn't link directory %s to %s: %w", sourcePath, destinationPath, err)
		}
	} else {
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

// Create a symlink of source path at the destination path,
func Link(sourcePath, destinationPath string) error {
	err := os.Symlink(destinationPath, sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't create symlink %s: %w", destinationPath, err)
	}
	return nil
}

// Remove the symlink file at the source, move the destination file to the
// original source path. Basically undo-ing the MoveAndLink function
func UnLink(sourcePath, destinationPath string, isDirectory bool) error {
	err := deleteFile(sourcePath)
	if err != nil {
		return err
	}

	if isDirectory {
		err := os.Rename(destinationPath, sourcePath)
		if err != nil {
			return fmt.Errorf("Couldn't move directory %s to %s: %w", sourcePath, destinationPath, err)
		}
	} else {
		err := moveFile(destinationPath, sourcePath)
		if err != nil {
			return err
		}
	}
	return nil
}

// Create a a file at the destination, copy all contents of the source to the
// destination and then remove the source. This method allows better handling
// when linking across file system than just renaming files
func moveFile(source, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("Failed to open file: %s: %w", source, err)
	}
	defer src.Close()

	err = os.MkdirAll(filepath.Dir(destination), 0o755)
	if err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", filepath.Dir(destination), err)
	}

	dst, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("Failed to create file %s: %w", destination, err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("Failed to copy file %s to %s: %w", source, destination, err)
	}

	err = deleteFile(source)
	if err != nil {
		return err
	}
	return nil
}

// Delete the file at the given path
func deleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("Failed to Remove file %s. Please run with elevated privileges", path)
		} else {
			return fmt.Errorf("Failed to Remove file %s: %w", path, err)
		}
	}
	return nil
}
