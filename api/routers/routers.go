package routers

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"logs-monitoring/api/handlers"

	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()
	
	router.Use(handlers.PromotheusMiddleware())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.POST("/get/level", func(c *gin.Context) {
		handlers.LogLevelHandler(c)
	})

	router.POST("/get/type/log", func(c *gin.Context) {
		handlers.LogFormatHandler(c)
	})

	go router.Run() // default localhost and port 8080
}
