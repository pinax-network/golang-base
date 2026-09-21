package middleware

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	base_global "github.com/pinax-network/golang-base/global"
	"github.com/pinax-network/golang-base/helper"
	"github.com/pinax-network/golang-base/log"
	base_models "github.com/pinax-network/golang-base/models"
	"github.com/pinax-network/golang-base/response"
	base_service "github.com/pinax-network/golang-base/service"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type lookupFunc func(context.Context, string) (*base_models.User, *response.ApiError)

func (f lookupFunc) ExtractUserByGUID(ctx context.Context, id string) (*base_models.User, *response.ApiError) {
	return f(ctx, id)
}

func testJWT(t *testing.T, users base_service.UserService) (*JwksMiddleware, *rsa.PrivateKey) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	cert := &x509.Certificate{SerialNumber: big.NewInt(1), NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour)}
	der, err := x509.CreateCertificate(rand.Reader, cert, cert, &key.PublicKey, key)
	require.NoError(t, err)
	data, err := json.Marshal(Jwks{Keys: []JSONWebKeys{{Kid: "test-key", X5c: []string{base64.StdEncoding.EncodeToString(der)}}}})
	require.NoError(t, err)
	path := filepath.Join(t.TempDir(), "jwks.json")
	require.NoError(t, os.WriteFile(path, data, 0600))
	m, err := NewJwksMiddleware(users, &JwtMiddlewareConfig{Auth0Domain: "issuer.example", Auth0AllowedAudiences: []string{"pinax-api"}, Namespace: "https://pinax.example/", JwksFile: path})
	require.NoError(t, err)
	return m, key
}
func serviceClaims() jwt.MapClaims {
	return jwt.MapClaims{"iss": "https://issuer.example/", "aud": "pinax-api", "sub": "MixedCaseClient@clients", "azp": "MixedCaseClient", "gty": "client-credentials", "scope": "admin", "iat": time.Now().Unix(), "exp": time.Now().Add(time.Hour).Unix()}
}
func signTest(t *testing.T, key *rsa.PrivateKey, claims jwt.MapClaims, mutate func(*jwt.Token)) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = "test-key"
	if mutate != nil {
		mutate(token)
	}
	signed, err := token.SignedString(key)
	require.NoError(t, err)
	return signed
}
func requestJWT(router http.Handler, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", "/admin/auth-check", nil)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
func serviceRouter(m *JwksMiddleware, extract bool, clients []ServiceClientConfig, handler gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.Use(Recovery(false), Errors())
	r.GET("/admin/auth-check", m.AuthenticateWithServiceClients(extract, clients), NewAuthMiddleware().CheckPermissions([]string{"admin"}), handler)
	return r
}
func allowedClients() []ServiceClientConfig {
	return []ServiceClientConfig{{ClientID: "MixedCaseClient", PrincipalGUID: "service-guid"}}
}

func TestServiceJWTUsesServerMappingAndPreservesIdentity(t *testing.T) {
	calls := 0
	user := &base_models.User{ID: 42, GUID: "service-guid", Permissions: []string{"old"}}
	m, key := testJWT(t, lookupFunc(func(ctx context.Context, guid string) (*base_models.User, *response.ApiError) {
		calls++
		require.Equal(t, "service-guid", guid)
		c := ctx.(*gin.Context)
		require.Equal(t, "MixedCaseClient@clients", c.GetString(base_global.CONTEXT_AUTH0_FULLID))
		require.Equal(t, "MixedCaseClient", c.GetString(ServiceClientIDContextKey))
		return user, nil
	}))
	claims := serviceClaims()
	claims["https://pinax.example/user_id"] = "attacker-guid"
	r := serviceRouter(m, true, allowedClients(), func(c *gin.Context) {
		u, err := helper.ExtractUserFromContext(c)
		require.NoError(t, err)
		require.Equal(t, 42, u.ID)
		require.Equal(t, []string{"admin"}, u.Permissions)
		c.Status(204)
	})
	token := signTest(t, key, claims, nil)
	require.Equal(t, 204, requestJWT(r, token).Code)
	require.Equal(t, 204, requestJWT(r, token).Code)
	require.Equal(t, 2, calls)
	require.Equal(t, []string{"old"}, user.Permissions, "do not mutate cached user records")
}

func TestServiceJWTRejectsInvalidTokens(t *testing.T) {
	m, key := testJWT(t, nil)
	wrongKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	tests := []struct {
		name        string
		change      func(jwt.MapClaims)
		header      func(*jwt.Token)
		wrongSigner bool
	}{
		{name: "unknown client", change: func(c jwt.MapClaims) { c["sub"] = "unknown@clients"; c["azp"] = "unknown" }},
		{name: "case changed client", change: func(c jwt.MapClaims) { c["sub"] = "mixedcaseclient@clients"; c["azp"] = "mixedcaseclient" }},
		{name: "wrong issuer", change: func(c jwt.MapClaims) { c["iss"] = "https://attacker.example/" }},
		{name: "missing issuer", change: func(c jwt.MapClaims) { delete(c, "iss") }},
		{name: "wrong audience", change: func(c jwt.MapClaims) { c["aud"] = []string{"wrong"} }},
		{name: "missing audience", change: func(c jwt.MapClaims) { delete(c, "aud") }},
		{name: "expired", change: func(c jwt.MapClaims) { c["exp"] = time.Now().Add(-time.Minute).Unix() }},
		{name: "missing expiry", change: func(c jwt.MapClaims) { delete(c, "exp") }},
		{name: "wrong expiry type", change: func(c jwt.MapClaims) { c["exp"] = "tomorrow" }},
		{name: "future nbf", change: func(c jwt.MapClaims) { c["nbf"] = time.Now().Add(time.Hour).Unix() }},
		{name: "no admin scope", change: func(c jwt.MapClaims) { c["scope"] = "read:admin superadmin"; c["permissions"] = []string{"admin"} }},
		{name: "no grant", change: func(c jwt.MapClaims) { delete(c, "gty") }},
		{name: "wrong grant", change: func(c jwt.MapClaims) { c["gty"] = "password" }},
		{name: "human subject with machine grant", change: func(c jwt.MapClaims) { c["sub"] = "auth0|human" }},
		{name: "azp mismatch", change: func(c jwt.MapClaims) { c["azp"] = "wrong" }},
		{name: "missing client claim", change: func(c jwt.MapClaims) { delete(c, "azp") }},
		{name: "conflicting client claims", change: func(c jwt.MapClaims) { c["client_id"] = "wrong" }},
		{name: "missing kid", header: func(token *jwt.Token) { delete(token.Header, "kid") }},
		{name: "wrong kid type", header: func(token *jwt.Token) { token.Header["kid"] = 42 }},
		{name: "unknown kid", header: func(token *jwt.Token) { token.Header["kid"] = "missing" }},
		{name: "wrong signature", wrongSigner: true},
	}
	r := serviceRouter(m, false, allowedClients(), func(c *gin.Context) { c.Status(204) })
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			claims := serviceClaims()
			if tc.change != nil {
				tc.change(claims)
			}
			signer := key
			if tc.wrongSigner {
				signer = wrongKey
			}
			token := signTest(t, signer, claims, tc.header)
			w := requestJWT(r, token)
			require.Contains(t, []int{401, 403}, w.Code)
			require.NotContains(t, w.Body.String(), token)
		})
	}
	anonymous := requestJWT(r, "")
	require.Equal(t, 401, anonymous.Code)
	require.Equal(t, "1", anonymous.Header().Get("X-Pinax-Admin-Service-Auth"))
	require.Equal(t, "no-store", anonymous.Header().Get("Cache-Control"))
	unsigned := jwt.NewWithClaims(jwt.SigningMethodNone, serviceClaims())
	unsigned.Header["kid"] = "test-key"
	token, err := unsigned.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	require.Equal(t, 401, requestJWT(r, token).Code)
}

