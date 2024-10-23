package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {

	username := os.Getenv("BASIC_AUTH_USERNAME")
	password := os.Getenv("BASIC_AUTH_PASSWORD")
	return gin.BasicAuth(gin.Accounts{
		username: password,
	})
}
