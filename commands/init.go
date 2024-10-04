package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/config"
	"github.com/SwayKh/linksym/logger"
	"gopkg.in/yaml.v3"
)

// Initialise and empty config with cwd as init directory
func (app *Application) Init() error {
	err := initialiseConfig(app.ConfigName, app.HomeDirectory)
	if err != nil {
		return err
	}
	return nil
}

// Create a default config file with empty records and Current working directory
// variable for Init directory
func initialiseConfig(configPath, homeDir string) error {
	initDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Couldn't get the current working directory")
	}

	initDirectory = config.AliasPath(initDirectory, homeDir, initDirectory, true)

	configuration := config.AppConfig{}
	configuration.InitDirectory = initDirectory

	configuration.AliasConfig(homeDir, initDirectory)
	data, err := yaml.Marshal(configuration)
	if err != nil {
		return fmt.Errorf("Error marshalling data from configuration{}: %w", err)
	}

	err = os.WriteFile(configPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing record to config file: %w", err)
	}

	logger.Log("Initialising %s file in the current directory.", filepath.Base(configPath))

	return nil
}
