package base_service

import (
	"context"
	base_models "github.com/eosnationftw/eosn-base-api/models"
	"github.com/eosnationftw/eosn-base-api/response"
)

type UserService interface {
	// ExtractUserByEosnId should extract the user's base model from the data source by the EOS Nation ID.
	// This method might panic if no user is found, create a new one or return the response.ApiError in case
	// a specific error should be returned when trying to extract that user from the authentication middleware.
	ExtractUserByEosnId(ctx context.Context, eosnId string) (*base_models.User, *response.ApiError)
}
