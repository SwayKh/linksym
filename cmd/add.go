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

		// For Source and Destination paths, to Exist, !Exist, be a Dir or a File
		// respectively creates 16 different combination of booleans,
		isSourceDir := sourceFileExists && sourceFileInfo.IsDir()
		isSourceFile := sourceFileExists && !sourceFileInfo.IsDir()
		isDestinationDir := destinationFileExists && destinationFileInfo.IsDir()
		isDestinationFile := destinationFileExists && !destinationFileInfo.IsDir()

		switch {
		// Link Source File to inside of Destination directory
		case isSourceFile && isDestinationDir:
			destinationPath = appendToDestinationPath(sourcePath, destinationPath)
			return linker.MoveAndLink(sourcePath, destinationPath, isSourceDir)

		case isSourceFile && isDestinationFile:
			return fmt.Errorf("Destination path %s already exists", destinationPath)

		// Link Source file to Destination by using path as File or Directory based
		// on trailling / provided with argument
		case isSourceFile && !destinationFileExists:
			if strings.HasPrefix(destinationPath, string(os.PathSeparator)) {
				destinationPath = appendToDestinationPath(sourcePath, destinationPath)
			}
			return linker.MoveAndLink(sourcePath, destinationPath, isSourceDir)

		// Link Source Directory to inside of Destination directory
		case isSourceDir && isDestinationDir:
			destinationPath = appendToDestinationPath(sourcePath, destinationPath)
			return linker.MoveAndLink(sourcePath, destinationPath, isSourceDir)

		// Can't link a Directory to a File
		case isSourceDir && isDestinationFile:
			return fmt.Errorf("Can't Link a Directory %s to a File %s", sourcePath, destinationPath)

		// Link Source directory to Destination by using path as File or Directory
		// based on trailling / provided with argument. But can't link a Directory
		// to a File
		case isSourceDir && !destinationFileExists:
			if strings.HasPrefix(destinationPath, string(os.PathSeparator)) {
				destinationPath = appendToDestinationPath(sourcePath, destinationPath)
				return linker.MoveAndLink(sourcePath, destinationPath, isSourceDir)
			} else {
				return fmt.Errorf("Can't Link a Directory %s to a File %s", sourcePath, destinationPath)
			}

		// Source Doesn't exists(Can be file or dir), But Destination does, and is a file
		case !sourceFileExists && isDestinationFile:
			if strings.HasPrefix(sourcePath, string(os.PathSeparator)) {
				// Given Source path has a trailing /, hence it's a directory
				return fmt.Errorf("Can't Link a Directory %s to a File %s", sourcePath, destinationPath)
			} else {
				// Source is a file which doesn't exist, Destination is a file
				return linker.Link(sourcePath, destinationPath)
			}

		// Source Doesn't exists(Can be file or dir), But Destination does, and is a directory
		case !sourceFileExists && isDestinationDir:
			if strings.HasPrefix(sourcePath, string(os.PathSeparator)) {
				// Given Source path has a trailing /, hence it's a directory
				return linker.Link(sourcePath, destinationPath)
			} else {
				// Else Source is a file, and destination is a directory
				return fmt.Errorf("Can't Link a file %s to a directory %s", sourcePath, destinationPath)
			}

		// Source and Destination Both Don't Exist
		case !sourceFileExists && !destinationFileExists:
			return fmt.Errorf("Source and Destination paths don't exist, Nothing to Link")

		default:
			// return fmt.Errorf("Unable to link %s to %s. \nEither the Source or Destination path don't exist, \nor There is a mismatch of types, eg - Directory to a file", sourcePath, destinationPath)
			return fmt.Errorf("Invalid arguments provided")
		}

	default:
		return fmt.Errorf("Invalid number of arguments")
	}
}

// Append filename from Source path to Destination path
func appendToDestinationPath(sourcePath, destinationPath string) string {
	filename := filepath.Base(sourcePath)
	destinationPath = filepath.Join(destinationPath, filename)

	return destinationPath
}
