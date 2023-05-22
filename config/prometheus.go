package config

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequstTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Total number of requests",
		},
		[]string{"method", "endpoint"},
	)

	ResponseTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "response_time_seconds",
			Help:    "Time taken to respond to a request",
			Buckets: []float64{0.1, 0.2, 0.4, 0.8, 1.6, 3.2, 6.4, 12.8, 25.6, 51.2, 102.4},
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(RequstTotal)
	prometheus.MustRegister(ResponseTime)
}
