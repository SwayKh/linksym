package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/global"
	"github.com/SwayKh/linksym/pkg/logger"
	"github.com/SwayKh/linksym/pkg/utils"
	"gopkg.in/yaml.v3"
)

// Load the configuration from .linksym.yaml configuration file and unmarshall
// it into the global Configuration variable, and return pointer to this struct
func LoadConfig(configPath string) (*AppConfig, error) {
	// Check if config file exists
	logger.VerboseLog("Checking if config file exists...")

	config, err := utils.GetFileInfo(configPath)
	if err != nil {
		return nil, fmt.Errorf("Error getting File Info of %s: %w", configPath, err)
	} else if !config.Exists {
		return nil, fmt.Errorf("No .linksym.yaml file found. Please run linksym init.")
	}

	file, err := os.Open(config.AbsPath)
	if err != nil {
		return nil, fmt.Errorf("Error opening config file: %s ", filepath.Base(configPath))
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading data from config file: %w", err)
	}

	logger.VerboseLog("Getting data from config file...")
	configuration := &AppConfig{}

	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return nil, fmt.Errorf("Error getting data from config file: %w", err)
	}

	return configuration, nil
}

// Write the Configuration struct data to .linksym.yaml file
func WriteConfig(configuration *AppConfig) error {
	AliasConfig(configuration)
	data, err := yaml.Marshal(configuration)
	if err != nil {
		return fmt.Errorf("Error marshalling data from configuration{}: %w", err)
	}

	logger.VerboseLog("Updating config file...")

	err = os.WriteFile(global.ConfigPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing record to config file: %w", err)
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

	InitDirectory = utils.AliasPath(InitDirectory, true)

	configuration := AppConfig{
		InitDirectory: InitDirectory,
		Records:       []record{},
	}

	AliasConfig(&configuration)
	data, err := yaml.Marshal(configuration)
	if err != nil {
		return fmt.Errorf("Error marshalling data from configuration{}: %w", err)
	}

	err = os.WriteFile(configPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing record to config file: %w", err)
	}

	logger.Log("Creating an empty %s file in the current directory.\n", filepath.Base(configPath))

	return nil
}
