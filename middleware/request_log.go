package middleware

import "github.com/gin-gonic/gin"

// Log the matched route template, never credentials, headers, query parameters,
// or concrete path parameters (which can themselves contain API keys).
func safeRequestSummary(c *gin.Context) string {
	return c.Request.Method + " " + c.FullPath()
}
