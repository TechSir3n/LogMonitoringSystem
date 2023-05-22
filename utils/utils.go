package utils

import (
	"bufio"
	"encoding/json"
	"logs-monitoring/config"
	"os"
	"strings"
)

var logger = config.InitLogger()

const MAX_FILE_SIZE = 2048

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
		} else if level == "all" {
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

func ClearLogFiles() bool {
	file, err := os.OpenFile("loggers.log", os.O_RDWR, 0755)
	if err != nil {
		logger.Error("Failed open file loggers.log", err.Error())
		return false
	}
	defer file.Close()

	fileNginx, err := os.OpenFile("/var/log/nginx/error.log.1", os.O_RDWR, 0755)
	if err != nil {
		logger.Error("Failed open file error.log", err.Error())
		return false
	}
	defer fileNginx.Close()

	stat, err := file.Stat()
	if err != nil {
		return false
	}

	statNginx, err := fileNginx.Stat()
	if err != nil {
		return false
	}

	if stat.Size() > MAX_FILE_SIZE {
		err = file.Truncate(0)
		if err != nil {
			return false
		}
		return true
	} else if statNginx.Size() > MAX_FILE_SIZE {
		err = fileNginx.Truncate(0)
		if err != nil {
			return false
		}
		return true
	}

	return false
}
