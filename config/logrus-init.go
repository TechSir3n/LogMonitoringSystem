package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogger() *logrus.Logger {
	var log = logrus.New()
	file, err := os.OpenFile("loggers.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Error("Failed to log to file:", file.Name())
	}

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)

	return log
}
