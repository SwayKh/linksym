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
		path = strings.Replace(path, "$init_directory", Configuration.InitDirectory, 1)
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
	if !skipInitDir && strings.HasPrefix(path, Configuration.InitDirectory) {
		path = strings.Replace(path, Configuration.InitDirectory, "$init_directory", 1)
	}

	return path
}
