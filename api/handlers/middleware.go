package handlers

import (
	"time"

	"logs-monitoring/config"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func PromotheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		
		c.Next()

		end := time.Now()

		duration := end.Sub(start).Seconds()

		config.RequstTotal.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.Request.URL.Path}).Inc()

		config.ResponseTime.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.Request.URL.Path}).Observe(duration)
	}
}
