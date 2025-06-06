package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type fileInfo struct {
	AbsPath  string
	Exists   bool
	Info     os.FileInfo
	IsDir    bool
	HasSlash bool
}

// Handle the repeating function calls in one place
func GetFileInfo(path string) (fileInfo, error) {
	var err error
	info := fileInfo{}
	info.Exists = true

	if strings.HasSuffix(path, string(os.PathSeparator)) {
		info.HasSlash = true
	}

	info.AbsPath, err = filepath.Abs(path)
	if err != nil {
		return fileInfo{}, fmt.Errorf("error getting absolute path of file %s: %w", path, err)
	}

	info.Info, err = os.Stat(info.AbsPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			info.Exists = false
		} else {
			return fileInfo{}, fmt.Errorf("error getting file info: %w", err)
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
func ExpandPath(path, homeDir, initDir string) string {
	// the $init_directory strings comes from the yaml tags for AppConfig
	if strings.HasPrefix(path, "$init_directory") {
		path = strings.Replace(path, "$init_directory", initDir, 1)
	}
	if strings.HasPrefix(path, "~") {
		path = strings.Replace(path, "~", homeDir, 1)
	}
	return path
}

// Create aliases of ~ and $init_directory to make the paths and the
// configurations more portable
func AliasPath(path, homeDir, initDir string, aliasToHome bool) string {
	// the $init_directory strings comes from the yaml tags for AppConfig
	if !aliasToHome && strings.HasPrefix(path, initDir) {
		path = strings.Replace(path, initDir, "$init_directory", 1)
	}
	if strings.HasPrefix(path, homeDir) {
		path = strings.Replace(path, homeDir, "~", 1)
	}

	return path
}
