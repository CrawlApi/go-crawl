package common

import (
	"github.com/Sirupsen/logrus"
)

var logger *logrus.Logger

func initLogger() {
	logger = logrus.New()
	logger.Formatter = new(logrus.TextFormatter)
	if AppConfig.DebugMode {
		logger.Level = logrus.DebugLevel
	} else {
		logger.Level = logrus.InfoLevel
	}

}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}