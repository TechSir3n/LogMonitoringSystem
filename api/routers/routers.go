package routers

import (
	"github.com/gin-gonic/gin"
	"logs-monitoring/api/handlers"
)

func init() {
	router := gin.Default()

	router.POST("/get/level", func(c *gin.Context) {
		handlers.LogLevelHandler(c)
	})

	router.POST("/get/type/log", func(c *gin.Context) {
		handlers.LogFormatHandler(c)
	})

	go router.Run() // default localhost and port 8080
}
