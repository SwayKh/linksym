package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/config"
	"github.com/SwayKh/linksym/logger"
)

// Update the $init_directory variable in the config and update the record names
// based on the new destination path with $init_directory
func (app *Application) Update() error {
	// Alias config, to be allow expanding the $init_directory variable with the
	// new InitDirectory
	app.Configuration.AliasConfig(app.HomeDirectory, app.InitDirectory)

	logger.Log("Updating .linksym.yaml file...")

	InitDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Couldn't get the current working directory")
	}

	app.InitDirectory = config.ExpandPath(InitDirectory, app.HomeDirectory, InitDirectory)
	app.ConfigPath = filepath.Join(app.InitDirectory, app.ConfigName)
	app.Configuration.InitDirectory = InitDirectory

	for i := range app.Configuration.Records {
		destinationPath := config.ExpandPath(app.Configuration.Records[i].Paths[1], app.HomeDirectory, app.InitDirectory)
		filename := filepath.Base(destinationPath)
		dirname := filepath.Base(filepath.Dir(destinationPath))

		app.Configuration.Records[i].Name = filepath.Join(dirname, filename)

	}

	err = app.Configuration.WriteConfig(app.HomeDirectory, app.InitDirectory, app.ConfigPath)
	if err != nil {
		return nil
	}

	return nil
}
