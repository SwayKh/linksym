package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
)

func Add(args []string) error {
	switch len(args) {

	case 1:
		// Set first arg source path, get absolute path, check if it exists, set the
		// destination path as cwd+filename of source path

		sourcePath, fileExists, fileInfo, err := filePathInfo(args[0])
		if err != nil {
			return err
		} else if !fileExists {
			return fmt.Errorf("File %s doesn't exist", sourcePath)
		}

		filename := filepath.Base(sourcePath)
		destinationPath := filepath.Join(config.Configuration.InitDirectory, filename)

		return linker.MoveAndLink(sourcePath, destinationPath, fileInfo.IsDir())

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

		isSourceDir := sourceFileExists && sourceFileInfo.IsDir()
		isSourceFile := sourceFileExists && !sourceFileInfo.IsDir()
		isDestinationDir := destinationFileExists && destinationFileInfo.IsDir()
		isDestinationFile := destinationFileExists && !destinationFileInfo.IsDir()

		// For Source and Destination paths, to Exist, !Exist, be a Dir or a File
		// respectively creates 16 different combination of booleans,

		switch {
		// Destination is a directory so create symlink to destinationPath/SourceFileName
		case sourceFileExists && destinationFileExists && !sourceFileInfo.IsDir() && destinationFileInfo.IsDir():
			destinationPath = appendToDestinationPath(sourcePath, destinationPath)

			return linker.MoveAndLink(sourcePath, destinationPath, isSourceDir)

		// If the Destination path doesn't exist, It can't not be a directory or be
		// a file, check for a trailing / in the destination path to determine
		// whether to use destinationPath as a directory or as a file
		// Handles both case for !destinationFileExists and be a dir or file
		case sourceFileExists && !destinationFileExists && !sourceFileInfo.IsDir():
			if strings.HasPrefix(destinationPath, string(os.PathSeparator)) {
				destinationPath = appendToDestinationPath(sourcePath, destinationPath)
			}

			return linker.MoveAndLink(sourcePath, destinationPath, isSourceDir)

			// Put the Source Directory Path inside Destination Directory
		case sourceFileExists && destinationFileExists && sourceFileInfo.IsDir() && destinationFileInfo.IsDir():
			destinationPath = appendToDestinationPath(sourcePath, destinationPath)

			return linker.MoveAndLink(sourcePath, destinationPath, isSourceDir)

		case sourceFileExists && !destinationFileExists && sourceFileInfo.IsDir():
			if strings.HasPrefix(destinationPath, string(os.PathSeparator)) {
				destinationPath = appendToDestinationPath(sourcePath, destinationPath)
			}

			return linker.MoveAndLink(sourcePath, destinationPath, isSourceDir)

		case !sourceFileExists && destinationFileExists && !destinationFileInfo.IsDir():
			if strings.HasPrefix(sourcePath, string(os.PathSeparator)) {
				// Given Source path has a trailing /, hence it's a directory
				destinationPath = appendToDestinationPath(sourcePath, destinationPath)

				return linker.Link(sourcePath, destinationPath)

			} else {
				// Else it's a file, which should be linked the destinationPath
				return linker.Link(sourcePath, destinationPath)
			}

		// If the destination file already exists, Then the MoveAndLink() will fail
		case sourceFileExists && destinationFileExists && !sourceFileInfo.IsDir() && !destinationFileInfo.IsDir():
			return fmt.Errorf("Destination path %s already exists", destinationPath)

		case sourceFileExists && destinationFileExists && sourceFileInfo.IsDir() && !destinationFileInfo.IsDir():
		case !sourceFileExists && destinationFileExists && !sourceFileInfo.IsDir() && destinationFileInfo.IsDir():
		case !sourceFileExists && !destinationFileExists && !sourceFileInfo.IsDir() && !destinationFileInfo.IsDir():
		case !sourceFileExists && !destinationFileExists && !sourceFileInfo.IsDir() && destinationFileInfo.IsDir():
		case !sourceFileExists && destinationFileExists && sourceFileInfo.IsDir() && !destinationFileInfo.IsDir():
		case !sourceFileExists && !destinationFileExists && !sourceFileInfo.IsDir() && !destinationFileInfo.IsDir():
		case !sourceFileExists && !destinationFileExists && !sourceFileInfo.IsDir() && destinationFileInfo.IsDir():
			return fmt.Errorf("Unable to link %s to %s. \nEither the Source or Destination path don't exist, \nor There is a mismatch of types, eg - Directory to a file", sourcePath, destinationPath)

		default:
			return fmt.Errorf("Invalid arguments")

		}

	default:
		return fmt.Errorf("Invalid number of arguments")
	}
	return nil
}

// Append filename from Source path to Destination path
func appendToDestinationPath(sourcePath, destinationPath string) string {
	filename := filepath.Base(sourcePath)
	destinationPath = filepath.Join(destinationPath, filename)

	return destinationPath
}
