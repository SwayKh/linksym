package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Expand the ~ and $init_directory variables to their respective values
func ExpandPath(path string) string {
	// the $init_directory strings comes from the yaml tags for AppConfig
	if strings.HasPrefix(path, "$init_directory") {
		path = strings.Replace(path, "$init_directory", InitDirectory, 1)
	}
	if strings.HasPrefix(path, "~") {
		path = strings.Replace(path, "~", HomeDirectory, 1)
	}
	return path
}

// Create aliases of ~ and $init_directory to make the paths and the
// configurations more portable
func AliasPath(path string, skipInitDir bool) string {
	// the $init_directory strings comes from the yaml tags for AppConfig
	if !skipInitDir && strings.HasPrefix(path, InitDirectory) {
		path = strings.Replace(path, InitDirectory, "$init_directory", 1)
	}
	if strings.HasPrefix(path, HomeDirectory) {
		path = strings.Replace(path, HomeDirectory, "~", 1)
	}

	return path
}

// Set values to global variables of InitDirectory and ConfigPath
func SetupDirectories(initDir string, configName string) {
	InitDirectory = ExpandPath(initDir)
	ConfigPath = filepath.Join(InitDirectory, configName)
}

// Set the global HomeDirectory variable. Separated from SetupDirectories to be
// used with the Init Subcommand.
func InitialiseHomePath() error {
	var err error
	HomeDirectory, err = os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Couldn't get the home directory")
	}
	return nil
}
