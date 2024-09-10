package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// This package should, create the default config with the initialised directory
// Read the config for every other command to work and use the variable vales
// Add data under the "Links" variable whenever a new config is add via the
// project.
// The config file can be either .json .ini .toml .yaml
// I think yaml is a good file format for this

type appConfig struct {
	InitDirectory string     `yaml:"init_directory"`
	Record        [][]string `yaml:"record"`
}

var (
	HomeDirectory           string
	CurrentWorkingDirectory string
	ConfigPath              string
)

func setupDirectories() error {
	var err error
	HomeDirectory, err = os.UserHomeDir()
	if err != nil {
		return errors.New("Couldn't get the home directory")
	}

	CurrentWorkingDirectory, err = os.Getwd()
	if err != nil {
		return errors.New("Couldn't get the current working directory")
	}

	ConfigPath = CurrentWorkingDirectory + "/.linksym.yaml"
	return nil
}

func InitialiseConfig() error {
	err := setupDirectories()
	if err != nil {
		return fmt.Errorf("Initialising Env: %w", err)
	}

	configuration := appConfig{
		InitDirectory: CurrentWorkingDirectory,
		Record:        [][]string{},
	}

	data, err := yaml.Marshal(&configuration)
	if err != nil {
		return fmt.Errorf("Error marshalling Init Data from appConfig{}: %w", err)
	}

	err = os.WriteFile(ConfigPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing data to config file: %w", err)
	}
	return nil
}

func AddRecord(sourcePath, destinationPath string) error {
	configuration, err := loadConfig(ConfigPath)
	if err != nil {
		return err
	}

	recordSlice := []string{}
	recordSlice = append(recordSlice, sourcePath, destinationPath)
	configuration.Record = append(configuration.Record, recordSlice)

	if err := writeConfig(configuration); err != nil {
		return err
	}
	return nil
}
