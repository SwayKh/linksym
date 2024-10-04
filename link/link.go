package link

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/config"
	"github.com/SwayKh/linksym/logger"
)

type LinkPaths struct {
	SourcePath      string
	DestinationPath string
	HomeDir         string
	InitDir         string
	IsDirectory     bool
}

// Move the source file to destination and creates a symlink at the source
// pointing towards the destination path
func (paths LinkPaths) MoveAndLink() error {
	// If path is a directory, Rename it
	if paths.IsDirectory {
		err := os.Rename(paths.SourcePath, paths.DestinationPath)
		if err != nil {
			return fmt.Errorf("Couldn't link directory %s to %s: %w", paths.SourcePath, paths.DestinationPath, err)
		}
	} else {
		err := moveFile(paths.SourcePath, paths.DestinationPath, paths.HomeDir, paths.InitDir)
		if err != nil {
			return err
		}
	}

	err := paths.Link()
	if err != nil {
		return err
	}
	return nil
}

// Create a symlink of source path at the destination path,
func (paths LinkPaths) Link() error {
	err := os.Symlink(paths.DestinationPath, paths.SourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't create symlink %s: %w", paths.DestinationPath, err)
	}

	logger.Log(logger.SUCCESS, "Creating symlink...")
	return nil
}

// Remove the symlink file at the source, move the destination file to the
// original source path. Basically undo-ing the MoveAndLink function
func (paths LinkPaths) UnLink() error {
	err := deleteFile(paths.SourcePath)
	if err != nil {
		return err
	}

	if paths.IsDirectory {
		err := os.Rename(paths.DestinationPath, paths.SourcePath)
		if err != nil {
			return fmt.Errorf("Couldn't move directory %s to %s: %w", paths.SourcePath, paths.DestinationPath, err)
		}
	} else {
		err := moveFile(paths.DestinationPath, paths.SourcePath, paths.HomeDir, paths.InitDir)
		if err != nil {
			return err
		}
	}
	return nil
}

// Create a a file at the destination, copy all contents of the source to the
// destination and then remove the source. This method allows better handling
// when linking across file system than just renaming files
func moveFile(source, destination, homeDir, initDir string) error {
	aliasSourcePath := config.AliasPath(source, homeDir, initDir, true)
	aliasDestinationPath := config.AliasPath(destination, homeDir, initDir, true)

	logger.Log(logger.INFO, "Moving: %s to %s\n", aliasSourcePath, aliasDestinationPath)

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
	file, err := config.GetFileInfo(path)
	if err != nil {
		return err
	}

	if !file.Exists {
		return nil
	}

	err = os.Remove(path)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("Failed to Remove file %s. Please run with elevated privileges", path)
		} else {
			return fmt.Errorf("Failed to Remove file: %w", err)
		}
	}
	return nil
}
