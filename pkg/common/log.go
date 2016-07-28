package common

import (
	"github.com/Sirupsen/logrus"
)

type logger struct {
	logger *logrus.Logger
}

var Log logger

func initLogger() {
	Log.logger = logrus.New()
	Log.logger.Formatter = new(logrus.TextFormatter)
	if AppConfig.DebugMode {
		Log.logger.Level = logrus.DebugLevel
	} else {
		Log.logger.Level = logrus.InfoLevel
	}
}

func (l *logger) Info(args ...interface{}) {
	l.logger.Info(args)
}

func (l *logger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}

func (l *logger) Warn(args ...interface{}) {
	l.logger.Warn(args)
}
