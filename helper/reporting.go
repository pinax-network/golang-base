package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/pinax-network/golang-base/response"
)

func ReportPublicError(c *gin.Context, apiError *response.ApiError, meta interface{}) {
	reportError(c, apiError, meta, gin.ErrorTypePublic, false)
}

func ReportPublicErrorAndAbort(c *gin.Context, apiError *response.ApiError, meta interface{}) {
	reportError(c, apiError, meta, gin.ErrorTypePublic, true)
}

func ReportPrivateError(c *gin.Context, apiError *response.ApiError, meta interface{}) {
	reportError(c, apiError, meta, gin.ErrorTypePrivate, false)
}

func ReportPrivateErrorAndAbort(c *gin.Context, apiError *response.ApiError, meta interface{}) {
	reportError(c, apiError, meta, gin.ErrorTypePrivate, true)
}

func reportError(c *gin.Context, apiError *response.ApiError, meta interface{}, errType gin.ErrorType, abort bool) {
	_ = c.Error(apiError).SetMeta(meta).SetType(errType)
	if abort {
		c.Abort()
	}
}
