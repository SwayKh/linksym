package commands

import (
	"fmt"
	"path/filepath"

	"github.com/SwayKh/linksym/config"
	"github.com/SwayKh/linksym/logger"
)

func (app *Application) Record(args []string) error {
	if len(args) == 1 {
		source, err := config.GetFileInfo(args[0])
		if err != nil {
			return err
		}

		if !source.Exists {
			return fmt.Errorf("file %s doesn't exist", source.AbsPath)
		}

		sourcePath := source.AbsPath
		filename := filepath.Base(sourcePath)
		destinationPath := filepath.Join(app.InitDirectory, filename)

		aliasSourcePath := config.AliasPath(sourcePath, app.HomeDirectory, app.InitDirectory, true)
		aliasDestinationPath := config.AliasPath(destinationPath, app.HomeDirectory, app.InitDirectory, true)

		logger.VerboseLog(logger.SUCCESS, "Source path exists: %s", aliasSourcePath)
		logger.VerboseLog(logger.SUCCESS, "Destination path exists: %s", aliasDestinationPath)

		// Check if sourcePath is a symlink to destinationPath
		isLink, err := checkSymlink(sourcePath, destinationPath)
		if err != nil {
			return err
		}

		if isLink {
			logger.Log(logger.WARNING, "Symlink already exists")
			// Add the record if source exists, and it's a symlink
			app.Configuration.AddRecord(sourcePath, destinationPath)
		} else {
			logger.VerboseLog(logger.WARNING, "Source and destination are both files.")
			logger.VerboseLog(logger.WARNING, "Source file is not a symlink to destination file")
			logger.VerboseLog(logger.WARNING, "Please use another commands to handle creating a symlink")
			logger.Log(logger.ERROR, "Unable to create a record in .linksym.yaml")
		}
		return nil
	} else if len(args) == 2 {

		source, err := config.GetFileInfo(args[0])
		if err != nil {
			return err
		}

		destination, err := config.GetFileInfo(args[1])
		if err != nil {
			return err
		}

		sourcePath := source.AbsPath
		destinationPath := destination.AbsPath

		aliasSourcePath := config.AliasPath(source.AbsPath, app.HomeDirectory, app.InitDirectory, true)
		aliasDestinationPath := config.AliasPath(destination.AbsPath, app.HomeDirectory, app.InitDirectory, true)

		logger.VerboseLog(logger.SUCCESS, "Source path: %s", aliasSourcePath)
		logger.VerboseLog(logger.SUCCESS, "Destination path: %s", aliasDestinationPath)

		// Check if sourcePath is a symlink to destinationPath
		isLink, err := checkSymlink(sourcePath, destinationPath)
		if err != nil {
			return err
		}

		if isLink {
			logger.Log(logger.WARNING, "Symlink already exists")
			// Add the record if source exists, and it's a symlink
			app.Configuration.AddRecord(sourcePath, destinationPath)
		} else {
			logger.VerboseLog(logger.WARNING, "Source and destination are both files.")
			logger.VerboseLog(logger.WARNING, "Source file is not a symlink to destination file")
			logger.VerboseLog(logger.WARNING, "Please use another commands to handle creating a symlink")
			logger.Log(logger.ERROR, "Unable to create a record in .linksym.yaml")
		}
		return nil
	}
	return nil
}
