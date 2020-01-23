package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yuxiang660/little-bee-server/internal/app/auther"
	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/internal/app/errors"
	"github.com/yuxiang660/little-bee-server/internal/app/ginhelper"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
)

// UserAuthMiddleware verifies user's authentication. 
func UserAuthMiddleware(a auther.Auther, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if t := ginhelper.GetToken(c); t != "" {
			id, err := a.ParseUserID(t)
			if err != nil {
				logger.Error(err.Error())
				ginhelper.RespondError(c, errors.ErrInvalidToken)
				return
			} else if id != "" {
				ginhelper.SetUserID(c, id)
				c.Next()
				return
			}
		}

		if skipHandler(c, skippers...) {
			c.Next()
			return
		}

		cfg := config.Global()
		if cfg.IsDebugMode() {
			ginhelper.SetUserID(c, cfg.Root.UserName)
			c.Next()
			return
		}
		ginhelper.RespondError(c, errors.ErrInvalidToken)
	}	
}

