package main

import (
	"flag"
	"fmt"

	"github.com/SwayKh/linksym/cmd"
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
		cmd.Add(args[1:])
	case "remove":
		cmd.Remove()
	default:
		fmt.Println("Unknown subcommand")
		cmd.Help()
	}
}
