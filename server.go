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

	// Will bypass the middleware (Auth) for health check
	router.GET("/health", handler.HealthCheckHandler())

	router.Use(prometheus.PrometheusTrackMetrics())
	router.GET("/metrics", handler.PrometheusHandler())

	router.Use(middleware.BasicAuth())
	router.GET("/", handler.PlaygroundHandler())
	router.POST("/query", handler.GraphqlHandler())

	log.Printf("Server is running on http://localhost:%s", os.Getenv("PORT"))
	log.Fatal(router.Run())
}
