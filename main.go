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

	// Since the Init Command creates the config file, the LoadConfig function
	// can't be called before handling the init subcommand.
	// But Init function calls aliasPath, which requires HomeDirectory variable,
	// and hence function SetupDirectories was split up
	if flag.Arg(0) == "init" {
		return cmd.Init(configName)
	}

	if *cmd.HelpFlag {
		cmd.Help()
		os.Exit(0)
	}

	configuration, err := config.LoadConfig(configName)
	if err != nil {
		return err
	}

	// Some issues here. UnAliasPath requires global.InitDirectory and which is
	// unaliasing the config, then the InitDirectory, ConfigPath will have ~ as
	// setup by SetupDirectories(), But if setup SetupDirectories is run before
	// alias for user home directory.
	// Solved, call the ExpandPath function on the InitDir inside SetupDirectories
	utils.SetupDirectories(configuration.InitDirectory, configName)
	config.UnAliasConfig(configuration)

	switch flag.Arg(0) {
	case "":
		cmd.Help()
	case "init":
		break
	case "add":
		err = cmd.Add(configuration, os.Args[2:])
	case "remove":
		err = cmd.Remove(configuration, os.Args[2:])
	case "source":
		err = cmd.Source()
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
