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
		return fmt.Errorf("Error initialising config: %v\n", err)
	}
	return nil
}

func Add(args []string) error {
	var sourcePath string
	var destinationPath string

	if len(args) == 1 {
		sourcePath = args[0]
		filename := filepath.Base(sourcePath)
		destinationPath = filepath.Join(config.CurrentWorkingDirectory, filename)
	} else if len(args) == 2 {
		sourcePath = args[0]
		destinationPath = args[1]

		fileExists, fileInfo, err := config.CheckFile(destinationPath)

		if err != nil {
			return err
		} else if fileInfo.IsDir() {
			filename := filepath.Base(sourcePath)
			destinationPath = filepath.Join(destinationPath, filename)

			destinationPath, err = filepath.Abs(destinationPath)
			if err != nil {
				return fmt.Errorf("Error getting absolute path of file %s: %v\n", destinationPath, err)
			}
		} else if fileExists {
			// Need to cover special case of linking a already existing config
			return fmt.Errorf("File %s already exists", destinationPath)
		}
	} else {
		return fmt.Errorf("Invalid number of arguments")
	}

	err := linker.Link(sourcePath, destinationPath)
	if err != nil {
		return err
	}
	return nil
}

func Remove() error {
	fmt.Println(config.HomeDirectory, config.ConfigPath, config.CurrentWorkingDirectory)
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
