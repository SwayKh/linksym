package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/global"
	"github.com/SwayKh/linksym/pkg/utils"
)

// Update the $init_directory variable in the config and update the record names
// based on the new destination path with $init_directory
func Update(configuration *config.AppConfig) error {
	// Alias config, to be allow expanding the $init_directory variable with the
	// new InitDirectory
	config.AliasConfig(configuration)

	InitDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Couldn't get the current working directory")
	}

	utils.SetupDirectories(InitDirectory, global.ConfigPath)
	configuration.InitDirectory = InitDirectory

	for i := range configuration.Records {
		destinationPath := utils.ExpandPath(configuration.Records[i].Paths[1])
		filename := filepath.Base(destinationPath)
		dirname := filepath.Base(filepath.Dir(destinationPath))

		configuration.Records[i].Name = filepath.Join(dirname, filename)

	}

	err = config.WriteConfig(configuration, global.ConfigPath)
	if err != nil {
		return nil
	}

	return nil
}