func TestServiceAccessIsOptInAndRejectsBadConfiguration(t *testing.T) {
	m, key := testJWT(t, nil)
	token := signTest(t, key, serviceClaims(), nil)
	for _, clients := range [][]ServiceClientConfig{nil, {}, {{ClientID: "*", PrincipalGUID: "service-guid"}}, {{ClientID: "MixedCaseClient", PrincipalGUID: ""}}, {{ClientID: "MixedCaseClient", PrincipalGUID: "service-guid"}, {ClientID: "MixedCaseClient", PrincipalGUID: "other"}}, {{ClientID: "MixedCaseClient", PrincipalGUID: "service-guid"}, {ClientID: " ", PrincipalGUID: "other"}}} {
		r := serviceRouter(m, false, clients, func(c *gin.Context) { c.Status(204) })
		require.Equal(t, 403, requestJWT(r, token).Code)
	}
	r := gin.New()
	r.Use(Errors())
	r.GET("/admin/auth-check", m.Authenticate(false, false), func(c *gin.Context) { c.Status(204) })
	require.Equal(t, 403, requestJWT(r, token).Code)
	clients := allowedClients()
	r = serviceRouter(m, false, clients, func(c *gin.Context) { c.Status(204) })
	clients[0].ClientID = "changed"
	require.Equal(t, 204, requestJWT(r, token).Code)
}

