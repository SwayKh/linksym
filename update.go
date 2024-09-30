package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Update the $init_directory variable in the config and update the record names
// based on the new destination path with $init_directory
func Update(configuration *AppConfig) error {
	// Alias config, to be allow expanding the $init_directory variable with the
	// new InitDirectory
	AliasConfig(configuration)

	VerboseLog("Updating .linksym.yaml file...")

	InitDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Couldn't get the current working directory")
	}

	SetupDirectories(InitDirectory, filepath.Base(ConfigPath))
	configuration.InitDirectory = InitDirectory

	for i := range configuration.Records {
		destinationPath := ExpandPath(configuration.Records[i].Paths[1])
		filename := filepath.Base(destinationPath)
		dirname := filepath.Base(filepath.Dir(destinationPath))

		configuration.Records[i].Name = filepath.Join(dirname, filename)

	}

	err = WriteConfig(configuration)
	if err != nil {
		return nil
	}

	return nil
}
