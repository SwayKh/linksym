package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/config"
	"github.com/SwayKh/linksym/link"
	"github.com/SwayKh/linksym/logger"
)

// Add function, which handles the Add subcommand and handles all scenarios of
// file paths provided.
// Handling one argument is simple enough. But, handling 2 arguments creates
// lots of different scenario of combination of files and directory, and
// handling the special scenario of a File/Dir which is already moved by the
// user, and just needs to be linked, Skipping the moving of file step of the
// Linking process
// toLink boolean decided whether to perform the Move/Link action or just add
// record of "linking" to the .linksym.yaml file. Useful for when a symlink
// already exists, but they record of it doesn't
// updateRecord bool is used for running the Add function in source subcommand
// since the Source command doesn't need to update the Records, just Link stuff
func (app *Application) Add(args []string, toLink bool, updateRecord bool) error {
	toMove := true

	if len(args) == 1 {
		source, err := config.GetFileInfo(args[0])
		if err != nil {
			return err
		}

		if !source.Exists {
			return fmt.Errorf("file %s doesn't exist", source.AbsPath)
		}

		logger.VerboseLog(logger.SUCCESS, "Source path exists: %s", config.AliasPath(source.AbsPath, app.HomeDirectory, app.InitDirectory, true))

		sourcePath := source.AbsPath
		filename := filepath.Base(sourcePath)
		destinationPath := filepath.Join(app.InitDirectory, filename)

		logger.VerboseLog(logger.SUCCESS, "Destination path exists: %s", config.AliasPath(destinationPath, app.HomeDirectory, app.InitDirectory, true))

		// Stop linking if the source path is already a symlink pointing towards the
		// destination path
		isLink, err := checkSymlink(sourcePath, destinationPath)
		if err != nil {
			return err
		}

		if isLink {
			logger.Log(logger.WARNING, "Symlink already exists")
			return nil
		}

		paths := link.LinkPaths{
			SourcePath:      sourcePath,
			DestinationPath: destinationPath,
			HomeDir:         app.HomeDirectory,
			InitDir:         app.InitDirectory,
			IsDirectory:     source.IsDir,
		}

		if toLink {
			err = paths.MoveAndLink()
			if err != nil {
				return err
			}
		}
		if updateRecord {
			app.Configuration.AddRecord(sourcePath, destinationPath)
		}

	} else if len(args) == 2 {
		source, err := config.GetFileInfo(args[0])
		if err != nil {
			return err
		}

		destination, err := config.GetFileInfo(args[1])
		if err != nil {
			return err
		}

		// For Source and Destination paths, to Exist, !Exist, be a Dir or a File
		// respectively creates 16 different combination of booleans,
		isSourceDir := source.Exists && source.IsDir
		isSourceFile := source.Exists && !source.IsDir
		isDestinationDir := destination.Exists && destination.IsDir
		isDestinationFile := destination.Exists && !destination.IsDir

		sourcePath := source.AbsPath
		destinationPath := destination.AbsPath

		aliasSourcePath := config.AliasPath(source.AbsPath, app.HomeDirectory, app.InitDirectory, true)
		aliasDestinationPath := config.AliasPath(destination.AbsPath, app.HomeDirectory, app.InitDirectory, true)

		logger.VerboseLog(logger.SUCCESS, "Source path: %s", aliasSourcePath)
		logger.VerboseLog(logger.SUCCESS, "Destination path: %s", aliasDestinationPath)

		// Stop linking if the source path is already a symlink pointing towards the
		// destination path
		isLink, err := checkSymlink(sourcePath, destinationPath)
		if err != nil {
			return err
		}

		if isLink {
			logger.Log(logger.WARNING, "Symlink already exists")
			return nil
		}

		switch {
		// Link Source File to inside of Destination directory
		case isSourceFile && isDestinationDir:
			destinationPath = appendToDestinationPath(source.AbsPath, destination.AbsPath)

		case isSourceFile && isDestinationFile:
			// call Link Function, which removes the sourceFile and creates and symlink
			err = link.DeleteFile(sourcePath)
			if err != nil {
				return err
			}
			toMove = false

			// The Files should be overwritten, just like directories are, So this
			// error shouldn't be returned, like 'ln' utility returns error when
			// linking 2 files.
			// return fmt.Errorf("destination file %s already exists", aliasDestinationPath)

		// Link Source file to Destination by using path as File or Directory based
		// on trailling / provided with argument
		case isSourceFile && !destination.Exists:
			if destination.HasSlash {
				err := os.MkdirAll(destinationPath, 0o755)
				if err != nil {
					return err
				}
				destinationPath = appendToDestinationPath(source.AbsPath, destination.AbsPath)
			}

		// Link Source Directory to inside of Destination directory
		case isSourceDir && isDestinationDir:
			destinationPath = appendToDestinationPath(source.AbsPath, destination.AbsPath)

		// Can't link a Directory to a File
		case isSourceDir && isDestinationFile:
			return fmt.Errorf("can't link a Directory: %s to a File: %s", aliasSourcePath, aliasDestinationPath)

		// Link Source directory to Destination by using path as File or Directory
		// based on trailling / provided with argument. But can't link a Directory
		// to a File
		case isSourceDir && !destination.Exists:
			if destination.HasSlash {
				err := os.MkdirAll(destinationPath, 0o755)
				if err != nil {
					return err
				}
				destinationPath = appendToDestinationPath(source.AbsPath, destination.AbsPath)
			} else {
				return fmt.Errorf("can't link a Directory: %s to a File: %s", aliasSourcePath, aliasDestinationPath)
			}

		// Source Doesn't exists, But Destination does, and is a file and the Source
		// can be a directory path or a file path
		case !source.Exists && isDestinationFile:
			// Don't need the whole HasSlash check, symlink gets created correctly
			// anyway, since the MoveAndLink does not get called, so it doesn't matter
			// if Source path provided is of a directory or not
			toMove = false

		// Source Doesn't exists(Can be file or dir), But Destination does, and is a directory
		case !source.Exists && isDestinationDir:
			toMove = false

		// Source and Destination Both Don't Exist
		case !source.Exists && !destination.Exists:
			return fmt.Errorf("source and Destination paths don't exist, Nothing to Link")

		default:
			return fmt.Errorf("invalid arguments provided")
		}

		paths := link.LinkPaths{
			SourcePath:      sourcePath,
			DestinationPath: destinationPath,
			HomeDir:         app.HomeDirectory,
			InitDir:         app.InitDirectory,
			IsDirectory:     source.IsDir,
		}

		if toLink {
			if toMove {
				err = paths.MoveAndLink()
			} else {
				err = paths.Link()
			}
			if err != nil {
				return err
			}
		}
		if updateRecord {
			app.Configuration.AddRecord(sourcePath, destinationPath)
		}

	} else {
		return fmt.Errorf("invalid number of arguments")
	}
	return nil
}

// Append filename from Source path to Destination path
func appendToDestinationPath(sourcePath, destinationPath string) string {
	filename := filepath.Base(sourcePath)
	destinationPath = filepath.Join(destinationPath, filename)

	return destinationPath
}

func checkSymlink(src, dst string) (bool, error) {
	srcIsLinkToDst := false

	sourceData, err := config.GetFileInfo(src)
	if err != nil {
		return srcIsLinkToDst, err
	}

	if !sourceData.Exists {
		return false, nil
	}

	sourceLink, err := os.Lstat(src)
	if err != nil {
		return srcIsLinkToDst, fmt.Errorf("error reading symlink: %s", src)
	}

	// True for symlink
	if sourceLink.Mode()&os.ModeSymlink != 0 {
		link, err := os.Readlink(src)
		if err != nil {
			return srcIsLinkToDst, fmt.Errorf("error reading destination of link: %s", src)
		}
		if link == dst {
			srcIsLinkToDst = true
		}
	}

	return srcIsLinkToDst, nil
}
