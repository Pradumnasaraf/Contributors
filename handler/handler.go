package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Pradumnasaraf/Contributors/graph"
	prom "github.com/Pradumnasaraf/Contributors/prometheus"
	"github.com/Pradumnasaraf/Contributors/redis"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Graphql handler
func GraphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {

		prom.HttpRequestTotal()
		metrics := prom.HttpRequestDuration()
		timer := prometheus.NewTimer(metrics.HttpRequestDuration.WithLabelValues("/query"))
		defer timer.ObserveDuration()

		err := redis.RateLimiter(c.ClientIP())
		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Playground handler
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Health handler
func HealthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	}
}

// Prometheus handler
func PrometheusHandler() gin.HandlerFunc {
	registry := prom.Registry
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
