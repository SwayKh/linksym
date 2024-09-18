package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
)

func Add(args []string) error {
	var sourcePath, destinationPath string
	var err error
	var isDirectory bool

	switch len(args) {

	case 1:
		// Set first arg source path, get absolute path, check if it exists, set the
		// destination path as cwd+filename of source path

		sourcePath, fileExists, fileInfo, err := filePathInfo(args[0])
		if err != nil {
			return err
		} else if !fileExists {
			return fmt.Errorf("File %s doesn't exist", sourcePath)
		} else if fileInfo.IsDir() {
			isDirectory = true
		}

		filename := filepath.Base(sourcePath)
		destinationPath = filepath.Join(config.Configuration.InitDirectory, filename)

		err = linker.MoveAndLink(sourcePath, destinationPath, isDirectory)
		if err != nil {
			return err
		}
		return nil

	case 2:
		// set first and second args as source and destination path, get absolute
		// paths, check if the paths exist, plus handle the special case of source
		// path not existing but destination path exists, hence creating a link
		// without the moving the files
		destinationPath, destinationFileExists, destinationFileInfo, err := filePathInfo(args[1])
		if err != nil {
			return err
		}

		sourcePath, sourceFileExists, sourceFileInfo, err := filePathInfo(args[0])
		if err != nil {
			return err
		}

		// For Source and Destination paths, to Exist, !Exist, be a Dir or a File
		// respectively creates 16 different combination of booleans, out of which 8
		// should result in an error, but still need to be handled to provide useful
		// information/error to the user.

		// This is the branching bullshit that happens with so many booleans
		// if sourceFileExists {
		// 	// Source is a file
		// 	if !sourceFileInfo.IsDir() {
		// 		if destinationFileExists {
		// 			// Destination is a file
		// 			if !destinationFileInfo.IsDir() {
		// 				// Destructive action, existing Destination file will be deleted
		// 				// Link it with a different name i.e., Destination file name
		// 				linker.Link()
		// 			} else if destinationFileInfo.IsDir() {
		// 				// Can't links a file to a directory
		// 				return err
		// 			}
		// 		} else if !destinationFileExists {
		// 			if !destinationFileInfo.IsDir() {
		// 				// Link it with a different name i.e., Destination file name
		// 				linker.Link()
		// 			} else if destinationFileInfo.IsDir() {
		// 				// given destination is in form dir/filename
		// 				linker.Link()
		// 			}
		// 		}
		// 	} else if sourceFileInfo.IsDir() {
		// 		if destinationFileExists {
		// 			if !destinationFileInfo.IsDir() {
		// 				// Can't link a directory to a file
		// 				return err
		// 			} else if destinationFileInfo.IsDir() {
		// 				// destination = destination/sourceDir/
		// 				linker.Link()
		// 			}
		// 		} else if !destinationFileExists {
		// 			// Treat the destinationPath as a Directory
		// 			// destination = destinationPath/sourceDir/
		// 			linker.Link()
		// 		}
		// 	}
		// } else if !sourceFileExists {
		// 	if !sourceFileInfo.IsDir() {
		// 		if destinationFileExists {
		// 			if !destinationFileInfo.IsDir() {
		// 				// Source doesn't exist, but destination exists, so create a "fake"
		// 				// link, skips a step in the function call for Link()
		// 				linker.Link()
		// 			} else if destinationFileInfo.IsDir() {
		// 				// Can't link a file to a directory
		// 				return err
		// 			}
		// 		} else if !destinationFileExists {
		// 			if !destinationFileInfo.IsDir() {
		// 				// if Source and Destination both don't exists then...????
		// 				return err
		// 			} else if destinationFileInfo.IsDir() {
		// 				// if Source and Destination both don't exists then...????
		// 				return err
		// 			}
		// 		}
		// 	} else if sourceFileInfo.IsDir() {
		// 		if destinationFileExists {
		// 			if !destinationFileInfo.IsDir() {
		// 				// Can't link a directory to a file
		// 				return err
		// 			} else if destinationFileInfo.IsDir() {
		// 				// Source is a directory that doesn't exists, Destination is a
		// 				// directory that exists, Same case as both source and destination
		// 				// paths being a file, Create an "Fake" link
		// 				linker.Link()
		// 			}
		// 		} else if !destinationFileExists {
		// 			// Treat the destinationPath as a Directory
		// 			return err
		// 		}
		// 	}
		// }

		if sourceFileExists && sourceFileInfo.IsDir() && destinationFileExists && destinationFileInfo.IsDir() {
			filename := filepath.Base(sourcePath)
			destinationPath = filepath.Join(destinationPath, filename)
			isDirectory = true

			err = linker.MoveAndLink(sourcePath, destinationPath, isDirectory)
			if err != nil {
				return err
			}
			return nil
		}

		if sourceFileExists && sourceFileInfo.IsDir() && destinationFileExists {
			filename := filepath.Base(destinationPath)
			sourcePath = filepath.Join(sourcePath, filename)

			err := linker.Link(sourcePath, destinationPath)
			if err != nil {
				return err
			}
			return nil
		}

		if destinationFileExists && !sourceFileExists {
			err := linker.Link(sourcePath, destinationPath)
			if err != nil {
				return err
			}
			return nil
		}

	default:
		return fmt.Errorf("Invalid number of arguments")
	}

	err = linker.MoveAndLink(sourcePath, destinationPath, isDirectory)
	if err != nil {
		return err
	}
	return nil
}
