package main

import (
	"log"

	"github.com/Pradumnasaraf/Contributors/handler"
	"github.com/Pradumnasaraf/Contributors/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// Will bypass the middleware (Auth) for health check
	router.GET("/health", handler.HealthCheckHandler())

	router.Use(middleware.BasicAuth())
	router.GET("/", handler.PlaygroundHandler())
	router.POST("/query", handler.GraphqlHandler())

	log.Fatal(router.Run())
}
