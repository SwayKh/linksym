package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// This package should, create the default config with the initialised directory
// Read the config for every other command to work and use the variable vales
// Add data under the "Links" variable whenever a new config is add via the
// project.
// The config file can be either .json .ini .toml .yaml
// I think yaml is a good file format for this

var (
	HomeDirectory string
	ConfigPath    string
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

func SetupDirectories() error {
	var err error
	HomeDirectory, err = os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Couldn't get the home directory")
	}

	ConfigName := ".linksym.yaml"
	ConfigPath = filepath.Join(Configuration.InitDirectory, ConfigName)

	return nil
}

// Create a array of Path provided and a Link Name which is appended in the
// Records of the global Configuration Struct
func AddRecord(sourcePath, destinationPath string) error {
	record := record{}

	recordSlice := []string{}
	recordSlice = append(recordSlice, sourcePath, destinationPath)

	filename := filepath.Base(destinationPath)
	dirname := filepath.Base(filepath.Dir(destinationPath))

	fileAndDirName := filepath.Join(dirname, filename)

	record.Name = fileAndDirName
	record.Paths = recordSlice

	Configuration.Records = append(Configuration.Records, record)

	return nil
}

// Remove a Record of Link Name and Path array from the global configuration
// struct, which is written to file at the end of program execution
func RemoveRecord(i int) {
	Configuration.Records = append(Configuration.Records[:i], Configuration.Records[i+1:]...)
}
