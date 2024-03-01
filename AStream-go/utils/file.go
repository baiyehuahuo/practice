package utils

import (
	"AStream-go/consts"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func CleanFiles(dirPath string) {
	err := filepath.Walk(dirPath, func(filePath string, info fs.FileInfo, err error) error {
		return os.Remove(filePath)
	})
	if err != nil {
		log.Fatalf("%s clean files failed: %s", consts.UtilError, err.Error())
	}
}
