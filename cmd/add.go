package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
)

func Add(args []string) error {
	var sourcePath, destinationPath string
	var err error
	var isDirectory bool

	switch len(args) {

	case 1:
		// Set first arg source path, get absolute path, check if it exists, set the
		// destination path as cwd+filename of source path

		sourcePath, fileExists, fileInfo, err := filePathInfo(args[0])
		if err != nil {
			return err
		} else if !fileExists {
			return fmt.Errorf("File %s doesn't exist", sourcePath)
		} else if fileInfo.IsDir() {
			isDirectory = true
		}

		filename := filepath.Base(sourcePath)
		destinationPath = filepath.Join(config.Configuration.InitDirectory, filename)

		err = linker.MoveAndLink(sourcePath, destinationPath, isDirectory)
		if err != nil {
			return err
		}
		return nil

	case 2:
		// set first and second args as source and destination path, get absolute
		// paths, check if the paths exist, plus handle the special case of source
		// path not existing but destination path exists, hence creating a link
		// without the moving the files
		destinationPath, destinationFileExists, destinationFileInfo, err := filePathInfo(args[1])
		if err != nil {
			return err
		}

		sourcePath, sourceFileExists, sourceFileInfo, err := filePathInfo(args[0])
		if err != nil {
			return err
		}

		if sourceFileExists && sourceFileInfo.IsDir() && destinationFileExists && destinationFileInfo.IsDir() {
			filename := filepath.Base(sourcePath)
			destinationPath = filepath.Join(destinationPath, filename)
			isDirectory = true

			err = linker.MoveAndLink(sourcePath, destinationPath, isDirectory)
			if err != nil {
				return err
			}
			return nil
		}

		if sourceFileExists && sourceFileInfo.IsDir() && destinationFileExists {
			filename := filepath.Base(destinationPath)
			sourcePath = filepath.Join(sourcePath, filename)

			err := linker.Link(sourcePath, destinationPath)
			if err != nil {
				return err
			}
			return nil
		}

		if destinationFileExists && !sourceFileExists {
			err := linker.Link(sourcePath, destinationPath)
			if err != nil {
				return err
			}
			return nil
		}

	default:
		return fmt.Errorf("Invalid number of arguments")
	}

	err = linker.MoveAndLink(sourcePath, destinationPath, isDirectory)
	if err != nil {
		return err
	}
	return nil
}
