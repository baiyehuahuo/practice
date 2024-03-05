package utils

import (
	"AStream-go/consts"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func CleanFiles(dirPath string) {
	err := filepath.Walk(dirPath, func(filePath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		return os.Remove(filePath)
	})
	if err != nil {
		Fatalf("%s clean files failed: %s", consts.UtilError, err.Error())
	}
}

func DeleteFiles(dirPath string, keyword string) {
	err := filepath.Walk(dirPath, func(filePath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(info.Name(), keyword) {
			return os.Remove(filePath)
		}
		return nil
	})
	if err != nil {
		Fatalf("%s clean files failed: %s", consts.UtilError, err.Error())
	}
}
