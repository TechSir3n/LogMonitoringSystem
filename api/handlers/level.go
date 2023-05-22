package handlers

import (
	"github.com/gin-gonic/gin"
	"logs-monitoring/config"
	"logs-monitoring/models"
	"logs-monitoring/utils"
	"net/http"
)

var logger = config.InitLogger()

// this function to select a certain level of logs and show it to the user
func LogLevelHandler(c *gin.Context) {
	var level models.LogLevel
	if err := c.ShouldBindJSON(&level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json fields"})
		logger.Error("Failed to convert struct to json format", err.Error())
	}

	msg, err := utils.ReadFile(level.Level)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find level logs"})
		logger.Error(err.Error())
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"level": level.Level, "message": msg})
	}
}
