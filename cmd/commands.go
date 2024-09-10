package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
)

// Initialise and empty config with cwd as init directory
func Init() {
	err := config.InitialiseConfig()
	if err != nil {
		fmt.Printf("Error initialising config: %s\n", err)
		os.Exit(1)
	}
}

func Add(args []string) {
	var sourcePath string
	var destinationPath string

	if len(args) == 1 {
		sourcePath = args[0]
		filename := filepath.Base(sourcePath)
		destinationPath = filepath.Join(config.CurrentWorkingDirectory, filename)
	} else if len(args) == 2 {
		sourcePath = args[0]
		destinationPath = args[1]
	} else {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}

	fmt.Println(sourcePath, destinationPath)

	err := linker.Link(sourcePath, destinationPath)
	if err != nil {
		fmt.Println(err)
	}
}

func Remove() {
	fmt.Println(config.HomeDirectory, config.ConfigPath, config.CurrentWorkingDirectory)
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
	os.Exit(1)
}
