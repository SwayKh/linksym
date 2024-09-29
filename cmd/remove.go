package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
	"github.com/SwayKh/linksym/pkg/utils"
)

// Get the absolute path of "LinkName", which should be the path relative from
// the Init Directory, and Find the matching LinkName through the []Records in
// .linksym.yaml, UnLink it, and Remove from the []Records. Expects one
// argument, and throws error on multiple arguments provided
func Remove(configuration *config.AppConfig, args []string) error {
	switch len(args) {
	case 1:
		var sourcePath, destinationPath string
		var err error
		var found bool

		// Get the File Info of LinkName provided from the arguments
		link, err := utils.GetFileInfo(args[0])
		if err != nil {
			return err
		} else if !link.Exists {
			return fmt.Errorf("File %s doesn't exist", link.AbsPath)
		}

		// Since the "filename" of the record can be the same with a different file,
		// just linked in separate directories, getting the filename and the above
		// directory name should make the name unique enough to be checked in
		// []Records
		recordPathName := filepath.Join(filepath.Base(filepath.Dir(link.AbsPath)), filepath.Base(link.AbsPath))

		// Can't use range over Configuration.Records. since the slice gets modified
		// during iteration with the RemoveRecord function, so iterate over the
		// slice manually in reverse order to fix this
		for i := len(configuration.Records) - 1; i >= 0; i-- {
			if configuration.Records[i].Name == recordPathName {
				sourcePath = configuration.Records[i].Paths[0]
				destinationPath = configuration.Records[i].Paths[1]
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("No record found for %s", utils.AliasPath(link.AbsPath, true))
		}

		err = linker.UnLink(sourcePath, destinationPath, link.IsDir)
		if err != nil {
			return err
		}

		configuration.RemoveRecord(recordPathName)

	default:
		return fmt.Errorf("Invalid number of arguments")
	}
	return nil
}
