package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(filepath string) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(&lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    100, // M
		MaxBackups: 7,
		MaxAge:     30,
	})

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}
