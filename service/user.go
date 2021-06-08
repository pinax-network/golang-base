package service

import (
	"context"
	"github.com/eosnationftw/eosn-base-api/dto"
	"github.com/friendsofgo/errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService interface {
	GetUserByAuth0Id(ctx context.Context, auth0Provider, auth0Id string) (*dto.User, error)
}
