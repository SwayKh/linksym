package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SwayKh/linksym/pkg/config"
)

// Handle the repeating function calls in one place
func filePathInfo(path string) (absPath string, exists bool, info os.FileInfo, err error) {
	absPath, err = filepath.Abs(path)
	if err != nil {
		return "", false, nil, fmt.Errorf("Error getting absolute path of file %s", path)
	}

	exists, info, err = config.CheckFile(absPath)
	if err != nil {
		return "", false, nil, err
	}
	return absPath, exists, info, nil
}
