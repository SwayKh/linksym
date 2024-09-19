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
		return fmt.Errorf("Error checking if .linksym.yaml exists: \n %w", err)
	} else if !fileExists {
		return fmt.Errorf("No .linksym.yaml file found, please run linksym init: \n %w", err)
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

	Configuration.InitDirectory = expandPath(Configuration.InitDirectory)

	for i, v := range Configuration.Records {
		for j := range v.Paths {
			Configuration.Records[i].Paths[j] = expandPath(Configuration.Records[i].Paths[j])
		}
	}

	return nil
}

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
		return fmt.Errorf("Error marshalling data from appConfig{}: \n%w", err)
	}

	err = os.WriteFile(ConfigPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing record to config file: \n%w", err)
	}
	return nil
}

func InitialiseConfig() error {
	err := SetupDirectories()
	if err != nil {
		return fmt.Errorf("Initialising Env: \n%w", err)
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
		return fmt.Errorf("Error marshalling Init Data from appConfig{}: \n%w", err)
	}

	err = os.WriteFile(ConfigPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing data to config file: \n%w", err)
	}
	return nil
}
