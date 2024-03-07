package utils

import (
	"AStream-go/consts"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var logger *logrus.Logger
var logMutex sync.Mutex

func ConfigureLogFile(logFilePath string) {
	logMutex.Lock()
	defer logMutex.Unlock()
	// 设置日志路径
	logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	// todo set formatter better
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:03:04",
	})

	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logger.Fatalf("%s %s open log file failed: %s", consts.UtilError, GetCallerName(), err)
	}
	//defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		logger.Fatalf("%s %s truncate log file failed: %s", consts.UtilError, GetCallerName(), err)
	}

	logger.SetOutput(file)
}

func Info(args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	logger.Info(args)
}

func Infof(format string, args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	logger.Warn(args)
}

func Warnf(format string, args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	logger.Warnf(format, args...)
}

func Fatal(args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	logger.Fatal(args)
}

func Fatalf(format string, args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	logger.Fatalf(format, args...)
}

func Error(args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	logger.Errorf(format, args...)
}

func Debug(args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	logger.Debugf(format, args...)
}
