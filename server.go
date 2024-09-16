package main

import (
	"log"

	"github.com/Pradumnasaraf/Contributors/handler"
	"github.com/Pradumnasaraf/Contributors/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(middleware.BasicAuth())
	router.GET("/", handler.PlaygroundHandler())
	router.GET("/health", handler.HealthCheckHandler())
	router.POST("/query", handler.GraphqlHandler())

	log.Fatal(router.Run())
}
