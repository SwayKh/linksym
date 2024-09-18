package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
)

func Remove(args []string) error {
	switch len(args) {
	case 1:
		linkName := args[0]
		var linkPath string
		var sourcePath, destinationPath string
		var err error
		var isDirectory bool

		linkPath, err = filepath.Abs(linkName)
		if err != nil {
			return fmt.Errorf("Error getting absolute path of file %s: \n%w", linkPath, err)
		}

		fileExists, fileInfo, err := config.CheckFile(linkPath)
		if err != nil {
			return err
		} else if !fileExists {
			return fmt.Errorf("File %s doesn't exist", linkPath)
		} else if fileInfo.IsDir() {
			isDirectory = true
		}

		recordPathName := filepath.Join(filepath.Base(filepath.Dir(linkPath)), filepath.Base(linkPath))

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

		err = linker.UnLink(sourcePath, destinationPath, isDirectory)
		if err != nil {
			return nil
		}

	default:
		return fmt.Errorf("Invalid number of arguments")
	}
	return nil
}
