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

// SkipperFunc defines a check function to skip the middleware.
type SkipperFunc func(*gin.Context) bool

// URLPrefixBlackList defines a black list for the middleware.
// If the URL path includes one of the prefixes, the request will skip the middleware.
func URLPrefixBlackList(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// URLPrefixWhiteList defines a white list for the middleware.
// If the URL path doesn't include any of the prefixes, the request will skip the middleware.
func URLPrefixWhiteList(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

func skipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}
