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
	record := record{}

	recordSlice := []string{}
	recordSlice = append(recordSlice, sourcePath, destinationPath)

	filename := filepath.Base(destinationPath)
	dirname := filepath.Base(filepath.Dir(destinationPath))

	fileAndDirName := filepath.Join(dirname, filename)

	record.Name = fileAndDirName
	record.Paths = recordSlice

	c.Records = append(c.Records, record)

	logger.VerboseLog(logger.INFO, "Adding record to .linksym.yaml...")
}

// Remove a Record of Link Name and Path array from the AppConfig struct, which
// is written to file at the end of program execution
func (c *AppConfig) RemoveRecord(name string) {
	for i := len(c.Records) - 1; i >= 0; i-- {
		if c.Records[i].Name == name {
			c.Records = append(c.Records[:i], c.Records[i+1:]...)
		}
	}

	logger.VerboseLog(logger.INFO, "Removing record from .linksym.yaml...")
}

func (c *AppConfig) AliasConfig(homeDir, initDir string) {
	c.InitDirectory = AliasPath(c.InitDirectory, homeDir, initDir, true)

	// Alias path absolute paths before writing to config file
	for i, v := range c.Records {
		for j := range v.Paths {
			c.Records[i].Paths[j] = AliasPath(c.Records[i].Paths[j], homeDir, initDir, false)
		}
	}
}

func (c *AppConfig) UnAliasConfig(homeDir, initDir string) {
	c.InitDirectory = ExpandPath(c.InitDirectory, homeDir, initDir)

	for i, v := range c.Records {
		for j := range v.Paths {
			c.Records[i].Paths[j] = ExpandPath(c.Records[i].Paths[j], homeDir, initDir)
		}
	}
}

func InitialiseHomePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Couldn't get the home directory")
	}
	return homeDir, nil
}
