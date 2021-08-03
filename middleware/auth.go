package middleware

import (
	"fmt"
	"github.com/eosnationftw/eosn-base-api/global"
	"github.com/eosnationftw/eosn-base-api/helper"
	"github.com/eosnationftw/eosn-base-api/response"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (a *AuthMiddleware) CheckPermissions(oneOf []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(oneOf) == 0 {
			panic(fmt.Errorf("restricting by permissions needs to set at least one allowed permission that is not empty"))
		}

		permissionInterface, exists := c.Get(base_global.CONTEXT_USER_PERMISIONS)

		if !exists {
			helper.ReportPrivateErrorAndAbort(c, response.Forbidden, nil)
			return
		}

		permissions, ok := permissionInterface.([]interface{})
		if !ok {
			panic(fmt.Errorf("jwt permissions expected to be []interface{}, instead got: '%T', %v", permissionInterface, permissionInterface))
		}

		for _, p := range oneOf {
			if hasPermission(permissions, p) {
				c.Next()
				return
			}
		}

		helper.ReportPrivateErrorAndAbort(c, response.Forbidden, nil)
	}
}

func hasPermission(permissions []interface{}, requiredPermission string) bool {

	// make sure we don't allow any empty permissions
	if len(permissions) == 0 || requiredPermission == "" {
		return false
	}

	for _, p := range permissions {
		if p == requiredPermission {
			return true
		}
	}

	return false
}
