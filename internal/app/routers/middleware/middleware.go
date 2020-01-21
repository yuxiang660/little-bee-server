package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yuxiang660/little-bee-server/internal/app/errors"
)

// NoMethodHandler handles unexpected methods.
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(errors.ErrorNotFound.Code(), errors.ErrorNotFound.Body())
	}
}

// NoRouteHandler handles unexpected routers.
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(errors.ErrorNotFound.Code(), errors.ErrorNotFound.Body())
	}
}