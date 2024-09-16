package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(configPath string) (appConfig, error) {
	// Check if config file exists
	if fileExists, _, err := CheckFile(configPath); err != nil {
		return appConfig{}, fmt.Errorf("Error checking if .linksym.yaml exists: \n%w", err)
	} else if !fileExists {
		return appConfig{}, fmt.Errorf("No .linksym.yaml file found, please run linksym init: \n%w", err)
	}

	file, err := os.Open(configPath)
	if err != nil {
		return appConfig{}, fmt.Errorf("Error opening config file: %s: \n%w", configPath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return appConfig{}, fmt.Errorf("Error reading data from config file: \n%w", err)
	}

	config := appConfig{}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return appConfig{}, fmt.Errorf("Error loading data to appConfig{}: \n%w", err)
	}

	for i := range config.Records {
		config.Records[i].Paths = expandPath(config.Records[i].Paths)
	}
	return config, nil
}

func writeConfig(configuration appConfig) error {
	data, err := yaml.Marshal(&configuration)
	if err != nil {
		return fmt.Errorf("Error marshalling data from appConfig{}: \n%w", err)
	}

	err = os.WriteFile(ConfigPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing record to config file: \n%w", err)
	}
	return nil
}
