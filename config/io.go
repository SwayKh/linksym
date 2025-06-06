package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/logger"
	"gopkg.in/yaml.v3"
)

// Load the configuration from .linksym.yaml configuration file and unmarshall
// it into the AppConfig struct, and return pointer to this struct
func LoadConfig(configPath string) (*AppConfig, error) {
	// Check if config file exists
	logger.VerboseLog(logger.INFO, "Checking if config file exists...")

	config, err := GetFileInfo(configPath)
	if err != nil {
		return nil, fmt.Errorf("error getting File Info of %s: %w", configPath, err)
	} else if !config.Exists {
		return nil, fmt.Errorf("no .linksym.yaml file found. Please run linksym init")
	}

	file, err := os.Open(config.AbsPath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %s ", filepath.Base(configPath))
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading data from config file: %w", err)
	}

	logger.VerboseLog(logger.INFO, "Getting data from config file...")
	configuration := &AppConfig{}

	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return nil, fmt.Errorf("error getting data from config file: %w", err)
	}

	return configuration, nil
}

// Write the Configuration struct data to .linksym.yaml file
func (configuration *AppConfig) WriteConfig(homeDir, initDir, configPath string) error {
	data, err := yaml.Marshal(configuration)
	if err != nil {
		return fmt.Errorf("error marshalling data from configuration{}: %w", err)
	}

	logger.VerboseLog(logger.SUCCESS, "Updating config file...")

	err = os.WriteFile(configPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("error writing record to config file: %w", err)
	}
	return nil
}
