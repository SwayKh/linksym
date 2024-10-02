package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SwayKh/linksym/flags"
)

type App struct {
	Configuration *AppConfig
	HomeDirectory string
	InitDirectory string
	ConfigPath    string
}

var (
	HomeDirectory string
	ConfigPath    string
	InitDirectory string
)

func main() {
	if err := Run(); err != nil {
		Log("Error: %v", err)
		os.Exit(1)
	}
}

// Load config, Setup up Global variables and handle all subcommand switching
func Run() error {
	flags.CreateFlags()
	flag.Parse()

	configName := ".linksym.yaml"

	err := InitialiseHomePath()
	if err != nil {
		return err
	}

	subcommand := flag.Arg(0)

	if len(flag.Args()) < 1 {
		Help()
		os.Exit(1)
	}

	args := flag.Args()[1:]

	// Since the Init Command creates the config file, the LoadConfig function
	// can't be called before handling the init subcommand.
	// But Init function calls aliasPath, which requires HomeDirectory variable,
	// and InitialiseHomePath needs be called before this.
	if subcommand == "init" {
		if len(args) > 0 {
			return fmt.Errorf("'init' subcommand doesn't accept any arguments.\nUsage: linksym init")
		}
		return Init(configName)
	}

	if *flags.HelpFlag {
		Help()
		os.Exit(0)
	}

	configuration, err := LoadConfig(configName)
	if err != nil {
		return err
	}

	SetupDirectories(configuration.InitDirectory, configName)
	UnAliasConfig(configuration)

	switch subcommand {
	case "init":
		break
	case "add":
		if len(args) > 2 {
			return fmt.Errorf("'add' subcommand doesn't accept more than 2 arguments.\nUsage: linksym add <source> <destination>")
		}
		err = Add(configuration, args, true)
	case "remove":
		if len(args) > 1 {
			return fmt.Errorf("'remove' subcommand doesn't accept more than 1 argument.\nUsage: linksym remove <file name>")
		}
		err = Remove(configuration, args)
	case "source":
		if len(args) > 0 {
			return fmt.Errorf("'source' subcommand doesn't accept any arguments.\nUsage: linksym source")
		}
		err = Source(configuration)
	case "update":
		if len(args) > 0 {
			return fmt.Errorf("'update subcommand doesn't accept any arguments.\nUsage: linksym update")
		}
		err = Update(configuration)

	default:
		err = fmt.Errorf("Invalid Command. Please use -h or --help flags to see available commands.")
	}

	if err != nil {
		return err
	}

	if err := WriteConfig(configuration); err != nil {
		return err
	}
	return nil
}
