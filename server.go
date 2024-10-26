package main

import (
	"log"
	"os"

	"github.com/Pradumnasaraf/Contributors/handler"
	"github.com/Pradumnasaraf/Contributors/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	mynumber prometheus.Gauge
}

func main() {
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)
	m.mynumber.Set(35)

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
	m := &metrics{
		mynumber: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "my_number",
			Help:      "Number of My Number",
		}),
	}
	reg.MustRegister(m.mynumber)
	return m
}
