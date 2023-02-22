package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pradumnasaraf/go-api/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.ServeHomepage)
	router.GET("/api/movie/:id", controller.Movie)
	router.GET("/api/movies", controller.AllMovie)
	router.POST("/api/movie", controller.CreateMovie)
	router.PUT("/api/movie/:id", controller.Watched)
	router.DELETE("/api/movie/:id", controller.DeleteMovie)
	router.DELETE("/api/movies", controller.DeleteAllMovie)

	return router
}
