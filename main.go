package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	configPath, err := config.Initialise()
	if err != nil {
		// fmt.Errorf("Error initialising config: %s", err)
		fmt.Println(err)
	}

	cfg, err := config.LoadConfig("./.linksym.yaml")
	if err != nil {
		// fmt.Errorf("Error loading config: %s", err)
		fmt.Println(err)
	}

	sourcePath := os.Args[1]
	filename := filepath.Base(sourcePath)
	// destinationPath := os.Args[2]
	destinationPath := cfg.InitDirectory + "/" + filename

	fmt.Println(sourcePath, filename, destinationPath)

	err = linker.Link(sourcePath, destinationPath, configPath)
	if err != nil {
		fmt.Println(err)
	}
}
