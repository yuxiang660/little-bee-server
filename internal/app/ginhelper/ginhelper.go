package ginhelper

import (
	"strings"
	"encoding/json"

	"github.com/yuxiang660/little-bee-server/internal/app/errors"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
	"github.com/gin-gonic/gin"
)

const (
	userIDKey = "user-id"
	resBodyKey = "res-body"
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

// GetUserID returns user id from gin context.
func GetUserID(c *gin.Context) string {
	return c.GetString(userIDKey)
}

// SetUserID sets user id to gin context.
func SetUserID(c *gin.Context, userID string) {
	c.Set(userIDKey, userID)
}

// GetResponseBody returns the contents of the response.
func GetResponseBody(c *gin.Context) string {
	if v, ok := c.Get(resBodyKey); ok && v != nil {
		if b, ok := v.([]byte); ok {
			return string(b)
		}
	}
	return ""
}

// RespondError writes error code and error message with JSON format into the response body.
func RespondError(c *gin.Context, err errors.Error) {
	respondJSON(c, err.Code(), err.Body())
}

// RespondOK writes ok string into the response body with successful status code.
func RespondOK(c *gin.Context) {
	respondJSON(c, errors.NoError.Code(), errors.NoError.Error())
}

// RespondSuccess writes message with JSON format into the response body with successful status code.
func RespondSuccess(c *gin.Context, v interface{}) {
	respondJSON(c, errors.NoError.Code(), v)
}

// respondJSON is the lowest layer of RespondXXX functions, which calls gin API.
// The body of the respond is JSON format.
func respondJSON(c *gin.Context, status int, v interface{}) {
	body, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		logger.Error(err.Error())
		c.JSON(errors.ErrInternalServerError.Code(), errors.ErrInternalServerError.Body())
		return
	}
	c.Set(resBodyKey, body)
	c.Data(status, "application/json; charset=utf-8", body)
}
