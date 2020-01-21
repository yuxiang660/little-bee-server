package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterAPI register all APIs to the router.
func RegisterAPI(router *gin.Engine, container *dig.Container) error {
	return nil
}