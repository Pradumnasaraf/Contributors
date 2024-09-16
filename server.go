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
	router.POST("/query", handler.GraphqlHandler())

	// Auto catch the PORT variable if it exists
	log.Fatal(router.Run())

}
