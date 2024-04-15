package validate

import (
	"testing"

	"github.com/volatiletech/null/v8"

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

func TestSearchPairValidation(t *testing.T) {
	v := &JsonValidator{}
	err := v.ValidateStruct(struct {
		Foo string `binding:"searchpair"`
	}{})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"searchpair"`
	}{
		Foo: "field:test123-124ASAD_PASD",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"searchpair"`
	}{
		Foo: "field:desc",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"searchpair"`
	}{
		Foo: "invalid:",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"searchpair"`
	}{
		Foo: "invalid",
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
		Foo: "https://github.com/org-dash.dot_underscore1/2repo-dash3.dot_underscore4/issues/6",
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

func TestGithubOrgRepoValidation(t *testing.T) {
	v := &JsonValidator{}
	err := v.ValidateStruct(struct {
		Foo string `binding:"githubrepo"`
	}{
		Foo: "https://github.com/org/repo",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubrepo"`
	}{
		Foo: "https://github.com/org-dash.dot_underscore1/2repo-dash3.dot_underscore4",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubrepo"`
	}{
		Foo: "https://github.com/org1",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubrepo"`
	}{
		Foo: "https://github.com/org1/",
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubrepo"`
	}{
		Foo: "https://github.com/org/repo/issues/",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubrepo"`
	}{
		Foo: "http://github.com/org/repo",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubrepo"`
	}{
		Foo: "github.com/org/repo",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubrepo"`
	}{
		Foo: "https://github.com",
	})
	assert.Error(t, err)

	err = v.ValidateStruct(struct {
		Foo string `binding:"githubrepo"`
	}{
		Foo: "https://gitlab.com",
	})
	assert.Error(t, err)
}

func TestNullStringValidation(t *testing.T) {

	v := &JsonValidator{}
	err := v.ValidateStruct(struct {
		Foo null.String `binding:"required,lt=10"`
	}{
		Foo: null.StringFrom("length_ok"),
	})
	assert.NoError(t, err)

	err = v.ValidateStruct(struct {
		Foo null.String `binding:"required,lt=10"`
	}{
		Foo: null.StringFrom("way_to_long_of_a_string"),
	})
	assert.Error(t, err)
}
