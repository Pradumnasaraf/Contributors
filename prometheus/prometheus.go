package prometheus

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	HttpRequestTotal      *prometheus.CounterVec
	HttpRequestErrorTotal *prometheus.CounterVec
	HttpRequestLatency    *prometheus.GaugeVec
}

// PrometheusRegistry is Capitalize to use in PrometheusHandler() handler in handlers.go
var PrometheusRegistry = prometheus.NewRegistry()

var prometheusMetrics = initializeMetrics(PrometheusRegistry)

func initializeMetrics(reg prometheus.Registerer) *metrics {
	requestMetrics := &metrics{
		HttpRequestTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "api_http_request_total",
			Help: "Total number of requests processed by the API",
		},
			[]string{"path", "status"}),
		HttpRequestErrorTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "api_http_request_error_total",
			Help: "Total number of errors processed by the API",
		},
			[]string{"path", "status"}),
		HttpRequestLatency: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "api_http_request_latency_milliseconds",
			Help: "Captures latency or time taken by a route to server the requests",
		},
			[]string{"path"}),
	}

	reg.MustRegister(requestMetrics.HttpRequestTotal, requestMetrics.HttpRequestErrorTotal, requestMetrics.HttpRequestLatency)
	return requestMetrics
}

func RequestMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Next()
		status := c.Writer.Status()
		if status <= 400 {
			prometheusMetrics.HttpRequestTotal.WithLabelValues(path, strconv.Itoa(status)).Inc()
			return
		}
		prometheusMetrics.HttpRequestErrorTotal.WithLabelValues(path, strconv.Itoa(status)).Inc()
	}
}

func RecordRequestLatency() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		path := c.Request.URL.Path

		c.Next()

		status := c.Writer.Size()
		latency := time.Since(t)
		if status <= 400 {
			prometheusMetrics.HttpRequestLatency.WithLabelValues(path).Set(float64(latency.Microseconds()))
			return
		}
	}
}
