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

		sourcePath, err = filepath.Abs(args[0])
		if err != nil {
			return fmt.Errorf("Error getting absolute path of file %s: \n%w", sourcePath, err)
		}

		fileExists, fileInfo, err := config.CheckFile(sourcePath)
		if err != nil {
			return err
		} else if !fileExists {
			return fmt.Errorf("File %s doesn't exist", sourcePath)
		} else if fileInfo.IsDir() {
			isDirectory = true
		}

		filename := filepath.Base(sourcePath)
		destinationPath = filepath.Join(config.Configuration.InitDirectory, filename)

	case 2:
		// set first and second args as source and destination path, get absolute
		// paths, check if the paths exist, plus handle the special case of source
		// path not existing but destination path exists, hence creating a link
		// without the moving the files

		sourcePath, err = filepath.Abs(args[0])
		if err != nil {
			return fmt.Errorf("Error getting absolute path of file %s: \n%w", sourcePath, err)
		}

		destinationPath, err = filepath.Abs(args[1])
		if err != nil {
			return fmt.Errorf("Error getting absolute path of file %s: \n%w", destinationPath, err)
		}

		sourceFileExists, sourceFileInfo, err := config.CheckFile(sourcePath)
		if err != nil {
			return err
		}

		destinationFileExists, DestinationFileInfo, err := config.CheckFile(destinationPath)
		if err != nil {
			return err
		}

		if sourceFileExists && sourceFileInfo.IsDir() && destinationFileExists && DestinationFileInfo.IsDir() {
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
