package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "thrive_request_total",
			Help: "Total nubmer of requests processed by the Thrive API",
		},
		[]string{"path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "thrive_requests_error_total",
			Help: "Total number of error requests processed by Thrive API",
		},
		[]string{"path", "status"},
	)
)

func PrometheusInit() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
}

func TrackMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Middleware TrackMetrics Executando")

		path := c.Request.URL.Path
		c.Next()

		status := strconv.Itoa(c.Writer.Status())

		RequestCount.WithLabelValues(path, status).Inc()

		if c.Writer.Status() >= http.StatusBadRequest {
			ErrorCount.WithLabelValues(path, status).Inc()
		}
	}
}
