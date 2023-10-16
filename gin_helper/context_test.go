package gin_helper

import (
	base_global "github.com/eosnationftw/eosn-base-api/global"
	base_models "github.com/eosnationftw/eosn-base-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testUser = &base_models.User{
	ID:            99,
	GUID:          "999",
	Email:         "test@example.org",
	EmailVerified: true,
	Permissions:   []string{"admin"},
	CreatedAt:     time.Now(),
	UpdatedAt:     time.Now(),
}

var testEmail = "test@example.org"
var testGUID = "999"

func TestGetAuthenticatedUserFromContext(t *testing.T) {

	c := &gin.Context{}
	c.Set(base_global.CONTEXT_USER, testUser)

	user, err := GetAuthenticatedUserFromContext(c)
	assert.NoError(t, err)
	assert.Equal(t, testUser, user)
}

func TestGetAuthenticatedUserFromContextNoUserContext(t *testing.T) {

	c := &gin.Context{}

	user, err := GetAuthenticatedUserFromContext(c)
	assert.Error(t, ErrMissingContextValue, err)
	assert.Nil(t, user)
}

func TestGetAuthenticatedUserFromContextInvalidUserContext(t *testing.T) {

	c := &gin.Context{}
	c.Set(base_global.CONTEXT_USER, "this is not a base_model.User struct")

	user, err := GetAuthenticatedUserFromContext(c)
	assert.Error(t, ErrInvalidContextType, err)
	assert.Nil(t, user)
}

func TestMustGetAuthenticatedUserFromContext(t *testing.T) {

	c := &gin.Context{}
	c.Set(base_global.CONTEXT_USER, testUser)

	user := MustGetAuthenticatedUserFromContext(c)
	assert.Equal(t, testUser, user)
}

func TestMustGetAuthenticatedUserFromContextPanics(t *testing.T) {

	c := &gin.Context{}
	assert.Panics(t, func() { MustGetAuthenticatedUserFromContext(c) })
}

func TestGetUserEmailFromContext(t *testing.T) {

	c := &gin.Context{}
	c.Set(base_global.CONTEXT_USER_EMAIL, testEmail)

	email, err := GetUserEmailFromContext(c)
	assert.NoError(t, err)
	assert.Equal(t, testEmail, email)
}

func TestGetUserGUIDFromContext(t *testing.T) {

	c := &gin.Context{}
	c.Set(base_global.CONTEXT_USER_GUID, testGUID)

	guid, err := GetUserGUIDFromContext(c)
	assert.NoError(t, err)
	assert.Equal(t, testGUID, guid)
}
