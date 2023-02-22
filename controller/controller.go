package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pradumnasaraf/go-api/helper"
	"github.com/pradumnasaraf/go-api/model"
)

func CreateMovie(c *gin.Context) {
	var movie model.Netflix
	err := c.BindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	helper.InsertOneMovie(movie)
	c.JSON(http.StatusCreated, gin.H{"message": "added one entry."})
}

func Watched(c *gin.Context) {

	id := c.Param("id")
	helper.UpdateOneMovie(id)
	c.JSON(http.StatusOK, gin.H{"message": "updated one entry."})
}

func DeleteMovie(c *gin.Context) {

	id := c.Param("id")
	helper.DeleteOneMovie(id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted one entry."})
}

func DeleteAllMovie(c *gin.Context) {

	deltedCount := helper.DeleteAllMovies()
	c.JSON(http.StatusOK, gin.H{"message": "deleted all entries.", "deletedCount": deltedCount})
}

func AllMovie(c *gin.Context) {

	allMovies := helper.GetAllMovies()
	if allMovies == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No movies found"})
		return
	}
	c.JSON(http.StatusOK, allMovies)
}

func Movie(c *gin.Context) {

	id := c.Param("id")
	movie := helper.GetOneMovie(id)
	c.JSON(http.StatusOK, movie)
}

func ServeHomepage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API working fine"})
}
