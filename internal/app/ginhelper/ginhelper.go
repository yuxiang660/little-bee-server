package ginhelper

import (
	"strings"

	"github.com/yuxiang660/little-bee-server/internal/app/errors"
	"github.com/gin-gonic/gin"
)

const (
	userIDKey = "user-id"
)

// GetToken gets token string from gin context.
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetUserID gets user id from gin context.
func GetUserID(c *gin.Context) string {
	return c.GetString(userIDKey)
}

// SetUserID sets user id to gin context.
func SetUserID(c *gin.Context, userID string) {
	c.Set(userIDKey, userID)
}

// RespondError writes error code and error message with JSON format into the response body.
func RespondError(c *gin.Context, err errors.Error) {
	c.JSON(err.Code(), err.Body())
}
