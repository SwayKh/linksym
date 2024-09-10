package config

import (
	"errors"
	"os"
)

func CheckFile(path string) (bool, os.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil, errors.New("File path provided doesn't exist")
		} else {
			return false, nil, err
		}
	}
	return true, fileInfo, nil
}
