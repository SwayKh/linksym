package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig() error {
	// Check if config file exists
	if fileExists, _, err := CheckFile(ConfigPath); err != nil {
		return fmt.Errorf("Error checking if .linksym.yaml exists: \n%w", err)
	} else if !fileExists {
		return fmt.Errorf("No .linksym.yaml file found, please run linksym init: \n%w", err)
	}

	file, err := os.Open(ConfigPath)
	if err != nil {
		return fmt.Errorf("Error opening config file: %s: \n%w", ConfigPath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Error reading data from config file: \n%w", err)
	}

	err = yaml.Unmarshal(data, &Configuration)
	if err != nil {
		return fmt.Errorf("Error loading data to appConfig{}: \n%w", err)
	}

	for i := range Configuration.Records {
		Configuration.Records[i].Paths = expandPath(Configuration.Records[i].Paths)
	}
	return nil
}

func WriteConfig() error {
	// Alias path absolute paths before writing to config file
	for i := range Configuration.Records {
		Configuration.Records[i].Paths = aliasPath(Configuration.Records[i].Paths)
	}

	data, err := yaml.Marshal(&Configuration)
	if err != nil {
		return fmt.Errorf("Error marshalling data from appConfig{}: \n%w", err)
	}

	err = os.WriteFile(ConfigPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing record to config file: \n%w", err)
	}
	return nil
}
