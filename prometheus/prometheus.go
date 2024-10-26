package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	noOfRequests prometheus.Gauge
	info         *prometheus.GaugeVec
	counter      *prometheus.CounterVec
}

var Registry = prometheus.NewRegistry()
var myMetrics = newMetrics(Registry)

func WriteMetrics() {

	myMetrics.info.With(prometheus.Labels{"version": "2.0.3"}).Set(6)
}

func newMetrics(reg prometheus.Registerer) *metrics {
	addMetrics := &metrics{
		noOfRequests: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "no_of_requests",
			Help: "Total number of request to the API server",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "info",
			Help: "Info about app version",
		},
			[]string{"version"}),
		counter: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "my_counter",
			Help: "count",
		},
			[]string{"type"}),
	}

	reg.MustRegister(addMetrics.noOfRequests, addMetrics.info, addMetrics.counter)
	return addMetrics
}

func NumberOfRequests() {

	myMetrics.noOfRequests.Inc()

}
