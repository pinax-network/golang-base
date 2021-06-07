package service

import (
	"github.com/eosnationftw/eosn-base-api/dto"
	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrLanguageNotFound = errors.New("language not found")
)

type LanguageService interface {
	GetByCode(c *gin.Context, languageCode string) (*dto.Language, error)
	ListSupportedLanguageCodes(c *gin.Context) []string
}
