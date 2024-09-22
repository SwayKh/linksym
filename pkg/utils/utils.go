package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/SwayKh/linksym/pkg/global"
)

type fileInfo struct {
	AbsPath  string
	Exists   bool
	Info     os.FileInfo
	IsDir    bool
	HasSlash bool
}

// Handle the repeating function calls in one place
func GetFileInfo(path string) (info fileInfo, err error) {
	info = fileInfo{}
	info.Exists = true

	if strings.HasSuffix(path, string(os.PathSeparator)) {
		info.HasSlash = true
	}

	info.AbsPath, err = filepath.Abs(path)
	if err != nil {
		return fileInfo{}, fmt.Errorf("Error getting absolute path of file %s: %w", path, err)
	}

	info.Info, err = os.Stat(info.AbsPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			info.Exists = false
		} else {
			return fileInfo{}, fmt.Errorf("Error getting file info: %w", err)
		}
	}

	// If file doesn't exist, the info.IsDir check will return an nil pointer
	// dereferencing error
	if info.Exists && info.Info.IsDir() {
		info.IsDir = true
	}

	return info, nil
}

// Expand the ~ and $init_directory variables to their respective values
func ExpandPath(path string) string {
	// the $init_directory strings comes from the yaml tags for AppConfig
	if strings.HasPrefix(path, "$init_directory") {
		path = strings.Replace(path, "$init_directory", global.InitDirectory, 1)
	}
	if strings.HasPrefix(path, "~") {
		path = strings.Replace(path, "~", global.HomeDirectory, 1)
	}
	return path
}

// Create aliases of ~ and $init_directory to make the paths and the
// configurations more portable
func AliasPath(path string, skipInitDir bool) string {
	// the $init_directory strings comes from the yaml tags for AppConfig
	if !skipInitDir && strings.HasPrefix(path, global.InitDirectory) {
		path = strings.Replace(path, global.InitDirectory, "$init_directory", 1)
	}
	if strings.HasPrefix(path, global.HomeDirectory) {
		path = strings.Replace(path, global.HomeDirectory, "~", 1)
	}

	return path
}

// Set values to global variables of InitDirectory and ConfigPath
func SetupDirectories(initDir string, configName string) {
	global.InitDirectory = ExpandPath(initDir)
	global.ConfigPath = filepath.Join(global.InitDirectory, configName)
}

// Set the global HomeDirectory variable. Separated from SetupDirectories to be
// used with the Init Subcommand.
func InitialiseHomePath() error {
	var err error
	global.HomeDirectory, err = os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Couldn't get the home directory")
	}
	return nil
}
