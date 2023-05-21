package utils

import (
	"bufio"
	"encoding/json"
	"logs-monitoring/config"
	"os"
	"strings"
)

var logger = config.InitLogger()

type LogLevel struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func ReadFile(level string) (string, error) {
	file, err := os.Open("/var/log/nginx/error.log.1")
	if err != nil {
		logger.Fatal("Error -> function[ReadFile]: ", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result strings.Builder

	for scanner.Scan() {
		var log LogLevel
		err := json.Unmarshal(scanner.Bytes(), &log)
		if err != nil {
			logger.Error("Failed unmarhshal,function[ReadFile]", err.Error())
		}

		if level == log.Level {
			result.WriteString(log.Message)
		} else {
			logger.Error("Didn't find match level in file")
		}
	}

	if err := scanner.Err(); err != nil {
		logger.Error(err.Error())
	}

	return result.String(), nil
}
