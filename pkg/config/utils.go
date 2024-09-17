package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func SetupDirectories() error {
	var err error
	HomeDirectory, err = os.UserHomeDir()
	if err != nil {
		return errors.New("Couldn't get the home directory")
	}

	InitDirectory, err = os.Getwd()
	if err != nil {
		return errors.New("Couldn't get the current working directory")
	}

	ConfigName := ".linksym.yaml"

	ConfigPath = filepath.Join(InitDirectory, ConfigName)
	return nil
}

func CheckFile(path string) (bool, os.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil, nil
		} else {
			return false, nil, fmt.Errorf("Error getting file info: \n%w", err)
		}
	}
	return true, fileInfo, nil
}

func expandPath(paths []string) []string {
	for i, path := range paths {
		if strings.HasPrefix(path, "$init_directory") {
			paths[i] = strings.Replace(path, "$init_directory", InitDirectory, 1)
		} else if strings.HasPrefix(path, "~") {
			paths[i] = strings.Replace(path, "~", HomeDirectory, 1)
		}
	}
	return paths
}

func aliasPath(paths []string) []string {
	for i, path := range paths {
		if strings.HasPrefix(path, InitDirectory) {
			paths[i] = strings.Replace(path, InitDirectory, "$init_directory", 1)
		} else if strings.HasPrefix(path, HomeDirectory) {
			paths[i] = strings.Replace(path, HomeDirectory, "~", 1)
		}
	}
	return paths
}
