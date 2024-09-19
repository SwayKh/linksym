package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/SwayKh/linksym/pkg/config"
)

// Handle the repeating function calls in one place
func filePathInfo(path string) (absPath string, exists bool, info os.FileInfo, hasSlash bool, err error) {
	if strings.HasSuffix(path, string(os.PathSeparator)) {
		hasSlash = true
	}
	absPath, err = filepath.Abs(path)
	if err != nil {
		return "", false, nil, false, fmt.Errorf("Error getting absolute path of file %s", path)
	}

	exists, info, err = config.CheckFile(absPath)
	if err != nil {
		return "", false, nil, false, err
	}
	return absPath, exists, info, hasSlash, nil
}
