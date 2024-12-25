package commands

import (
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/logger"
)

// Loop over the configuration []Records, for each entry get the source and
// destination paths. Run the Link command for each entry.
func (app *Application) Source() error {
	logger.VerboseLog(logger.INFO, "Creating Symlinks from .linksym.yaml Records...")
	for _, record := range app.Configuration.Records {
		sourcePath := record.Paths[0]
		destinationPath := record.Paths[1]

		err := os.MkdirAll(filepath.Dir(destinationPath), 0o755)
		if err != nil {
			return err
		}

		pathArgs := []string{sourcePath, destinationPath}

		// Don't stop the program wehn encountering a error when looping over
		// records for source command. Move on to linking next records.
		err = app.Add(pathArgs, true, false)
		if err != nil {
			logger.Log(logger.ERROR, err.Error())
		}
	}
	logger.Log(logger.SUCCESS, "Success")
	return nil
}
