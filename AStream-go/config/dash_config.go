package config

import (
	"AStream-go/consts"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

const (
	// The configuration file for the AStream module
	// create logger
	LogName  = ""
	LogLevel = ""

	LogFolder       = "AStreamLogs/"
	NormalThreshold = 4

	InitialBufferingCount = 2
	RebufferingCount      = 1

	DownloadPath = "DownloadedSegment/"
	Limit        = 1 // <= InitialBufferingCount is not limit
)

var (
	now         = time.Now()
	timeSuffix  = now.Format("2006-01-02.15_04_05")
	LogFilename = "python.log"

	BufferLogFilename = path.Join(LogFolder, fmt.Sprintf("DASH_BUFFER_LOG_%s.csv", timeSuffix))
	JsonLogPath       string
	JsonABRLogPath    string
	Version           string
)

func InitConfig(version string) {
	initLogFolder(LogFolder)
	Version = version
	JsonLogPath = path.Join(LogFolder, fmt.Sprintf("%s_%s.json", version, timeSuffix))
	JsonABRLogPath = path.Join(LogFolder, fmt.Sprintf("%s-ABR_%s.json", version, timeSuffix))
}

func initLogFolder(folderPath string) {
	s, err := os.Stat(LogFolder)

	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("%s initialize log folder failed: %s", consts.ConfigError, err.Error())
		}

		// err is not exist
		if err = os.Mkdir(folderPath, os.ModePerm); err != nil {
			log.Fatalf("%s initialize log folder failed: %s", consts.ConfigError, err.Error())
		}

		return
	}

	if !s.IsDir() {
		err = errors.New(fmt.Sprintf("%s is not a folder", folderPath))
		log.Fatalf("%s initialize log folder failed: %s", consts.ConfigError, err.Error())
	}
}
