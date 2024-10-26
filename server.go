package main

import (
	"log"
	"math/rand/v2"
	"os"
	"time"

	"github.com/Pradumnasaraf/Contributors/handler"
	"github.com/Pradumnasaraf/Contributors/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	mynumber prometheus.Gauge
	info     *prometheus.GaugeVec
}

func main() {
	reg := prometheus.NewRegistry()
	metrics := NewMetrics(reg)

	go func() {
		for {
			metrics.mynumber.Set(float64(rand.IntN(100)))
			time.Sleep(5 * time.Second)
		}
	}()
	metrics.info.With(prometheus.Labels{"version": "2.1.2"}).Set(1)

	router := gin.Default()

	// Will bypass the middleware (Auth) for health check
	router.GET("/health", handler.HealthCheckHandler())
	router.GET("/metrics", handler.PrometheusHandler(reg))

	router.Use(middleware.BasicAuth())

	router.GET("/", handler.PlaygroundHandler())
	router.POST("/query", handler.GraphqlHandler())

	log.Printf("Server is running on http://localhost:%s", os.Getenv("PORT"))
	log.Fatal(router.Run())
}

func NewMetrics(reg prometheus.Registerer) *metrics {
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
	}
	reg.MustRegister(addMetrics.mynumber, addMetrics.info)
	return addMetrics
}
