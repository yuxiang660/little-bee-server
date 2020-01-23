package controller

import (
	"github.com/gin-gonic/gin"
)

// ILogin defines the interface to manager login controller.
type ILogin interface {
	In(c *gin.Context)
	Out(c *gin.Context)
}

// IUser defines the interface to manager user controller.
type IUser interface {
	Create(c *gin.Context)
	Query(c *gin.Context)
}
