package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SwayKh/linksym/cmd"
	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/utils"
)

func main() {
	if err := Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

// Load config, Setup up Global variables and handle all subcommand switching
func Run() error {
	cmd.CreateFlags()
	flag.Parse()

	configName := ".linksym.yaml"

	err := utils.InitialiseHomePath()
	if err != nil {
		return err
	}

	subcommand := flag.Arg(0)

	if len(flag.Args()) < 1 {
		cmd.Help()
		os.Exit(1)
	}

	args := flag.Args()[1:]

	// Since the Init Command creates the config file, the LoadConfig function
	// can't be called before handling the init subcommand.
	// But Init function calls aliasPath, which requires HomeDirectory variable,
	// and InitialiseHomePath needs be called before this.
	if subcommand == "init" {
		if err = cmd.InitFlag.Parse(args); err != nil {
			return err
		}
		if remainingArgs := cmd.InitFlag.Args(); len(remainingArgs) > 0 {
			return fmt.Errorf("'init' subcommand doesn't accept any arguments.\nUsage: linksym init [-u, --update]")
		}
		if !cmd.UpdateInitBool {
			return cmd.Init(configName)
		}
	}

	if *cmd.HelpFlag {
		cmd.Help()
		os.Exit(0)
	}

	configuration, err := config.LoadConfig(configName)
	if err != nil {
		return err
	}

	utils.SetupDirectories(configuration.InitDirectory, configName)
	config.UnAliasConfig(configuration)

	switch subcommand {
	case "init":
		if cmd.UpdateInitBool {
			return cmd.UpdateInit(configuration, configName)
		}
	case "add":
		if len(args) > 2 {
			return fmt.Errorf("'add' subcommand doesn't accept more than 2 arguments.\nUsage: linksym add <source> <destination>")
		}
		err = cmd.Add(configuration, args, true)
	case "remove":
		if len(args) > 1 {
			return fmt.Errorf("'remove' subcommand doesn't accept more than 1 argument.\nUsage: linksym remove <file name>")
		}
		err = cmd.Remove(configuration, args)
	case "source":
		if len(args) > 0 {
			return fmt.Errorf("'source' subcommand doesn't accept any arguments.\nUsage: linksym source")
		}
		err = cmd.Source(configuration)
	default:
		err = fmt.Errorf("Invalid Command. Please use -h or --help flags to see available commands.")
	}

	if err != nil {
		return err
	}

	config.AliasConfig(configuration)
	if err := config.WriteConfig(configuration, configName); err != nil {
		return err
	}
	return nil
}
