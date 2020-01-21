package middleware

import (
	"github.com/gin-gonic/gin"
)

// NoMethodHandler handles unexpected methods.
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(400, "method error")
	}
}

// NoRouteHandler handles unexpected routers.
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(400, "router error")
	}
}