func TestServicePrincipalLookupFailsClosed(t *testing.T) {
	for _, name := range []string{"missing", "denied", "wrong mapping", "no local ID"} {
		t.Run(name, func(t *testing.T) {
			m, key := testJWT(t, lookupFunc(func(context.Context, string) (*base_models.User, *response.ApiError) {
				switch name {
				case "missing":
					return nil, nil
				case "denied":
					return nil, response.Forbidden
				case "wrong mapping":
					return &base_models.User{ID: 1, GUID: "other"}, nil
				default:
					return &base_models.User{GUID: "service-guid"}, nil
				}
			}))
			r := serviceRouter(m, true, allowedClients(), func(c *gin.Context) { c.Status(204) })
			require.Equal(t, 403, requestJWT(r, signTest(t, key, serviceClaims(), nil)).Code)
		})
	}
}

func TestLegacyUserAuthenticationIsPreserved(t *testing.T) {
	m, key := testJWT(t, nil)
	claims := serviceClaims()
	delete(claims, "gty")
	claims["sub"] = "auth0|existing-admin"
	claims["https://pinax.example/user_id"] = "existing-guid"
	claims["permissions"] = []string{"admin"}
	r := serviceRouter(m, false, allowedClients(), func(c *gin.Context) { require.Empty(t, c.GetString(ServiceClientIDContextKey)); c.Status(204) })
	require.Equal(t, 204, requestJWT(r, signTest(t, key, claims, nil)).Code)
	for _, value := range []interface{}{[]interface{}{42}, "admin", map[string]string{"scope": "admin"}} {
		claims["permissions"] = value
		require.Equal(t, 401, requestJWT(r, signTest(t, key, claims, nil)).Code)
	}
	delete(claims, "permissions")
	require.Equal(t, 403, requestJWT(r, signTest(t, key, claims, nil)).Code)
	delete(claims, "https://pinax.example/user_id")
	token := signTest(t, key, claims, nil)
	w := requestJWT(r, token)
	require.Equal(t, 401, w.Code)
	require.NotContains(t, w.Body.String(), token)
}

func TestErrorAndPanicLogsExcludeRequestSecrets(t *testing.T) {
	core, logs := observer.New(zap.DebugLevel)
	old := log.ZapLogger
	log.ZapLogger = zap.New(core)
	t.Cleanup(func() { log.ZapLogger = old })
	for _, panics := range []bool{false, true} {
		r := gin.New()
		r.Use(Recovery(false), Errors())
		r.GET("/resource/:id", func(c *gin.Context) {
			if panics {
				panic("synthetic error")
			}
			helper.ReportPrivateErrorAndAbort(c, response.InternalServerError, "synthetic failure")
		})
		req := httptest.NewRequest("GET", "/resource/secret-path?token=secret-query", strings.NewReader("secret-body"))
		req.Header.Set("Authorization", "Bearer secret-access-token")
		req.Header.Set("Cookie", "session=secret-cookie")
		req.Header.Set("X-Api-Key", "secret-api-key")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		require.Equal(t, 500, w.Code)
	}
	require.Len(t, logs.All(), 2)
	for _, entry := range logs.All() {
		data, err := json.Marshal(entry.ContextMap())
		require.NoError(t, err)
		for _, secret := range []string{"secret-path", "secret-query", "secret-body", "secret-access-token", "secret-cookie", "secret-api-key"} {
			require.NotContains(t, string(data), secret)
		}
		require.Contains(t, string(data), "/resource/:id")
	}
}

func TestConcurrentServiceRequestsAndKeyReads(t *testing.T) {
	m, key := testJWT(t, nil)
	token := signTest(t, key, serviceClaims(), nil)
	r := serviceRouter(m, false, allowedClients(), func(c *gin.Context) { c.Status(204) })
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.certHandler.refreshMu.Lock()
			m.certHandler.certs = map[string]*rsa.PublicKey{"test-key": &key.PublicKey}
			m.certHandler.lastRefresh = time.Now()
			m.certHandler.refreshMu.Unlock()
			w := requestJWT(r, token)
			if w.Code != 204 {
				t.Errorf("status %d", w.Code)
			}
		}()
	}
	wg.Wait()
}

func TestProtectedOptionsAlwaysVerifiesTheSignature(t *testing.T) {
	m, key := testJWT(t, nil)
	wrongKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	r := gin.New()
	r.Use(Errors())
	r.OPTIONS("/admin/auth-check", m.AuthenticateWithServiceClients(false, allowedClients()), NewAuthMiddleware().CheckPermissions([]string{"admin"}), func(c *gin.Context) { c.Status(204) })
	for _, signer := range []*rsa.PrivateKey{key, wrongKey} {
		req := httptest.NewRequest("OPTIONS", "/admin/auth-check", nil)
		req.Header.Set("Authorization", "Bearer "+signTest(t, signer, serviceClaims(), nil))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if signer == key {
			require.Equal(t, 204, w.Code)
		} else {
			require.Equal(t, 401, w.Code)
		}
	}
}
