package config

import (
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/utils"
)

// This package should, create the default config with the initialised directory
// Read the config for every other command to work and use the variable vales
// Add data under the "Links" variable whenever a new config is add via the
// project.
// The config file can be either .json .ini .toml .yaml
// I think yaml is a good file format for this

type AppConfig struct {
	InitDirectory string   `yaml:"init_directory"`
	Records       []record `yaml:"records"`
}

type record struct {
	Name  string   `yaml:"name"`
	Paths []string `yaml:"paths"`
}

// Create a array of Path provided and a Link Name which is appended in the
// Records of the global Configuration Struct
func (c *AppConfig) AddRecord(sourcePath string, destinationPath string) {
	record := record{}

	recordSlice := []string{}
	recordSlice = append(recordSlice, sourcePath, destinationPath)

	filename := filepath.Base(destinationPath)
	dirname := filepath.Base(filepath.Dir(destinationPath))

	fileAndDirName := filepath.Join(dirname, filename)

	record.Name = fileAndDirName
	record.Paths = recordSlice

	c.Records = append(c.Records, record)
}

// Remove a Record of Link Name and Path array from the global configuration
// struct, which is written to file at the end of program execution
func (c *AppConfig) RemoveRecord(name string) {
	for i := len(c.Records) - 1; i >= 0; i-- {
		if c.Records[i].Name == name {
			c.Records = append(c.Records[:i], c.Records[i+1:]...)
		}
	}
}

func UnAliasConfig(configuration *AppConfig) {
	configuration.InitDirectory = utils.ExpandPath(configuration.InitDirectory)

	for i, v := range configuration.Records {
		for j := range v.Paths {
			configuration.Records[i].Paths[j] = utils.ExpandPath(configuration.Records[i].Paths[j])
		}
	}
}

func AliasConfig(configuration *AppConfig) {
	configuration.InitDirectory = utils.AliasPath(configuration.InitDirectory, true)

	// Alias path absolute paths before writing to config file
	for i, v := range configuration.Records {
		for j := range v.Paths {
			configuration.Records[i].Paths[j] = utils.AliasPath(configuration.Records[i].Paths[j], false)
		}
	}
}
