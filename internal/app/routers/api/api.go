package api

import (
	"github.com/yuxiang660/little-bee-server/internal/app/auther"
	"github.com/yuxiang660/little-bee-server/internal/app/routers/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterAPI register all APIs to the router.
func RegisterAPI(router *gin.Engine, container *dig.Container) error {
	return container.Invoke(func(
		a auther.Auther,
	) error {
		api := router.Group("/api")
		{
			api.Use(middleware.UserAuthMiddleware(a,
				middleware.SkipPrefixList("/api/v1/pub/login"),
			))
		}

		return nil
	})
}