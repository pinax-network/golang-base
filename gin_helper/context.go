package gin_helper

import (
	"context"
	"errors"
	"fmt"
	base_global "github.com/eosnationftw/eosn-base-api/global"
	base_models "github.com/eosnationftw/eosn-base-api/models"
)

var (
	ErrMissingContextValue = errors.New("missing context value for the given key")
	ErrInvalidContextType  = errors.New("context value did not match the expected type")
)

// GetAuthenticatedUserFromContext returns the authenticated user from the context, ErrMissingContextValue if no user is
// found, or ErrInvalidContextType if a user is found but cannot be cast into a base_models.User struct.
func GetAuthenticatedUserFromContext(ctx context.Context) (user *base_models.User, err error) {

	var ok bool
	if ctx.Value(base_global.CONTEXT_USER) == nil {
		err = ErrMissingContextValue
	} else if user, ok = ctx.Value(base_global.CONTEXT_USER).(*base_models.User); !ok {
		err = ErrInvalidContextType
	}

	return
}

// MustGetAuthenticatedUserFromContext returns the authenticated user from the context. It panics if the no user context
// is available, or the context cannot be parsed properly.
func MustGetAuthenticatedUserFromContext(ctx context.Context) (user *base_models.User) {

	user, err := GetAuthenticatedUserFromContext(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to get the user from context: %v", err))
	}

	return
}

// GetUserEmailFromContext returns the user's email from the context, ErrMissingContextValue if no email is found, or
// ErrInvalidContextType if an email key is available, but cannot be cast into a string.
func GetUserEmailFromContext(ctx context.Context) (email string, err error) {
	return extractStringFromContext(ctx, base_global.CONTEXT_USER_EMAIL)
}

// GetUserGUIDFromContext returns the user's guid from the context, ErrMissingContextValue if no guid is found, or
// ErrInvalidContextType if a guid key is available, but cannot be cast into a string.
func GetUserGUIDFromContext(ctx context.Context) (email string, err error) {
	return extractStringFromContext(ctx, base_global.CONTEXT_USER_GUID)
}

func extractStringFromContext(ctx context.Context, key string) (value string, err error) {

	var ok bool
	if ctx.Value(key) == nil {
		err = ErrMissingContextValue
	} else if value, ok = ctx.Value(key).(string); !ok {
		err = ErrInvalidContextType
	}

	return
}
