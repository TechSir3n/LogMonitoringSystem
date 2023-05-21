package routers

import (
	"logs-monitoring/api/handlers"
	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()

	router.POST("/get/level", func(c *gin.Context) {
		handlers.LogLevelHandler(c)
	})

	router.Run() // default localhost and port 8080 
}
