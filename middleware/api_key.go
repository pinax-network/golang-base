package middleware

import (
	"crypto/subtle"
	"errors"
	"github.com/eosnationftw/eosn-base-api/helper"
	"github.com/eosnationftw/eosn-base-api/response"
	"github.com/gin-gonic/gin"
)

type ApiKeyMiddleware struct {
	apiKey       string
	apiKeyHeader string
}

func NewApiKeyMiddleware(apiKey, apiKeyHeader string) (*ApiKeyMiddleware, error) {

	if apiKey == "" || len(apiKey) < 16 {
		return nil, errors.New("invalid or empty api key given")
	}

	if apiKeyHeader == "" {
		return nil, errors.New("api key header is empty")
	}

	return &ApiKeyMiddleware{apiKey: apiKey, apiKeyHeader: apiKeyHeader}, nil
}

func (a *ApiKeyMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.GetHeader(a.apiKeyHeader)

		if apiKey == "" {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "no api key set")
			return
		}

		if subtle.ConstantTimeCompare([]byte(apiKey), []byte(a.apiKey)) != 1 {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "invalid api key given")
			return
		}

		c.Next()
	}
}
