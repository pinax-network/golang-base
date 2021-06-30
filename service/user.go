package service

import (
	"context"
	base_models "github.com/eosnationftw/eosn-base-api/models"
)

type UserService interface {
	// ExtractUserByEosnId assumes the user should be available in the database and must find it by it's EOS Nation ID.
	// This method might panic if no user is found with the given id.
	ExtractUserByEosnId(ctx context.Context, eosnId string) *base_models.User
}
