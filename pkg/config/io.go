package config

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func CheckFile(path string) (bool, os.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil, nil
		} else {
			return false, nil, fmt.Errorf("Error getting file info: %w", err)
		}
	}
	return true, fileInfo, nil
}

func loadConfig(configPath string) (appConfig, error) {
	// Check if config file exists
	if fileExists, _, err := CheckFile(configPath); err != nil {
		return appConfig{}, fmt.Errorf("Error checking if .linksym.yaml exists: %w", err)
	} else if !fileExists {
		return appConfig{}, fmt.Errorf("No .linksym.yaml file found, please run linksym init: %w", err)
	}

	file, err := os.Open(configPath)
	if err != nil {
		return appConfig{}, fmt.Errorf("Error opening config file: %s: %w", configPath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return appConfig{}, fmt.Errorf("Error reading data from config file: %w", err)
	}

	config := appConfig{}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return appConfig{}, fmt.Errorf("Error loading data to appConfig{}: %w", err)
	}
	return config, nil
}

func writeConfig(configuration appConfig) error {
	data, err := yaml.Marshal(&configuration)
	if err != nil {
		return fmt.Errorf("Error marshalling data from appConfig{}: %w", err)
	}

	err = os.WriteFile(ConfigPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing record to config file: %w", err)
	}
	return nil
}
