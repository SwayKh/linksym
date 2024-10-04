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
			return fmt.Errorf("File %s doesn't exist", linkInfo.AbsPath)
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

		aliasLinkPath := config.AliasPath(linkInfo.AbsPath, app.HomeDirectory, app.InitDirectory, true)
		aliasSourcePath := config.AliasPath(sourcePath, app.HomeDirectory, app.InitDirectory, true)
		aliasDestinationPath := config.AliasPath(destinationPath, app.HomeDirectory, app.InitDirectory, true)

		if !found {
			return fmt.Errorf("No record found for %s", aliasLinkPath)
		}

		logger.Log(logger.INFO, "Moving: %s to %s", aliasSourcePath, aliasDestinationPath)
		err = link.UnLink(sourcePath, destinationPath, linkInfo.IsDir)
		if err != nil {
			return err
		}

		app.Configuration.RemoveRecord(recordPathName)
	}
	return nil
}
