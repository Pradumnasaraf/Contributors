package middleware

import (
	"os"

	"github.com/Pradumnasaraf/Contributors/config"
	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {

	config.Config()

	username := os.Getenv("BASIC_AUTH_USERNAME")
	password := os.Getenv("BASIC_AUTH_PASSWORD")
	return gin.BasicAuth(gin.Accounts{
		username: password,
	})
}
