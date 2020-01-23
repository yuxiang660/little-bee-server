package user

import (
	"github.com/gin-gonic/gin"
	"github.com/yuxiang660/little-bee-server/internal/app/errors"
	"github.com/yuxiang660/little-bee-server/internal/app/controller"
	"github.com/yuxiang660/little-bee-server/internal/app/ginhelper"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
	"github.com/yuxiang660/little-bee-server/internal/app/model"
	"github.com/yuxiang660/little-bee-server/internal/app/model/schema"
)

// User defines the structure about user controller.
type User struct {
	model model.IUser
}

// New creates user controller.
func New(m model.IUser) controller.IUser {
	return &User{
		model: m,
	}
}

// Create creates a user with username and password.
func (u *User) Create(c *gin.Context) {
	var user schema.User
	if err := c.ShouldBind(&user); err != nil {
		logger.Error(err.Error())
		ginhelper.RespondError(c, errors.ErrBadRequestParam)
		return
	}

	if err := u.model.Create(user); err != nil {
		logger.Error(err.Error())
		ginhelper.RespondError(c, errors.ErrInternalServerError)
		return
	}

	ginhelper.RespondOK(c)
}

// Query query users with a username from client.
// Query sting: ...?user_name=xxx
func (u *User) Query(c *gin.Context) {
	var user schema.User
	if err := c.ShouldBind(&user); err != nil {
		logger.Error(err.Error())
		ginhelper.RespondError(c, errors.ErrBadRequestParam)
		return
	}

	results, err := u.model.Query(user.UserName)
	if err != nil {
		logger.Error(err.Error())
		ginhelper.RespondError(c, errors.ErrInternalServerError)
		return
	}

	ginhelper.RespondSuccess(c, results.Users)
}
