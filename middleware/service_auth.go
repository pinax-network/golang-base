package middleware

import (
	"strings"
	"time"
	"unicode"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	base_global "github.com/pinax-network/golang-base/global"
	"github.com/pinax-network/golang-base/helper"
	"github.com/pinax-network/golang-base/response"
)

// ServiceClientIDContextKey identifies the authenticated application separately
// from any human operator attributed by a trusted administrative frontend.
const ServiceClientIDContextKey = "pinax_service_client_id"

func copyServiceClients(clients []ServiceClientConfig) map[string]string {
	copy := make(map[string]string, len(clients))
	for _, config := range clients {
		client, principal := config.ClientID, config.PrincipalGUID
		if _, duplicate := copy[client]; duplicate {
			return nil
		}
		if client == "" || len(client) > 128 || strings.ContainsFunc(client, func(r rune) bool {
			return !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '_' || r == '-')
		}) || principal == "" || len(principal) > 256 || strings.TrimSpace(principal) != principal || strings.ContainsFunc(principal, unicode.IsControl) {
			// A malformed allowlist must never partially enable service access.
			return nil
		}
		copy[client] = principal
	}
	return copy
}

func permissionClaims(value interface{}) ([]interface{}, bool) {
	if value == nil {
		return nil, true
	}
	permissions, ok := value.([]interface{})
	if !ok {
		return nil, false
	}
	for _, p := range permissions {
		if _, ok := p.(string); !ok {
			return nil, false
		}
	}
	return permissions, true
}

// Called only after the normal RS256 signature, issuer, audience and time checks.
func (j *JwksMiddleware) authenticateService(c *gin.Context, claims jwt.MapClaims, subject string, extractUser bool, clients map[string]string) {
	deny := func() { helper.ReportPublicErrorAndAbort(c, response.Forbidden, "service access denied") }
	clientID := strings.TrimSuffix(subject, "@clients")
	principal, allowed := clients[clientID]
	if !allowed || subject != clientID+"@clients" || claims["gty"] != "client-credentials" {
		deny()
		return
	}
	// Auth0's standard profile uses azp; its RFC 9068 profile uses client_id.
	clientClaim := false
	for _, key := range []string{"azp", "client_id"} {
		if value, present := claims[key]; present {
			if id, ok := value.(string); !ok || id != clientID {
				deny()
				return
			}
			clientClaim = true
		}
	}
	expires, validExpiry := claims["exp"].(float64)
	if !clientClaim || !validExpiry || expires <= float64(time.Now().Unix()) {
		deny()
		return
	}
	scope, ok := claims["scope"].(string)
	if !ok || !hasScope(scope, "admin") {
		deny()
		return
	}

	// Do not trust a user_id/custom claim to choose the local service account.
	// Its GUID comes exclusively from the server's exact client allowlist.
	c.Set(ServiceClientIDContextKey, clientID)
	c.Set(base_global.CONTEXT_AUTH0_FULLID, subject)
	c.Set(base_global.CONTEXT_AUTH0_PROVIDER, "client-credentials")
	c.Set(base_global.CONTEXT_AUTH0_ID, clientID)
	c.Set(base_global.CONTEXT_USER_GUID, principal)
	c.Set(base_global.CONTEXT_USER_PERMISIONS, []interface{}{"admin"})
	if extractUser {
		if j.userService == nil {
			deny()
			return
		}
		user, apiErr := j.userService.ExtractUserByGUID(c, principal)
		if apiErr != nil || user == nil || user.ID <= 0 || user.GUID != principal {
			deny()
			return
		}
		copy := *user
		copy.Permissions = []string{"admin"}
		c.Set(base_global.CONTEXT_USER, &copy)
	}
	c.Next()
}

func hasScope(scope, wanted string) bool {
	for _, item := range strings.Fields(scope) {
		if item == wanted {
			return true
		}
	}
	return false
}
