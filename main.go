package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SwayKh/linksym/cmd"
	"github.com/SwayKh/linksym/pkg/config"
)

func main() {
	err := config.SetupDirectories()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
		err = cmd.Init()
	case "add":
		err = cmd.Add(args[1:])
	case "remove":
		err = cmd.Remove(args[1])
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
}
