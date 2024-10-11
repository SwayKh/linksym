package commands

import (
	"fmt"
	"path/filepath"

	"github.com/SwayKh/linksym/config"
	"github.com/SwayKh/linksym/link"
	"github.com/SwayKh/linksym/logger"
)

func (app *Application) Restore(args []string) error {
	for _, path := range args {
		var sourcePath, destinationPath string
		var err error
		var found bool

		// Get the File Info of LinkName provided from the arguments
		linkInfo, err := config.GetFileInfo(path)
		if err != nil {
			return err
		} else if !linkInfo.Exists {
			return fmt.Errorf("File %s doesn't exist", linkInfo.AbsPath)
		}

		fileName := filepath.Base(linkInfo.AbsPath)
		dirName := filepath.Base(filepath.Dir(linkInfo.AbsPath))
		recordPathName := filepath.Join(dirName, fileName)

		for i := range app.Configuration.Records {
			if app.Configuration.Records[i].Name == recordPathName {
				sourcePath = app.Configuration.Records[i].Paths[0]
				destinationPath = app.Configuration.Records[i].Paths[1]
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("No record found for %s", config.AliasPath(linkInfo.AbsPath, app.HomeDirectory, app.InitDirectory, true))
		}

		paths := link.LinkPaths{
			SourcePath:      sourcePath,
			DestinationPath: destinationPath,
			HomeDir:         app.HomeDirectory,
			InitDir:         app.InitDirectory,
			IsDirectory:     linkInfo.IsDir,
		}

		logger.Log(logger.WARNING, "Restoring %s", config.AliasPath(linkInfo.AbsPath, app.HomeDirectory, app.InitDirectory, true))
		err = paths.Link()
		if err != nil {
			return err
		}

	}
	return nil
}
