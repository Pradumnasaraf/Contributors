package main

import (
	"log"
	"os"

	"github.com/Pradumnasaraf/Contributors/config"
	"github.com/Pradumnasaraf/Contributors/graph"
	"github.com/Pradumnasaraf/Contributors/handler"
	"github.com/Pradumnasaraf/Contributors/middleware"
	database "github.com/Pradumnasaraf/Contributors/mongo"
	"github.com/Pradumnasaraf/Contributors/prometheus"
	"github.com/Pradumnasaraf/Contributors/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	// Config setup
	config.Config()

	// Database connection
	redis.RedisInit()
	mongoClient := database.MongoInit()
	graph.GetMongoClient(mongoClient)
	defer redis.RedisClose()

	// Server setup
	router := gin.Default()

	// Bypasses Auth
	router.Use(prometheus.RecordRequestLatency())
	router.Use(prometheus.RequestMetricsMiddleware())
	router.GET("/health", handler.HealthCheckHandler())
	router.GET("/metrics", handler.PrometheusHandler())

	router.Use(middleware.BasicAuthMiddleware())
	router.Use(middleware.RedisRateLimiterMiddleware())
	router.GET("/", handler.PlaygroundHandler())
	router.POST("/query", handler.GraphqlHandler())

	log.Printf("Server is running on http://localhost:%s", os.Getenv("PORT"))
	log.Fatal(router.Run())
}
