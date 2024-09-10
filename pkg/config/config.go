package config

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// This package should, create the default config with the initialised directory
// Read the config for every other command to work and use the variable vales
// Add data under the "Links" variable whenever a new config is add via the
// project.
// The config file can be either .json .ini .toml .yaml
// I think yaml is a good file format for this

type Config struct {
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

// Create a init function, that create the config files, stores the working
// directory, and other stuff, every other command needs to check if the config
// file exists, before it works.
// The config package will be separates, that adds and reads config, the init
// function should probably call that package

func InitialiseConfig() error {
	err := setupDirectories()
	if err != nil {
		return err
	}

	cfg := Config{
		InitDirectory: CurrentWorkingDirectory,
		Record:        [][]string{},
	}

	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigPath, data, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig(configPath string) (*Config, error) {
	// Check if config file exists
	if fileExists, _, err := CheckFile(configPath); err != nil {
		return nil, err
	} else if !fileExists {
		return nil, errors.New("Config file doesn't exist")
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %s, %w", configPath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config := Config{}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func AddRecord(sourcePath, destinationPath, configPath string) error {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return err
	}

	record := []string{}
	record = append(record, sourcePath, destinationPath)

	cfg.Record = append(cfg.Record, record)

	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return fmt.Errorf("Error marshalling data from cofnig struct\n %w", err)
	}

	err = os.WriteFile(configPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing record to config file\n %w", err)
	}

	return nil
}
