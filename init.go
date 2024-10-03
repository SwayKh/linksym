package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Initialise and empty config with cwd as init directory
func Init(configName string) error {
	err := InitialiseConfig(configName)
	if err != nil {
		return err
	}
	return nil
}

// Create a default config file with empty records and Current working directory
// variable for Init directory
func InitialiseConfig(configPath string) error {
	InitDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Couldn't get the current working directory")
	}

	InitDirectory = AliasPath(InitDirectory, true)

	configuration := AppConfig{
		InitDirectory: InitDirectory,
		Records:       []record{},
	}

	configuration.AliasConfig()
	data, err := yaml.Marshal(configuration)
	if err != nil {
		return fmt.Errorf("Error marshalling data from configuration{}: %w", err)
	}

	err = os.WriteFile(configPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing record to config file: %w", err)
	}

	Log("Initialising %s file in the current directory.", filepath.Base(configPath))

	return nil
}
