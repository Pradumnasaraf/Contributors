package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	HttpRequestTotal    *prometheus.CounterVec
	HttpRequestDuration *prometheus.HistogramVec
}

var PrometheusRegistry = prometheus.NewRegistry()
var prometheusMetrics = initializeMetrics(PrometheusRegistry)

func initializeMetrics(reg prometheus.Registerer) *metrics {
	requestMetrics := &metrics{
		HttpRequestTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Total number of HTTP requests to the API.",
		},
			[]string{"path"}),
		HttpRequestDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{0.01, 0.015, 0.02, 0.025, 0.03},
		},
			[]string{"path"}),
	}

	reg.MustRegister(requestMetrics.HttpRequestDuration, requestMetrics.HttpRequestTotal)
	return requestMetrics
}

func HttpRequestTotal() {

	prometheusMetrics.HttpRequestTotal.WithLabelValues("/query").Inc()
}

func HttpRequestDuration() *metrics {
	return prometheusMetrics
}
