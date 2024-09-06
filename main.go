package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	homeDirectory   string
	configDirectory string
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
		return err
	}

	configDirectory, err = os.UserConfigDir()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := setupDirectories()
	if err != nil {
		fmt.Println(err)
	}
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	// flag.Parse()
	//
	// if *help || *h {
	// 	fmt.Println(usage)
	// 	os.Exit(1)
	// }

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	path := os.Args[1]
	if _, err = os.Stat(path); os.IsNotExist(err) {
		fmt.Println("File Doesn't exist", err)
		os.Exit(1)
	}
	filename := filepath.Base(path)

	fmt.Println(homeDirectory, configDirectory, cwd, path, filename)
	err = os.Symlink(path, cwd+"/"+filename)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
