package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/SwayKh/linksym/pkg/config"
)

func Run() error {
	var err error

	CreateFlags()
	flag.Parse()

	// Since the Init Command creates the config file, the LoadConfig function
	// can't be called before handling the init subcommand.
	if flag.Arg(0) == "init" {
		return Init()
	}

	err = config.SetupDirectories()
	if err != nil {
		return err
	}

	err = config.LoadConfig()
	if err != nil {
		return err
	}

	if *HelpFlag {
		Help()
		os.Exit(0)
	}

	switch flag.Arg(0) {
	case "":
		Help()
	case "init":
		break
	case "add":
		return Add(os.Args[2:])
	case "remove":
		return Remove(os.Args[2:])
	case "source":
		return Source()
	default:
		return fmt.Errorf("Invalid Command. Please use -h or --help flags to see available commands.")
	}

	if err := config.WriteConfig(); err != nil {
		return err
	}
	return nil
}
