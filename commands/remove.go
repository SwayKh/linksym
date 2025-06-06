package commands

import (
	"fmt"
	"path/filepath"

	"github.com/SwayKh/linksym/config"
	"github.com/SwayKh/linksym/link"
	"github.com/SwayKh/linksym/logger"
)

// Get the absolute path of "LinkName", which should be the path relative from
// the Init Directory, and Find the matching LinkName through the []Records in
// .linksym.yaml, UnLink it, and Remove from the []Records. can take multiple
// arguments and looks over them Removing each one
func (app *Application) Remove(args []string) error {
	for _, path := range args {
		var sourcePath, destinationPath string
		var err error
		var found bool

		// Get the File Info of LinkName provided from the arguments
		linkInfo, err := config.GetFileInfo(path)
		if err != nil {
			return err
		} else if !linkInfo.Exists {
			return fmt.Errorf("file %s doesn't exist", linkInfo.AbsPath)
		}

		logger.Log(logger.WARNING, "Unlinking %s", config.AliasPath(linkInfo.AbsPath, app.HomeDirectory, app.InitDirectory, true))

		// Since the "filename" of the record can be the same with a different file,
		// just linked in separate directories, getting the filename and the parent
		// directory name should make the name unique enough to be checked in
		// []Records
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
			return fmt.Errorf("no record found for %s", config.AliasPath(linkInfo.AbsPath, app.HomeDirectory, app.InitDirectory, true))
		}

		paths := link.LinkPaths{
			SourcePath:      sourcePath,
			DestinationPath: destinationPath,
			HomeDir:         app.HomeDirectory,
			InitDir:         app.InitDirectory,
			IsDirectory:     linkInfo.IsDir,
		}

		err = paths.UnLink()
		if err != nil {
			return err
		}

		app.Configuration.RemoveRecord(recordPathName)

		// Save the config after removing each records, since if out of multiple
		// arguments provided and one of them is not present in the records or
		// returns some argument, the .linksym.yaml won't be saved with the removed
		// records that were removed.
		if err := app.Configuration.WriteConfig(app.HomeDirectory, app.InitDirectory, app.ConfigPath); err != nil {
			return err
		}
	}
	return nil
}
