package middleware

import (
	"errors"
	"github.com/eosnationftw/eosn-base-api/helper"
	"github.com/eosnationftw/eosn-base-api/response"
	"github.com/gin-gonic/gin"
)

type ApiKeyMiddleware struct {
	apiKey string
}

func NewApiKeyMiddleware(apiKey string) (*ApiKeyMiddleware, error) {

	if apiKey == "" || len(apiKey) < 16 {
		return nil, errors.New("invalid or empty api key given")
	}

	return &ApiKeyMiddleware{apiKey: apiKey}, nil
}

func (a *ApiKeyMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.GetHeader("X-API-KEY")

		if apiKey == "" {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "no api key set")
			return
		}
		if apiKey != a.apiKey {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "invalid api key given")
			return
		}

		c.Next()
	}
}
