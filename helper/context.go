package helper

import (
	"fmt"
	"github.com/eosnationftw/eosn-base-api/global"
	"github.com/eosnationftw/eosn-base-api/models"
	"github.com/gin-gonic/gin"
)

func ExtractLanguageFromContext(c *gin.Context) (language *base_models.Language, err error) {

	langInterface, exists := c.Get(base_global.CONTEXT_LANGUAGE)
	if !exists {
		err = fmt.Errorf("failed to extract language from context, does not exist")
		return
	}

	language, ok := langInterface.(*base_models.Language)
	if !ok {
		err = fmt.Errorf("failed to convert language context model")
		return
	}

	return
}

func MustExtractLanguageFromContext(c *gin.Context) *base_models.Language {

	language, err := ExtractLanguageFromContext(c)
	if err != nil {
		panic(err)
	}

	return language
}

func ExtractUserFromContext(c *gin.Context) (user *base_models.User, err error) {

	userInterface, exists := c.Get(base_global.CONTEXT_USER)
	if !exists {
		err = fmt.Errorf("failed to extract user from context, does not exist")
		return
	}

	user, ok := userInterface.(*base_models.User)
	if !ok {
		err = fmt.Errorf("failed to convert user context model")
		return
	}

	return
}

func MustExtractUserFromContext(c *gin.Context) *base_models.User {

	user, err := ExtractUserFromContext(c)
	if err != nil {
		panic(err)
	}

	return user
}

func ExtractFullAuth0IdFromContext(c *gin.Context) (auth0FullId string, err error) {

	userInterface, exists := c.Get(base_global.CONTEXT_USER)
	if !exists {
		err = fmt.Errorf("failed to extract user from context, does not exist")
		return
	}

	auth0FullId, ok := userInterface.(string)
	if !ok {
		err = fmt.Errorf("failed to convert user context model")
		return
	}

	return
}

func MustExtractFullAuth0IdFromContext(c *gin.Context) (auth0FullId string) {

	auth0FullId, err := ExtractFullAuth0IdFromContext(c)
	if err != nil {
		panic(err)
	}

	return
}
