package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	HttpRequestTotal    *prometheus.CounterVec
	HttpRequestDuration *prometheus.HistogramVec
}

var Registry = prometheus.NewRegistry()
var myMetrics = newMetrics(Registry)

func newMetrics(reg prometheus.Registerer) *metrics {
	defineMetrics := &metrics{
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

	reg.MustRegister(defineMetrics.HttpRequestDuration, defineMetrics.HttpRequestTotal)
	return defineMetrics
}

func HttpRequestTotal() {

	myMetrics.HttpRequestTotal.WithLabelValues("/query").Inc()
}

func HttpRequestDuration() *metrics {
	return myMetrics
}
