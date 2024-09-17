package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
)

// Initialise and empty config with cwd as init directory
func Init() error {
	err := config.InitialiseConfig()
	if err != nil {
		return err
	}
	return nil
}

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
		destinationPath = filepath.Join(config.InitDirectory, filename)

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

		if destinationFileExists && DestinationFileInfo.IsDir() {
			filename := filepath.Base(sourcePath)
			destinationPath = filepath.Join(destinationPath, filename)
			isDirectory = true
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

func Remove(args []string) error {
	switch len(args) {
	case 1:
		linkName := args[0]
		var linkPath string
		var sourcePath, destinationPath string
		var err error
		var isDirectory bool

		linkPath, err = filepath.Abs(linkName)
		if err != nil {
			return fmt.Errorf("Error getting absolute path of file %s: \n%w", linkPath, err)
		}

		fileExists, fileInfo, err := config.CheckFile(linkPath)
		if err != nil {
			return err
		} else if !fileExists {
			return fmt.Errorf("File %s doesn't exist", linkPath)
		} else if fileInfo.IsDir() {
			isDirectory = true
		}

		recordPathName := filepath.Join(filepath.Base(filepath.Dir(linkPath)), filepath.Base(linkPath))

		for i := range config.Configuration.Records {
			if config.Configuration.Records[i].Name == recordPathName {
				sourcePath = config.Configuration.Records[i].Paths[0]
				destinationPath = config.Configuration.Records[i].Paths[1]
				err = config.RemoveRecord(i)
				if err != nil {
					return nil
				}
			}
		}

		err = linker.UnLink(sourcePath, destinationPath, isDirectory)
		if err != nil {
			return nil
		}

	default:
		return fmt.Errorf("Invalid number of arguments")
	}
	return nil
}

func Source() error {
	return nil
}

func Help() {
	fmt.Println("Usage: linksym [subcommand] [flags]")

	fmt.Println("\n Subcommands:")
	fmt.Println("   add [Path] [(optional) Destination]:")
	fmt.Println("     Create a symlink for given path, optionally define a destination for symlink")
	fmt.Println("   remove [Path]")
	fmt.Println("     Remove the symlink and move the file to the original path")

	fmt.Println("\n Flags:")
	fmt.Println("   -h, --help")
	fmt.Println("     Print this help message")
	os.Exit(0)
}
