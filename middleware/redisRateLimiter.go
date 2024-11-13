package middleware

import (
	"net/http"

	"github.com/Pradumnasaraf/Contributors/redis"
	"github.com/gin-gonic/gin"
)

// Redis Rate Limiter middleware
func RedisRateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := redis.RateLimiter(c.ClientIP())
		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
