package login

import (
	"github.com/gin-gonic/gin"
	"github.com/yuxiang660/little-bee-server/internal/app/auther"
	"github.com/yuxiang660/little-bee-server/internal/app/errors"
	"github.com/yuxiang660/little-bee-server/internal/app/controller"
	"github.com/yuxiang660/little-bee-server/internal/app/ginhelper"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
	"github.com/yuxiang660/little-bee-server/internal/app/model"
	"github.com/yuxiang660/little-bee-server/internal/app/model/schema"
)

// Login defines the structure about login controller.
type Login struct {
	auth auther.Auther
	users model.IUser
}

// New creates login controller.
func New(a auther.Auther, u model.IUser) controller.ILogin {
	return &Login{
		auth: a,
		users: u,
	}
}

func (l *Login) respondWithToken(c *gin.Context, userID string) {
	ginhelper.SetUserID(c, userID)

	tokenInfo, err := l.auth.GenerateToken(userID)
	if err != nil {
		logger.Error(err.Error())
		ginhelper.RespondError(c, errors.ErrInternalServerError)
	}

	ginhelper.RespondSuccess(c, tokenInfo)
}

// In verifies username and password and generate a token to client.
// @Tags Login
// @Summary Login with username and password.
// @Param user body schema.LoginParam true "login with username and password"
// @Success 200 {object} schema.LoginTokenInfo
// @Failure 400 {object} errors.impl "bad request parameters or invalid username/password"
// @Router /api/v1/pub/login [post]
func (l *Login) In(c *gin.Context) {
	var cred schema.LoginParam
	if err := c.ShouldBind(&cred); err != nil {
		logger.Error(err.Error())
		ginhelper.RespondError(c, errors.ErrBadRequestParam)
		return
	}

	root := l.users.GetRootUser()
	if l.users.VerifyCredential(cred, root) {
		l.respondWithToken(c, root.RecordID)
		return
	}

	results, err := l.users.Query(cred.UserName)
	if err != nil {
		logger.Error(err.Error())
		ginhelper.RespondError(c, errors.ErrInternalServerError)
		return
	}

	if len(results.Users) == 0 {
		ginhelper.RespondError(c, errors.ErrInvalidUsername)
		return
	}

	for _, user := range results.Users {
		if l.users.VerifyCredential(cred, user) {
			l.respondWithToken(c, user.RecordID)
			return
		}
	}

	ginhelper.RespondError(c, errors.ErrInvalidPassword)
}

// Out destroys the token for the login client.
// @Tags Login
// @Summary Logout with a token.
// @Param Authorization header string false "Bearer token"
// @Success 200 {object} errors.impl "ok"
// @Router /api/v1/pub/login/exit [post]
func (l *Login) Out(c *gin.Context) {
	userID := ginhelper.GetUserID(c)
	if userID != "" {
		token := ginhelper.GetToken(c)
		err := l.auth.DestroyToken(token)
		if err != nil {
			// Swallow the error since client is logout.
			logger.Error(err.Error())
		}
	}

	ginhelper.RespondOK(c)
}
