package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pradumnasaraf/go-api/config"
	"github.com/pradumnasaraf/go-api/handler"
	"github.com/pradumnasaraf/go-api/middleware"
)

func main() {
	// Load .env file
	config.Config()

	router := gin.Default()

	router.Use(middleware.BasicAuth())
	router.GET("/", handler.PlaygroundHandler())
	router.POST("/query", handler.GraphqlHandler())

	// Auto catch the PORT variable if it exists
	log.Fatal(router.Run())
}
