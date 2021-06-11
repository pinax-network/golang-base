package service

import (
	"context"
	"github.com/eosnationftw/eosn-base-api/models"
	"github.com/friendsofgo/errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService interface {
	GetUserByAuth0Id(ctx context.Context, auth0Provider, auth0Id string) (*base_models.User, error)
}
