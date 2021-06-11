package service

import (
	"context"
	"github.com/eosnationftw/eosn-base-api/models"
	"github.com/friendsofgo/errors"
)

var (
	ErrLanguageNotFound = errors.New("language not found")
)

type LanguageService interface {
	GetByCode(ctx context.Context, languageCode string) (*base_models.Language, error)
	ListSupportedLanguageCodes(ctx context.Context) []string
}
