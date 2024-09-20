package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
	"github.com/SwayKh/linksym/pkg/linker"
)

func Add(args []string) error {
	switch len(args) {

	case 1:
		// Set first arg source path, get absolute path, check if it exists, set the
		// destination path as cwd+filename of source path
		source, err := filePathInfo(args[0])
		if err != nil {
			return err
		} else if !source.Exists {
			return fmt.Errorf("File %s doesn't exist", source.AbsPath)
		}

		filename := filepath.Base(source.AbsPath)
		destinationPath := filepath.Join(config.Configuration.InitDirectory, filename)

		return linker.MoveAndLink(source.AbsPath, destinationPath, source.IsDir)

	case 2:
		// set first and second args as source and destination path, get absolute
		// paths, check if the paths exist, plus handle the special case of source
		// path not existing but destination path exists, hence creating a link
		// without the moving the files
		destination, err := filePathInfo(args[1])
		if err != nil {
			return err
		}

		source, err := filePathInfo(args[0])
		if err != nil {
			return err
		}

		// For Source and Destination paths, to Exist, !Exist, be a Dir or a File
		// respectively creates 16 different combination of booleans,
		isSourceDir := source.Exists && source.IsDir
		isSourceFile := source.Exists && !source.IsDir
		isDestinationDir := destination.Exists && destination.IsDir
		isDestinationFile := destination.Exists && !destination.IsDir

		switch {
		// Link Source File to inside of Destination directory
		case isSourceFile && isDestinationDir:
			destination.AbsPath = appendToDestinationPath(source.AbsPath, destination.AbsPath)
			return linker.MoveAndLink(source.AbsPath, destination.AbsPath, isSourceDir)

		case isSourceFile && isDestinationFile:
			return fmt.Errorf("Destination path %s already exists", destination.AbsPath)

		// Link Source file to Destination by using path as File or Directory based
		// on trailling / provided with argument
		case isSourceFile && !destination.Exists:
			if destination.HasSlash {
				destination.AbsPath = appendToDestinationPath(source.AbsPath, destination.AbsPath)
			}
			return linker.MoveAndLink(source.AbsPath, destination.AbsPath, isSourceDir)

		// Link Source Directory to inside of Destination directory
		case isSourceDir && isDestinationDir:
			destination.AbsPath = appendToDestinationPath(source.AbsPath, destination.AbsPath)
			return linker.MoveAndLink(source.AbsPath, destination.AbsPath, isSourceDir)

		// Can't link a Directory to a File
		case isSourceDir && isDestinationFile:
			return fmt.Errorf("Can't Link a Directory %s to a File %s", source.AbsPath, destination.AbsPath)

		// Link Source directory to Destination by using path as File or Directory
		// based on trailling / provided with argument. But can't link a Directory
		// to a File
		case isSourceDir && !destination.Exists:
			if destination.HasSlash {
				destination.AbsPath = appendToDestinationPath(source.AbsPath, destination.AbsPath)
				return linker.MoveAndLink(source.AbsPath, destination.AbsPath, isSourceDir)
			} else {
				return fmt.Errorf("Can't Link a Directory %s to a File %s", source.AbsPath, destination.AbsPath)
			}

		// Source Doesn't exists(Can be file or dir), But Destination does, and is a file
		case !source.Exists && isDestinationFile:
			if source.HasSlash {
				// Given Source path has a trailing /, hence it's a directory
				return fmt.Errorf("Can't Link a Directory %s to a File %s", source.AbsPath, destination.AbsPath)
			} else {
				// Source is a file which doesn't exist, Destination is a file
				return linker.Link(source.AbsPath, destination.AbsPath)
			}

		// Source Doesn't exists(Can be file or dir), But Destination does, and is a directory
		case !source.Exists && isDestinationDir:
			if source.HasSlash {
				// Given Source path has a trailing /, hence it's a directory
				return linker.Link(source.AbsPath, destination.AbsPath)
			}
			// Else Source is a file, and destination is a directory
			return fmt.Errorf("Can't Link a file %s to a directory %s", source.AbsPath, destination.AbsPath)

		// Source and Destination Both Don't Exist
		case !source.Exists && !destination.Exists:
			return fmt.Errorf("Source and Destination paths don't exist, Nothing to Link")

		default:
			// return fmt.Errorf("Unable to link %s to %s. \nEither the Source or Destination path don't exist, \nor There is a mismatch of types, eg - Directory to a file", source.AbsPath, destination.AbsPath)
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
