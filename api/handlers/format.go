package handlers

import (
	"logs-monitoring/models"
	"logs-monitoring/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

func LogFormatHandler(c *gin.Context) {
	var logType models.LogType
	if err := c.ShouldBindJSON(&logType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't convernt to json"})
		logger.Error("Failed to convert struct to json format", err.Error())
		return 
	}

	data := utils.GetFormatTypeLog(logType.Type)
	if logType.Type == "xml" {
		c.Header("Content-Type", "text/xml")
		c.XML(http.StatusOK, data)
	} else if logType.Type == "json" {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"logs": data})
	} else if logType.Type == "csv" {
		c.Header("Content-Type", "text/csv")
		c.String(http.StatusOK, data)
	} else { 
		c.JSON(http.StatusBadRequest,gin.H{"error":" unsupported log type"})
		logger.Error("Unsupported log type", logType.Type)
	}
}
