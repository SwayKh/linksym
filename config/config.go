package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/logger"
)

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
	logger.VerboseLog("Adding record to .linksym.yaml...")
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
	logger.VerboseLog("Removing record from .linksym.yaml...")
	for i := len(c.Records) - 1; i >= 0; i-- {
		if c.Records[i].Name == name {
			c.Records = append(c.Records[:i], c.Records[i+1:]...)
		}
	}
}

// Un-Alias all paths with ~ and $init_directory with full absolute paths
func (configuration *AppConfig) UnAliasConfig(homeDir, initDir string) {
	configuration.InitDirectory = ExpandPath(configuration.InitDirectory, homeDir, initDir)

	for i, v := range configuration.Records {
		for j := range v.Paths {
			configuration.Records[i].Paths[j] = ExpandPath(configuration.Records[i].Paths[j], homeDir, initDir)
		}
	}
}

// Alias all mentions of HomeDirectory and InitDirectory with ~ and $init_directory
func (configuration *AppConfig) AliasConfig(homeDir, initDir string) {
	configuration.InitDirectory = AliasPath(configuration.InitDirectory, homeDir, initDir, true)

	// Alias path absolute paths before writing to config file
	for i, v := range configuration.Records {
		for j := range v.Paths {
			configuration.Records[i].Paths[j] = AliasPath(configuration.Records[i].Paths[j], homeDir, initDir, false)
		}
	}
}

// Set the global HomeDirectory variable. Separated from SetupDirectories to be
// used with the Init Subcommand.
func InitialiseHomePath() (string, error) {
	HomeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Couldn't get the home directory")
	}
	return HomeDirectory, nil
}
