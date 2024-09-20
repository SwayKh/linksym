package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func CheckFile(path string) (bool, os.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil, nil
		} else {
			return false, nil, fmt.Errorf("Error getting file info: %w", err)
		}
	}
	return true, fileInfo, nil
}

// Expand the ~ and $init_directory variables to their respective values
func expandPath(path string) string {
	if strings.HasPrefix(path, "$init_directory") {
		path = strings.Replace(path, "$init_directory", InitDirectory, 1)
	}
	if strings.HasPrefix(path, "~") {
		path = strings.Replace(path, "~", HomeDirectory, 1)
	}
	return path
}

// Create aliases of ~ and $init_directory to make the paths and the
// configurations more portable
func aliasPath(path string, skipInitDir bool) string {
	if strings.HasPrefix(path, HomeDirectory) {
		path = strings.Replace(path, HomeDirectory, "~", 1)
	}
	if !skipInitDir && strings.HasPrefix(path, InitDirectory) {
		path = strings.Replace(path, InitDirectory, "$init_directory", 1)
	}

	return path
}

func GetHomePath() error {
	var err error
	HomeDirectory, err = os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Couldn't get the home directory")
	}
	return nil
}
