package user

import (
	// TODO: remove
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuxiang660/little-bee-server/internal/app/errors"
	"github.com/yuxiang660/little-bee-server/internal/app/controller"
	"github.com/yuxiang660/little-bee-server/internal/app/ginhelper"
	"github.com/yuxiang660/little-bee-server/internal/app/model"
	"github.com/yuxiang660/little-bee-server/internal/app/schema"
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
	// TODO: remove.
	fmt.Println("User Create")

	var user schema.User
	if err := c.ShouldBind(&user); err != nil {
		ginhelper.RespondError(c, errors.ErrBadRequestParam)
		return
	}

	// TODO: remove.
	fmt.Println("username:", user.UserName)
	fmt.Println("password:", user.Password)
	fmt.Println("recod_id:", user.RecordID)

	ginhelper.RespondError(c, errors.NoError)
}
