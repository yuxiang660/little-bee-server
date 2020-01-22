package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/internal/app/routers/api"
	"github.com/yuxiang660/little-bee-server/internal/app/routers/middleware"
	"go.uber.org/dig"
)

// InitRouters initializes all routers.
func InitRouters(container *dig.Container) (*gin.Engine, error) {
	cfg := config.Global()

	gin.SetMode(cfg.RunMode)
	router := gin.New()
	
	router.NoMethod(middleware.NoMethodHandler())
	router.NoRoute(middleware.NoRouteHandler())

	router.Use(middleware.LoggerMiddleware(middleware.URLPrefixWhiteList([]string{"/api/"}...)))

	err := api.RegisterAPI(router, container)

	return router, err
}