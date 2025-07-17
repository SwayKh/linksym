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
	var err error
	aliasSourcePath := config.AliasPath(paths.SourcePath, paths.HomeDir, paths.InitDir, true)
	aliasDestinationPath := config.AliasPath(paths.DestinationPath, paths.HomeDir, paths.InitDir, true)

	// If path is a directory, Rename it
	if paths.IsDirectory {
		// Delete destination, if it exists
		err = DeleteFile(paths.DestinationPath)
		if err != nil {
			return err
		}

		err = os.Rename(paths.SourcePath, paths.DestinationPath)
		if err != nil {
			return fmt.Errorf("couldn't link directory %s to %s: %w", aliasSourcePath, aliasDestinationPath, err)
		}
		logger.Log(logger.INFO, "Moving: %s to %s", aliasSourcePath, aliasDestinationPath)
	} else {
		err = moveFile(paths.SourcePath, paths.DestinationPath, paths.HomeDir, paths.InitDir)
		if err != nil {
			return err
		}
	}

	err = paths.Link()
	if err != nil {
		return err
	}
	return nil
}

// Create a symlink of source path at the destination path,
func (paths LinkPaths) Link() error {
	var err error
	aliasDestinationPath := config.AliasPath(paths.DestinationPath, paths.HomeDir, paths.InitDir, true)

	// Create paths for source path, if it doesn't exist
	err = os.MkdirAll(filepath.Dir(paths.SourcePath), 0o755)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(paths.SourcePath), err)
	}

	err = os.Symlink(paths.DestinationPath, paths.SourcePath)
	if err != nil {
		return fmt.Errorf("couldn't create symlink %s: %w", aliasDestinationPath, err)
	}

	logger.Log(logger.SUCCESS, "Creating symlink for %s", filepath.Base(aliasDestinationPath))
	return nil
}

// Remove the symlink file at the source, move the destination file to the
// original source path. Basically undo-ing the MoveAndLink function
func (paths LinkPaths) UnLink() error {
	aliasSourcePath := config.AliasPath(paths.SourcePath, paths.HomeDir, paths.InitDir, true)
	aliasDestinationPath := config.AliasPath(paths.DestinationPath, paths.HomeDir, paths.InitDir, true)

	// Delete destination, if it exists
	err := DeleteFile(paths.SourcePath)
	if err != nil {
		return err
	}

	if paths.IsDirectory {
		err := os.Rename(paths.DestinationPath, paths.SourcePath)
		if err != nil {
			return fmt.Errorf("couldn't move directory %s to %s: %w", aliasSourcePath, aliasDestinationPath, err)
		}
		logger.Log(logger.INFO, "Moving: %s to %s", aliasSourcePath, aliasDestinationPath)
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

	logger.Log(logger.INFO, "Moving: %s to %s", aliasSourcePath, aliasDestinationPath)

	src, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("failed to open file: %s: %w", source, err)
	}
	defer src.Close()

	err = os.MkdirAll(filepath.Dir(destination), 0o755)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(destination), err)
	}

	dst, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", destination, err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("failed to copy file %s to %s: %w", source, destination, err)
	}

	err = DeleteFile(source)
	if err != nil {
		return err
	}
	return nil
}

// Delete the file at the given path
func DeleteFile(path string) error {
	file, err := config.GetFileInfo(path)
	if err != nil {
		return err
	}

	if !file.Exists {
		return nil
	}

	err = os.RemoveAll(path)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("failed to Remove file %s. Please run with elevated privileges", path)
		} else {
			return fmt.Errorf("failed to Remove file: %w", err)
		}
	}
	return nil
}
