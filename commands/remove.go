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
// .linksym.yaml, UnLink it, and Remove from the []Records. Expects one
// argument, and throws error on multiple arguments provided
func (app *Application) Remove(args []string) error {
	switch len(args) {
	case 1:
		var sourcePath, destinationPath string
		var err error
		var found bool

		// Get the File Info of LinkName provided from the arguments
		linkInfo, err := config.GetFileInfo(args[0])
		if err != nil {
			return err
		} else if !linkInfo.Exists {
			return fmt.Errorf("File %s doesn't exist", linkInfo.AbsPath)
		}

		logger.Log("Unlinking %s", config.AliasPath(linkInfo.AbsPath, app.HomeDirectory, app.InitDirectory, true))

		// Since the "filename" of the record can be the same with a different file,
		// just linked in separate directories, getting the filename and the above
		// directory name should make the name unique enough to be checked in
		// []Records
		recordPathName := filepath.Join(filepath.Base(filepath.Dir(linkInfo.AbsPath)), filepath.Base(linkInfo.AbsPath))

		// Can't use range over Configuration.Records. since the slice gets modified
		// during iteration with the RemoveRecord function, so iterate over the
		// slice manually in reverse order to fix this
		for i := len(app.Configuration.Records) - 1; i >= 0; i-- {
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

		err = link.UnLink(sourcePath, destinationPath, linkInfo.IsDir)
		if err != nil {
			return err
		}

		app.Configuration.RemoveRecord(recordPathName)

	default:
		return fmt.Errorf("Invalid number of arguments")
	}
	return nil
}
