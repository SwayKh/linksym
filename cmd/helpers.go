package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/SwayKh/linksym/pkg/config"
)

type fileInfo struct {
	AbsPath  string
	Exists   bool
	Info     os.FileInfo
	IsDir    bool
	HasSlash bool
}

// Handle the repeating function calls in one place
func filePathInfo(path string) (info fileInfo, err error) {
	info = fileInfo{}

	if strings.HasSuffix(path, string(os.PathSeparator)) {
		info.HasSlash = true
	}

	info.AbsPath, err = filepath.Abs(path)
	if err != nil {
		return fileInfo{}, fmt.Errorf("Error getting absolute path of file %s", path)
	}

	info.Exists, info.Info, err = config.CheckFile(info.AbsPath)
	if err != nil {
		return fileInfo{}, err
	}

	// If file doesn't exist, the info.IsDir check will return an nil pointer
	// dereferencing error
	if info.Exists && info.Info.IsDir() {
		info.IsDir = true
	}

	return info, nil
}
