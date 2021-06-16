package service

import (
	"context"
	base_models "github.com/eosnationftw/eosn-base-api/models"
	"github.com/friendsofgo/errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService interface {
	// deprecated GetUserByAuth0Id(ctx context.Context, auth0Provider, auth0Id string) (*base_models.User, error)
	GetUserByEosnId(ctx context.Context, eosnId string) (*base_models.User, error)
}
