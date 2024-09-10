package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/cmd"
	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
)

func main() {
	cmd.CreateFlags()
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		cmd.Help()
	}

	if *cmd.HelpFlag {
		cmd.Help()
	}

	subcommand := args[0]

	switch subcommand {
	case "init":
		cmd.Init()
	case "add":
		cmd.Add()
	case "remove":
		cmd.Remove()
	default:
		fmt.Println("Unknown subcommand")
		cmd.Help()
	}

	if fileExists, _, err := config.CheckFile("./.linksym.yaml"); err != nil {
		fmt.Println("Error checking if .linksym.yaml exists")
	} else if !fileExists {
		fmt.Println("No .linksym.yaml file found, please run linksym init")
		os.Exit(1)
	}

	err := config.InitialiseConfig()
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

	err = linker.Link(sourcePath, destinationPath, config.ConfigPath)
	if err != nil {
		fmt.Println(err)
	}
}
