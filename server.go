package main

import (
	"log"
	"os"

	"github.com/Pradumnasaraf/Contributors/handler"
	"github.com/Pradumnasaraf/Contributors/middleware"
	"github.com/Pradumnasaraf/Contributors/prometheus"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Above BasicAuth to bypass authentication for /metrics and /health
	router.Use(prometheus.RecordRequestLatency())
	router.Use(prometheus.RequestMetricsMiddleware())
	router.GET("/health", handler.HealthCheckHandler())
	router.GET("/metrics", handler.PrometheusHandler())

	router.Use(middleware.BasicAuth())
	router.GET("/", handler.PlaygroundHandler())
	router.POST("/query", handler.GraphqlHandler())

	log.Printf("Server is running on http://localhost:%s", os.Getenv("PORT"))
	log.Fatal(router.Run())
}
