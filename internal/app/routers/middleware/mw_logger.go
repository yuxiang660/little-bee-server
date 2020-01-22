package middleware

import (
	"bytes"
	"io/ioutil"
	"mime"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuxiang660/little-bee-server/internal/app/ginhelper"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
)

// LoggerMiddleware is a middleware for log.
func LoggerMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func (c *gin.Context) {
		if skipHandler(c, skippers...) {
			c.Next()
			return
		}

		method := c.Request.Method

		fields := make(logger.Fields)

		fields["ip"] = c.ClientIP()
		fields["method"] = method
		fields["url"] = c.Request.URL.String()
		fields["proto"] = c.Request.Proto

		if method == http.MethodPost || method == http.MethodPut {
			mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
			if mediaType == "application/json" {
				body, err := ioutil.ReadAll(c.Request.Body)
				c.Request.Body.Close()
				if err == nil {
					buf := bytes.NewBuffer(body)
					c.Request.Body = ioutil.NopCloser(buf)
					fields["content_length"] = c.Request.ContentLength
					fields["body"] = string(body)
				}
			}
		}

		start := time.Now()
		c.Next()
		timeConsuming := time.Since(start).Nanoseconds() / 1e6
		fields["time_consuming(ms)"] = timeConsuming

		if id := ginhelper.GetUserID(c); id != "" {
			fields["user_id"] = id
		}
		fields["res_status"] = c.Writer.Status()
		fields["res_length"] = c.Writer.Size()
		

		logger.InfoWithFields("API Log", fields)
	}
}
