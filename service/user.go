package service

import (
	"github.com/eosnationftw/eosn-base-api/dto"
	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService interface {
	GetUserByAuth0Id(c *gin.Context, auth0Provider, auth0Id string) (*dto.User, error)
}
