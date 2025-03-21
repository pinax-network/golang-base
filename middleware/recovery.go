package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pinax-network/golang-base/helper"
	"github.com/pinax-network/golang-base/log"
	"github.com/pinax-network/golang-base/response"
	"go.uber.org/zap"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func Recovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					log.Error(c.Request.URL.Path, zap.Any("error", err), zap.String("request", string(httpRequest)))
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				user, _ := helper.ExtractUserFromContext(c)

				if stack {
					log.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.Any("user", user),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					log.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.Any("user", user),
						zap.String("request", string(httpRequest)),
					)
				}

				errInternal := response.InternalServerError
				errResponse := response.ApiErrorResponse{Errors: []*response.ApiError{errInternal}}

				// if we are running in debug mode attach the stack trace to the error response
				if gin.IsDebugging() {
					errResponse.Errors[0].Detail = string(debug.Stack())
				}

				c.AbortWithStatusJSON(errInternal.Status, errResponse)
				return
			}
		}()
		c.Next()
	}
}
