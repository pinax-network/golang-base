package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	base_global "github.com/pinax-network/golang-base/global"
	"github.com/pinax-network/golang-base/helper"
	"github.com/pinax-network/golang-base/response"
	base_service "github.com/pinax-network/golang-base/service"
)

type LanguageMiddleware struct {
	languageService base_service.LanguageService
}

func NewLanguageMiddleware(languageService base_service.LanguageService) *LanguageMiddleware {
	return &LanguageMiddleware{
		languageService: languageService,
	}
}

func (l *LanguageMiddleware) ParseLanguageHeader(defaultLanguageCode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		languageCode := c.GetHeader("X-Accept-Language")

		if languageCode == "" {
			languageCode = defaultLanguageCode
		}

		language, err := l.languageService.GetByCode(c, languageCode)
		if err == base_service.ErrLanguageNotFound {
			availableCodes := l.languageService.ListSupportedLanguageCodes(c)
			helper.ReportPublicErrorAndAbort(c, response.NewApiErrorBadRequest(response.BAD_REQUEST_HEADER), fmt.Sprintf("invalid 'X-Accept-Language' header given: '%s', allowed values are: %v", languageCode, availableCodes))
			return
		}

		c.Set(base_global.CONTEXT_LANGUAGE, language)

		c.Next()
	}
}
