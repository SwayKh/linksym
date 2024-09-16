package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// This package should, create the default config with the initialised directory
// Read the config for every other command to work and use the variable vales
// Add data under the "Links" variable whenever a new config is add via the
// project.
// The config file can be either .json .ini .toml .yaml
// I think yaml is a good file format for this

var (
	HomeDirectory string
	InitDirectory string
	ConfigPath    string
	ConfigName    string
	Configuration appConfig
)

type appConfig struct {
	InitDirectory string   `yaml:"init_directory"`
	Records       []record `yaml:"records"`
}

type record struct {
	Name  string   `yaml:"name"`
	Paths []string `yaml:"paths"`
}

func InitialiseConfig() error {
	err := SetupDirectories()
	if err != nil {
		return fmt.Errorf("Initialising Env: \n%w", err)
	}

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

func AddRecord(sourcePath, destinationPath string) error {
	record := record{}

	recordSlice := []string{}
	recordSlice = append(recordSlice, sourcePath, destinationPath)

	filename := filepath.Base(destinationPath)
	dirname := filepath.Base(filepath.Dir(destinationPath))

	fileAndDirName := filepath.Join(dirname, filename)

	record.Name = fileAndDirName
	record.Paths = recordSlice

	// record.Paths = aliasPath(record.Paths)
	Configuration.Records = append(Configuration.Records, record)

	for i := range Configuration.Records {
		Configuration.Records[i].Paths = aliasPath(Configuration.Records[i].Paths)
	}

	if err := WriteConfig(); err != nil {
		return err
	}
	return nil
}

func RemoveRecord(index int) error {
	Configuration.Records = removeElement(Configuration.Records, index)

	if err := WriteConfig(); err != nil {
		return err
	}
	return nil
}
