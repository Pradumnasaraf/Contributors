package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pradumnasaraf/go-api/handler"
	"github.com/pradumnasaraf/go-api/config"

)

func main() {
	// Load .env file
	config.Config()

	router := gin.Default()

	router.GET("/", handler.PlaygroundHandler())
	router.POST("/query", handler.GraphqlHandler())

	log.Fatal(router.Run())
}
