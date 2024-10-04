package commands

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/config"
	"github.com/SwayKh/linksym/flags"
)

type Application struct {
	Configuration *config.AppConfig
	ConfigName    string
	ConfigPath    string
	HomeDirectory string
	InitDirectory string
}

func (app *Application) Run() error {
	flags.CreateFlags()
	flag.Parse()

	if len(flag.Args()) < 1 {
		Help()
		os.Exit(1)
	}

	if *flags.HelpFlag {
		Help()
		os.Exit(0)
	}

	subcommand := flag.Arg(0)
	args := flag.Args()[1:]

	// Since the Init Command creates the config file, the LoadConfig function
	// can't be called before handling the init subcommand.
	// But Init function calls aliasPath, which requires HomeDirectory variable,
	// and InitialiseHomePath needs be called before this.
	if subcommand == "init" {
		if len(args) > 0 {
			return fmt.Errorf("'init' subcommand doesn't accept any arguments.\nUsage: linksym init")
		}
		return app.Init()
	}

	configuration, err := config.LoadConfig(app.ConfigName)
	if err != nil {
		return err
	}

	app.Configuration = configuration
	app.InitDirectory = config.ExpandPath(configuration.InitDirectory, app.HomeDirectory, configuration.InitDirectory)
	app.ConfigPath = filepath.Join(app.InitDirectory, app.ConfigName)

	app.Configuration.UnAliasConfig(app.HomeDirectory, app.InitDirectory)

	switch subcommand {
	case "init":
		break
	case "add":
		if len(args) > 2 {
			return fmt.Errorf("'add' subcommand doesn't accept more than 2 arguments.\nUsage: linksym add <source> <destination>")
		}
		err = app.Add(args, true)
	case "remove":
		err = app.Remove(args)
	case "source":
		if len(args) > 0 {
			return fmt.Errorf("'source' subcommand doesn't accept any arguments.\nUsage: linksym source")
		}
		err = app.Source()
	case "update":
		if len(args) > 0 {
			return fmt.Errorf("'update subcommand doesn't accept any arguments.\nUsage: linksym update")
		}
		err = app.Update()

	default:
		err = fmt.Errorf("Invalid Command. Please use -h or --help flags to see available commands.")
	}

	if err != nil {
		return err
	}

	if err := app.Configuration.WriteConfig(app.HomeDirectory, app.InitDirectory, app.ConfigPath); err != nil {
		return err
	}
	return nil
}
