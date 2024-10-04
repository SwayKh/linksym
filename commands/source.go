package commands

import (
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/logger"
)

// Loop over the configuration []Records, for each entry get the source and
// destination paths. Run the Link command for each entry.
func (app *Application) Source() error {
	logger.VerboseLog("Creating Symlinks from .linksym.yaml Records...")
	for _, record := range app.Configuration.Records {
		sourcePath := record.Paths[0]
		destinationPath := filepath.Dir(record.Paths[1])

		err := os.MkdirAll(destinationPath, 0o755)
		if err != nil {
			return err
		}

		pathArgs := []string{sourcePath, destinationPath}

		err = app.Add(pathArgs, false)
		if err != nil {
			return err
		}
	}
	return nil
}
