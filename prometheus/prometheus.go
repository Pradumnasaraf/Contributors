package prometheus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	HttpRequestTotal      *prometheus.CounterVec
	HttpRequestErrorTotal *prometheus.CounterVec
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
	}

	reg.MustRegister(requestMetrics.HttpRequestTotal, requestMetrics.HttpRequestErrorTotal)
	return requestMetrics
}

func PrometheusTrackMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Next()
		status := c.Writer.Status()
		prometheusMetrics.HttpRequestTotal.WithLabelValues(path, http.StatusText(status)).Inc()
		if status >= 400 {
			prometheusMetrics.HttpRequestErrorTotal.WithLabelValues(path, http.StatusText(status)).Inc()
		}
	}
}
