package middleware

import (
	"fmt"
	"github.com/eosnationftw/eosn-base-api/global"
	"github.com/eosnationftw/eosn-base-api/helper"
	"github.com/eosnationftw/eosn-base-api/response"
	"github.com/eosnationftw/eosn-base-api/service"
	"github.com/gin-gonic/gin"
)

type LanguageMiddleware struct {
	languageService service.LanguageService
}

func NewLanguageMiddleware(languageService service.LanguageService) *LanguageMiddleware {
	return &LanguageMiddleware{
		languageService: languageService,
	}
}

func (l *LanguageMiddleware) ParseLanguageHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		languageCode := c.GetHeader("Accept-Language")

		if languageCode == "" {
			helper.ReportPublicErrorAndAbort(c, response.NewApiErrorBadRequest(response.BAD_REQUEST_HEADER_MISSING), "missing required header 'Accept-Language'")
			return
		}

		language, err := l.languageService.GetByCode(c, languageCode)
		if err == service.ErrLanguageNotFound {
			availableCodes := l.languageService.ListSupportedLanguageCodes(c)
			helper.ReportPublicErrorAndAbort(c, response.NewApiErrorBadRequest(response.BAD_REQUEST_HEADER), fmt.Sprintf("invalid 'Accept-Language' header given: '%s', allowed values are: %v", languageCode, availableCodes))
			return
		}

		c.Set(base_global.CONTEXT_LANGUAGE, language)

		c.Next()
	}
}
