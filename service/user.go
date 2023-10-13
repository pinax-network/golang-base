package base_service

import (
	"context"
	base_models "github.com/eosnationftw/eosn-base-api/models"
	"github.com/eosnationftw/eosn-base-api/response"
)

type UserService interface {
	// ExtractUserByGUID should extract the user's base model from the local data source using the user's unique global
	// ID (for example, the eosn_id or uuid).
	// This method might panic if no user is found, create a new one or return the response.ApiError in case
	// a specific error should be returned when trying to extract that user from the authentication middleware.
	ExtractUserByGUID(ctx context.Context, guid string) (*base_models.User, *response.ApiError)
}
