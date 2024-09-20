package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// Load the configuration from .linksym.yaml configuration file and unmarshall
// it into the global Configuration variable, and un-alias all paths
func LoadConfig(configPath string) (*AppConfig, error) {
	// Check if config file exists
	if fileExists, _, err := CheckFile(configPath); err != nil {
		return nil, fmt.Errorf("Error checking if .linksym.yaml exists: %w", err)
	} else if !fileExists {
		return nil, fmt.Errorf("No .linksym.yaml file found. Please run linksym init.")
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("Error opening config file: %s ", configPath)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading data from config file: %w", err)
	}

	configuration := &AppConfig{}

	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return nil, fmt.Errorf("Error loading data to appConfig{}: %w", err)
	}

	configuration.InitDirectory = expandPath(configuration.InitDirectory)

	for i, v := range configuration.Records {
		for j := range v.Paths {
			configuration.Records[i].Paths[j] = expandPath(configuration.Records[i].Paths[j])
		}
	}

	return configuration, nil
}

// Write the Configuration struct data to .linksym.yaml file after aliasing all
// paths with ~ and $init_directory
func WriteConfig() error {
	Configuration.InitDirectory = aliasPath(Configuration.InitDirectory, true)

	// Alias path absolute paths before writing to config file
	for i, v := range Configuration.Records {
		for j := range v.Paths {
			Configuration.Records[i].Paths[j] = aliasPath(Configuration.Records[i].Paths[j], false)
		}
	}

	data, err := yaml.Marshal(&Configuration)
	if err != nil {
		return fmt.Errorf("Error marshalling data from Configuration{}: %w", err)
	}

	err = os.WriteFile(ConfigPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing record to config file: %w", err)
	}
	return nil
}

// Create a default config file with empty records and Current working directory
// variable for Init directory
func InitialiseConfig() error {
	err := SetupDirectories()
	if err != nil {
		return fmt.Errorf("Initialising Env: %w", err)
	}

	InitDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Couldn't get the current working directory")
	}

	InitDirectory = aliasPath(InitDirectory, true)

	configuration := appConfig{
		InitDirectory: InitDirectory,
		Records:       []record{},
	}

	data, err := yaml.Marshal(&configuration)
	if err != nil {
		return fmt.Errorf("Error marshalling Init Data from Configuration{}: %w", err)
	}

	err = os.WriteFile(ConfigPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing data to config file: %w", err)
	}
	return nil
}
