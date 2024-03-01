package config

import (
	"AStream-go/consts"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

const (
	// The configuration file for the AStream module
	// create logger
	LogName  = ""
	LogLevel = ""

	LogFolder             = "AStreamLogs/"
	NormalThreshold       = 4
	BasicThreshold        = 10
	BasicUpperThreshold   = 1.2
	BasicDeltaCount       = 5
	InitialBufferingCount = 2
	RebufferingCount      = 1

	DownloadPath = "DownloadedSegment/"
)

var (
	now         = time.Now()
	timeSuffix  = fmt.Sprintf("%d-%d-%d.%d_%d_%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	LogFilename = path.Join(LogFolder, "DASH_RUNTIME_LOG")

	BufferLogFilename = path.Join(LogFolder, fmt.Sprintf("DASH_BUFFER_LOG_%s.csv", timeSuffix))
	JsonLogPath       string
	JsonABRLogPath    string
	ServerDomain      = ""
	JsonHandle        map[string]interface{}
	JsonHandleLock    = sync.Mutex{}
)

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

func InitConfig(version string) {
	initLogFolder(LogFolder)
	JsonLogPath = path.Join(LogFolder, fmt.Sprintf("%s_%s.json", version, timeSuffix))
	JsonABRLogPath = path.Join(LogFolder, fmt.Sprintf("%s-ABR_%s.json", version, timeSuffix))
	JsonHandle = make(map[string]interface{})
	JsonHandle["playback_info"] = map[string]interface{}{
		"start_time":                 nil,
		"end_time":                   nil,
		"initial_buffering_duration": nil,
		"interruptions": map[string]interface{}{
			"count":          0,
			"events":         new([][]float64),
			"total_duration": 0,
		},
		"up_shifts":   0,
		"down_shifts": 0,
	}

}
