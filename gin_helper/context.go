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

// GetUserEmailVerifiedFromContext returns if the user's email has been marked as verified, ErrMissingContextValue if no
// email verification is found, or ErrInvalidContextType if an email verification key is available, but cannot be cast
// into a bool.
func GetUserEmailVerifiedFromContext(ctx context.Context) (emailVerified bool, err error) {
	return extractBoolFromContext(ctx, base_global.CONTEXT_USER_EMAIL_VERIFIED)
}

// GetUserGUIDFromContext returns the user's guid from the context, ErrMissingContextValue if no guid is found, or
// ErrInvalidContextType if a guid key is available, but cannot be cast into a string.
func GetUserGUIDFromContext(ctx context.Context) (guid string, err error) {
	return extractStringFromContext(ctx, base_global.CONTEXT_USER_GUID)
}

// GetUserGithubIdFromContext returns the user's GitHub id, ErrMissingContextValue if no id is found, or
// ErrInvalidContextType if a GitHub id key is available, but cannot be cast into a string.
func GetUserGithubIdFromContext(ctx context.Context) (githubId string, err error) {
	return extractStringFromContext(ctx, base_global.CONTEXT_USER_GITHUB_ID)
}

// GetUserGithubUsernameFromContext returns the user's GitHub username, ErrMissingContextValue if no username is found,
// or ErrInvalidContextType if a GitHub username key is available, but cannot be cast into a string.
func GetUserGithubUsernameFromContext(ctx context.Context) (githubUsername string, err error) {
	return extractStringFromContext(ctx, base_global.CONTEXT_USER_GITHUB_USERNAME)
}

// GetFullAuth0IdFromContext returns the full auth0 id from the context (in the form of '<provider>|<id>'),
// ErrMissingContextValue if no auth0 id is found, or ErrInvalidContextType if an auth0 key is available,
// but cannot be cast into a string.
func GetFullAuth0IdFromContext(ctx context.Context) (githubUsername string, err error) {
	return extractStringFromContext(ctx, base_global.CONTEXT_AUTH0_FULLID)
}

// GetAuth0IdFromContext returns the user's auth0 id from the context (without the provider),
// ErrMissingContextValue if no auth0 id is found, or ErrInvalidContextType if an auth0 id key is available,
// but cannot be cast into a string.
func GetAuth0IdFromContext(ctx context.Context) (githubUsername string, err error) {
	return extractStringFromContext(ctx, base_global.CONTEXT_AUTH0_ID)
}

// GetAuth0ProviderFromContext returns the user's auth0 provider from the context (without the auth0 id),
// ErrMissingContextValue if no auth0 provider is found, or ErrInvalidContextType if an auth0 provider key is available,
// but cannot be cast into a string.
func GetAuth0ProviderFromContext(ctx context.Context) (githubUsername string, err error) {
	return extractStringFromContext(ctx, base_global.CONTEXT_AUTH0_PROVIDER)
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

func extractBoolFromContext(ctx context.Context, key string) (value bool, err error) {

	var ok bool
	if ctx.Value(key) == nil {
		err = ErrMissingContextValue
	} else if value, ok = ctx.Value(key).(bool); !ok {
		err = ErrInvalidContextType
	}

	return
}
