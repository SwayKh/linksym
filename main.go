package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SwayKh/linksym/cmd"
	"github.com/SwayKh/linksym/pkg/config"
)

func main() {
	var err error
	cmd.CreateFlags()
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		cmd.Help()
	}

	if *cmd.HelpFlag {
		cmd.Help()
	}

	if args[0] == "init" {
		err = cmd.Init()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	err = config.SetupDirectories()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	subcommand := args[0]

	// This mess is created since the Config needs to be loaded at startup, which
	// check if the configuration file is present or not, which doesn't won't
	// exist until the init function is called, so the LoadConfig function needs
	// to be called after a "init" subcommand call
	switch subcommand {
	case "init":
		break
	case "add":
		err = cmd.Add(args[1:])
	case "remove":
		err = cmd.Remove(args[1:])
	case "source":
		err = cmd.Source()
	default:
		fmt.Println("Unknown subcommand")
		cmd.Help()
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err := config.WriteConfig(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
