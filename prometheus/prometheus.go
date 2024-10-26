package prometheus

import (
	"math/rand/v2"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	mynumber prometheus.Gauge
	info     *prometheus.GaugeVec
	counter  *prometheus.CounterVec
}

func WriteMetrics(reg *prometheus.Registry) {

	metrics := newMetrics(reg)

	go func() {
		for {
			metrics.mynumber.Set(float64(rand.IntN(100)))
			time.Sleep(5 * time.Second)
		}
	}()
	metrics.info.With(prometheus.Labels{"version": "2.1.2"}).Set(1)
}
func newMetrics(reg prometheus.Registerer) *metrics {
	addMetrics := &metrics{
		mynumber: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "my_number",
			Help:      "Number of My Number",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "info",
			Help:      "Info about app version",
		},
			[]string{"version"}),
		counter: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "myapp",
			Name:      "my_counter",
			Help:      "count",
		},
			[]string{"type"}),
	}

	reg.MustRegister(addMetrics.mynumber, addMetrics.info, addMetrics.counter)
	return addMetrics
}
