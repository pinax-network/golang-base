package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotBlankValidation(t *testing.T) {
	v := &JsonValidator{}
	err := v.ValidateStruct(struct {
		Foo string `binding:"notblank"`
	}{})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"notblank"`
	}{
		Foo: "not blank",
	})
	assert.NoError(t, err)
}

func TestSortPairValidation(t *testing.T) {
	v := &JsonValidator{}
	err := v.ValidateStruct(struct {
		Foo string `binding:"sortpair"`
	}{})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"sortpair"`
	}{
		Foo: "field:asc",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"sortpair"`
	}{
		Foo: "field:desc",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"sortpair"`
	}{
		Foo: "invalid:",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"sortpair"`
	}{
		Foo: "foo:bar",
	})
	assert.Error(t, err)
}

func TestEosAccountValidation(t *testing.T) {
	v := &JsonValidator{}
	err := v.ValidateStruct(struct {
		Foo string `binding:"eosaccount"`
	}{
		Foo: "eosnationftw",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"eosaccount"`
	}{
		Foo: "",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"eosaccount"`
	}{
		Foo: "too_long_eos_account_name_1234567890",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"eosaccount"`
	}{
		Foo: "invalid#chars",
	})
	assert.Error(t, err)
}

func TestUsernameValidation(t *testing.T) {
	v := &JsonValidator{}
	err := v.ValidateStruct(struct {
		Foo string `binding:"username"`
	}{
		Foo: "johndoe",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"username"`
	}{
		Foo: "too_long_username_1234567890123456789012345678901234567890",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"username"`
	}{
		Foo: "invalid@chars",
	})
	assert.Error(t, err)
}

func TestGithubIssueValidation(t *testing.T) {
	v := &JsonValidator{}
	err := v.ValidateStruct(struct {
		Foo string `binding:"githubissue"`
	}{
		Foo: "https://github.com/org1/repo2/issues/123",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubissue"`
	}{
		Foo: "https://github.com/org/repo/invalid/123",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubissue"`
	}{
		Foo: "http://github.com/org/repo/123",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubissue"`
	}{
		Foo: "github.com/org/repo/123",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubissue"`
	}{
		Foo: "https://github.com/org/repo/issues/abc",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubissue"`
	}{
		Foo: "https://github.com/org/repo/issues/",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubissue"`
	}{
		Foo: "https://github.com/org$$$/repo/issues/123",
	})
	assert.Error(t, err)
}
