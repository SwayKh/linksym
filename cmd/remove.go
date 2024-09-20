package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
)

// Get the "LinkName" as an argument, which should be the path relative to the
// Init Directory, and Find the matching LinkName through the []Records in
// .linksym.yaml, UnLink it, and Remove from the []Records.
// Expects one argument, and throws error on multiple arguments provided
func Remove(args []string) error {
	switch len(args) {
	case 1:
		var sourcePath, destinationPath string
		var err error

		// Get the File Info of LinkName provided from the arguments
		linkPath, err := filePathInfo(args[0])
		if err != nil {
			return err
		} else if !linkPath.Exists {
			return fmt.Errorf("File %s doesn't exist", linkPath.AbsPath)
		}
		// Since the "filename" of the record can be the same with a different file,
		// just linked in separate directories, getting the filename and the above
		// directory name should make the name unique enough to be checked in
		// []Records
		recordPathName := filepath.Join(filepath.Base(filepath.Dir(linkPath.AbsPath)), filepath.Base(linkPath.AbsPath))

		// Can't use range over Configuration.Records. since the slice gets modified
		// during iteration with the RemoveRecord function, so iterate over the
		// slice manually in reverse order to fix this
		for i := len(config.Configuration.Records) - 1; i >= 0; i-- {
			if config.Configuration.Records[i].Name == recordPathName {
				sourcePath = config.Configuration.Records[i].Paths[0]
				destinationPath = config.Configuration.Records[i].Paths[1]

				config.RemoveRecord(i)
			}
		}

		err = linker.UnLink(sourcePath, destinationPath, linkPath.IsDir)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("Invalid number of arguments")
	}
	return nil
}
