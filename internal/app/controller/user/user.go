package user

import (
	// TODO: remove
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuxiang660/little-bee-server/internal/app/errors"
	"github.com/yuxiang660/little-bee-server/internal/app/controller"
	"github.com/yuxiang660/little-bee-server/internal/app/ginhelper"
	"github.com/yuxiang660/little-bee-server/internal/app/schema"
	"github.com/yuxiang660/little-bee-server/internal/app/store"
)

// User defines the structure about user controller.
type User struct {
	db store.Store
}

// New creates user controller.
func New(db store.Store) controller.IUser {
	return &User{
		db: db,
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
