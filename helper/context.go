package helper

import (
	"fmt"
	"github.com/eosnationftw/eosn-base-api/dto"
	"github.com/eosnationftw/eosn-base-api/global"
	"github.com/gin-gonic/gin"
)

func ExtractLanguageFromContext(c *gin.Context) (language *dto.Language, err error) {

	langInterface, exists := c.Get(global.CONTEXT_LANGUAGE)
	if !exists {
		err = fmt.Errorf("failed to extract language from context, does not exist")
		return
	}

	language, ok := langInterface.(*dto.Language)
	if !ok {
		err = fmt.Errorf("failed to convert language context model")
		return
	}

	return
}

func MustExtractLanguageFromContext(c *gin.Context) *dto.Language {

	language, err := ExtractLanguageFromContext(c)
	if err != nil {
		panic(err)
	}

	return language
}

func ExtractUserFromContext(c *gin.Context) (user *dto.User, err error) {

	userInterface, exists := c.Get(global.CONTEXT_USER)
	if !exists {
		err = fmt.Errorf("failed to extract user from context, does not exist")
		return
	}

	user, ok := userInterface.(*dto.User)
	if !ok {
		err = fmt.Errorf("failed to convert user context model")
		return
	}

	return
}

func MustExtractUserFromContext(c *gin.Context) *dto.User {

	user, err := ExtractUserFromContext(c)
	if err != nil {
		panic(err)
	}

	return user
}
