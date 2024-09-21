package cmd

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
func getFileInfo(path string) (info fileInfo, err error) {
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
