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

func CountSuffixFileSize(dirPath string, suffix string) int64 {
	var ans int64
	err := filepath.Walk(dirPath, func(filePath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), suffix) {
			ans += info.Size()
		}
		return nil
	})
	if err != nil {
		Fatalf("%s clean files failed: %s", consts.UtilError, err.Error())
	}
	return ans
}
