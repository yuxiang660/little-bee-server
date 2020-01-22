package login

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuxiang660/little-bee-server/internal/app/auther"
	"github.com/yuxiang660/little-bee-server/internal/app/errors"
	"github.com/yuxiang660/little-bee-server/internal/app/controller"
	"github.com/yuxiang660/little-bee-server/internal/app/ginhelper"
	"github.com/yuxiang660/little-bee-server/internal/app/schema"
)

// Login defines the structure about login controller.
type Login struct {
	a auther.Auther
}

// New creates login controller.
func New(a auther.Auther) controller.ILogin {
	return &Login{
		a: a,
	}
}

// In verifies username and password and generate a token to client.
func (a *Login) In(c *gin.Context) {
	fmt.Println("Login")

	var cred schema.LoginParam
	if err := c.ShouldBind(&cred); err != nil {
		ginhelper.RespondError(c, errors.ErrBadRequestParam)
		return
	}
	fmt.Println(cred.UserName)
	fmt.Println(cred.Password)

	user, err := verify(cred.UserName, cred.Password)
	if err != nil {
		ginhelper.RespondError(c, err.(errors.Error))
		return
	}

	userID := user.RecordID
	ginhelper.SetUserID(c, userID)

	ginhelper.RespondError(c, errors.NoError)
}

// TODO: retrieve username and password from database.
func verify(userName, password string) (*schema.User, error) {
	user := schema.User {
		RecordID: "aaa",
		UserName: "ben",
		Password: "123",
	}

	if userName != user.UserName {
		return nil, errors.ErrInvalidUsername
	}

	if password != user.Password {
		return nil, errors.ErrInvalidPassword
	}

	return &user, nil
}

// TODO: it is a placeholder here, implement this function.
// Out destroys the token for the login client.
func (a *Login) Out(c *gin.Context) {
	fmt.Println("Logout")
}
