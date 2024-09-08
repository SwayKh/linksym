package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/linker"
)

var (
	homeDirectory           string
	configDirectory         string
	currentWorkingDirectory string
)

// var (
// 	add    = flag.Bool("add", false, "add a symlink to given path")
// 	remove = flag.Bool("remove", false, "Remove a symlink")
// 	help   = flag.Bool("--help", false, "Print this help message")
// 	h      = flag.Bool("-h", false, "Print this help message")
// )

var usage string = `
Usage: linksym [option...] [path...]

Options: 
	add                   add a symlink to given path
	remove                Remove a symlink
	-h                    Print this help message
`

func setupDirectories() error {
	var err error
	homeDirectory, err = os.UserHomeDir()
	if err != nil {
		return errors.New("Couldn't get the home directory")
	}

	configDirectory, err = os.UserConfigDir()
	if err != nil {
		return errors.New("Couldn't get the config directory")
	}

	currentWorkingDirectory, err = os.Getwd()
	if err != nil {
		return errors.New("Couldn't get the current working directory")
	}

	return nil
}

func main() {
	// Get Home, config and current directory
	err := setupDirectories()
	if err != nil {
		fmt.Println(err)
	}

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	path := os.Args[1]
	filename := filepath.Base(path)
	newFilePath := currentWorkingDirectory + "/" + filename

	err = linker.Link(path, newFilePath)
	if err != nil {
		fmt.Println(err)
	}
}

// Create a init function, that create the config files, stores the working
// directory, and other stuff, every other command needs to check if the config
// file exists, before it works.
// The config package will be separates, that adds and reads config, the init
// function should probably call that package

func initialise() error {
	return nil
}
