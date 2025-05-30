package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(filepath string) *logrus.Logger {
	logger := logrus.New()

	logFile := &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    100, // M
		MaxBackups: 7,
		MaxAge:     30,
	}

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	mv := io.MultiWriter(os.Stdout, logFile)
	logger.SetOutput(mv)

	return logger
}
