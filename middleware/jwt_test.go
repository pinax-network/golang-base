package middleware

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
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

func userClaims(subject string) jwt.MapClaims {
	return jwt.MapClaims{
		"iss":                           "https://issuer.example/",
		"aud":                           "pinax-api",
		"sub":                           subject,
		"https://pinax.example/user_id": "existing-guid",
		"https://pinax.example/email":   "admin@pinax.example",
		"permissions":                   []string{"admin"},
		"iat":                           time.Now().Unix(),
		"exp":                           time.Now().Add(time.Hour).Unix(),
	}
}

func machineClaims() jwt.MapClaims {
	return jwt.MapClaims{"iss": "https://issuer.example/", "aud": "pinax-api", "sub": "MixedCaseClient@clients", "azp": "MixedCaseClient", "gty": "client-credentials", "scope": "admin", "permissions": []string{"admin"}, "iat": time.Now().Unix(), "exp": time.Now().Add(time.Hour).Unix()}
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
	req := httptest.NewRequest("GET", "/admin/resource", nil)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// adminRouter protects a route the way services protect their admin groups.
func adminRouter(m *JwksMiddleware, extractUser bool, handler gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.Use(Recovery(false), Errors())
	r.GET("/admin/resource", m.Authenticate(extractUser, false), NewAuthMiddleware().CheckPermissions([]string{"admin"}), handler)
	return r
}

func TestUserJWTSetsIdentityAndLoadsTheUser(t *testing.T) {
	calls := 0
	user := &base_models.User{ID: 42, GUID: "existing-guid", Permissions: []string{"old"}}
	m, key := testJWT(t, lookupFunc(func(ctx context.Context, guid string) (*base_models.User, *response.ApiError) {
		calls++
		require.Equal(t, "existing-guid", guid)
		return user, nil
	}))
	r := adminRouter(m, true, func(c *gin.Context) {
		u, err := helper.ExtractUserFromContext(c)
		require.NoError(t, err)
		require.Equal(t, 42, u.ID)
		require.Equal(t, []string{"admin"}, u.Permissions)
		require.Equal(t, "admin@pinax.example", u.Email)
		require.Equal(t, "auth0|existing-admin", c.GetString(base_global.CONTEXT_AUTH0_FULLID))
		require.Equal(t, "existing-guid", c.GetString(base_global.CONTEXT_USER_GUID))
		c.Status(204)
	})
	token := signTest(t, key, userClaims("auth0|existing-admin"), nil)
	require.Equal(t, 204, requestJWT(r, token).Code)
	require.Equal(t, 204, requestJWT(r, token).Code)
	require.Equal(t, 2, calls)
	require.Equal(t, []string{"old"}, user.Permissions, "do not mutate cached user records")
}

func TestUserLookupFailsClosed(t *testing.T) {
	for name, want := range map[string]int{"missing": 403, "denied": 403, "unavailable": 500} {
		t.Run(name, func(t *testing.T) {
			var users base_service.UserService = lookupFunc(func(context.Context, string) (*base_models.User, *response.ApiError) {
				if name == "denied" {
					return nil, response.Forbidden
				}
				return nil, nil
			})
			if name == "unavailable" {
				users = nil
			}
			m, key := testJWT(t, users)
			r := adminRouter(m, true, func(c *gin.Context) { c.Status(204) })
			require.Equal(t, want, requestJWT(r, signTest(t, key, userClaims("auth0|existing-admin"), nil)).Code)
		})
	}
}

func TestJWTRejectsInvalidTokens(t *testing.T) {
	m, key := testJWT(t, nil)
	wrongKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	tests := []struct {
		name        string
		change      func(jwt.MapClaims)
		header      func(*jwt.Token)
		wrongSigner bool
	}{
		{name: "wrong issuer", change: func(c jwt.MapClaims) { c["iss"] = "https://attacker.example/" }},
		{name: "missing issuer", change: func(c jwt.MapClaims) { delete(c, "iss") }},
		{name: "wrong audience", change: func(c jwt.MapClaims) { c["aud"] = []string{"wrong"} }},
		{name: "missing audience", change: func(c jwt.MapClaims) { delete(c, "aud") }},
		{name: "expired", change: func(c jwt.MapClaims) { c["exp"] = time.Now().Add(-time.Minute).Unix() }},
		{name: "future nbf", change: func(c jwt.MapClaims) { c["nbf"] = time.Now().Add(time.Hour).Unix() }},
		{name: "missing kid", header: func(token *jwt.Token) { delete(token.Header, "kid") }},
		{name: "wrong kid type", header: func(token *jwt.Token) { token.Header["kid"] = 42 }},
		{name: "unknown kid", header: func(token *jwt.Token) { token.Header["kid"] = "missing" }},
		{name: "wrong signature", wrongSigner: true},
	}
	r := adminRouter(m, false, func(c *gin.Context) { c.Status(204) })
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			claims := userClaims("auth0|existing-admin")
			if tc.change != nil {
				tc.change(claims)
			}
			signer := key
			if tc.wrongSigner {
				signer = wrongKey
			}
			token := signTest(t, signer, claims, tc.header)
			w := requestJWT(r, token)
			require.Equal(t, 401, w.Code)
			require.NotContains(t, w.Body.String(), token)
		})
	}
	require.Equal(t, 401, requestJWT(r, "").Code)
	unsigned := jwt.NewWithClaims(jwt.SigningMethodNone, userClaims("auth0|existing-admin"))
	unsigned.Header["kid"] = "test-key"
	token, err := unsigned.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	require.Equal(t, 401, requestJWT(r, token).Code)
}

// Machine tokens have no operator identity, so writes could not be attributed.
func TestMachineTokensAreRejected(t *testing.T) {
	m, key := testJWT(t, lookupFunc(func(context.Context, string) (*base_models.User, *response.ApiError) {
		t.Fatal("machine tokens must not reach the user lookup")
		return nil, nil
	}))
	for name, change := range map[string]func(jwt.MapClaims){
		"client credentials":                 func(jwt.MapClaims) {},
		"clients subject without grant type": func(c jwt.MapClaims) { delete(c, "gty") },
		"grant type with a user subject": func(c jwt.MapClaims) {
			c["sub"] = "auth0|existing-admin"
			c["https://pinax.example/user_id"] = "existing-guid"
		},
	} {
		t.Run(name, func(t *testing.T) {
			claims := machineClaims()
			change(claims)
			token := signTest(t, key, claims, nil)
			for _, extract := range []bool{false, true} {
				w := requestJWT(adminRouter(m, extract, func(c *gin.Context) { c.Status(204) }), token)
				require.Equal(t, 403, w.Code)
				require.NotContains(t, w.Body.String(), token)
			}
		})
	}
}

func TestUserPermissionAndIdentityClaimsAreValidated(t *testing.T) {
	m, key := testJWT(t, nil)
	r := adminRouter(m, false, func(c *gin.Context) { c.Status(204) })
	claims := userClaims("auth0|existing-admin")
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

func TestMultiPartAuth0SubjectsArePreserved(t *testing.T) {
	m, key := testJWT(t, nil)
	cases := map[string]struct{ provider, id string }{
		"auth0|existing-admin":      {"auth0", "existing-admin"},
		"samlp|enterprise|alice":    {"samlp", "enterprise|alice"},
		"auth0|MyConnection1|alice": {"auth0", "MyConnection1|alice"},
	}
	for subject, want := range cases {
		var got struct{ full, provider, id string }
		r := adminRouter(m, false, func(c *gin.Context) {
			got.full = c.GetString(base_global.CONTEXT_AUTH0_FULLID)
			got.provider = c.GetString(base_global.CONTEXT_AUTH0_PROVIDER)
			got.id = c.GetString(base_global.CONTEXT_AUTH0_ID)
			c.Status(204)
		})
		w := requestJWT(r, signTest(t, key, userClaims(subject), nil))
		require.Equal(t, 204, w.Code, subject)
		require.Equal(t, subject, got.full)
		require.Equal(t, want.provider, got.provider)
		require.Equal(t, want.id, got.id)
	}
}

func TestMalformedAuth0SubjectsAreStillRejected(t *testing.T) {
	m, key := testJWT(t, nil)
	r := adminRouter(m, false, func(c *gin.Context) { c.Status(204) })
	for _, subject := range []string{"", "no-separator", "auth0|", "|alice"} {
		require.Equal(t, 401, requestJWT(r, signTest(t, key, userClaims(subject), nil)).Code, "%q", subject)
	}
}

// onDemandJWKS turns a file-backed test middleware into a URL-backed one whose
// fetch is replaced, so unknown signing keys exercise the refresh path.
func onDemandJWKS(m *JwksMiddleware, load func() (map[string]*rsa.PublicKey, error)) {
	m.config.JwksFile = ""
	m.certLoader = load
	m.certHandler.refreshMu.Lock()
	m.certHandler.lastRefresh = time.Now().Add(-2 * onDemandRefreshInterval)
	m.certHandler.refreshMu.Unlock()
}

func concurrently(n int, run func()) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			run()
		}()
	}
	wg.Wait()
}

func TestUnknownSigningKeyRefreshIsCoalesced(t *testing.T) {
	m, key := testJWT(t, nil)
	rotated, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	var fetches atomic.Int32
	onDemandJWKS(m, func() (map[string]*rsa.PublicKey, error) {
		fetches.Add(1)
		time.Sleep(50 * time.Millisecond)
		return map[string]*rsa.PublicKey{"test-key": &key.PublicKey, "rotated-key": &rotated.PublicKey}, nil
	})
	token := signTest(t, rotated, userClaims("auth0|existing-admin"), func(token *jwt.Token) { token.Header["kid"] = "rotated-key" })
	r := adminRouter(m, false, func(c *gin.Context) { c.Status(204) })
	var failures atomic.Int32
	concurrently(40, func() {
		if requestJWT(r, token).Code != 204 {
			failures.Add(1)
		}
	})
	// Every request waits for the single in-flight refresh and then succeeds.
	require.Equal(t, int32(0), failures.Load())
	require.Equal(t, int32(1), fetches.Load())
	// An unknown key right after a successful refresh does not fetch again.
	missing := signTest(t, rotated, userClaims("auth0|existing-admin"), func(token *jwt.Token) { token.Header["kid"] = "missing-key" })
	require.Equal(t, 401, requestJWT(r, missing).Code)
	require.Equal(t, int32(1), fetches.Load())
}

func TestFailedSigningKeyRefreshBacksOff(t *testing.T) {
	m, key := testJWT(t, nil)
	var fetches atomic.Int32
	onDemandJWKS(m, func() (map[string]*rsa.PublicKey, error) {
		fetches.Add(1)
		return nil, errors.New("jwks unavailable")
	})
	token := signTest(t, key, userClaims("auth0|existing-admin"), func(token *jwt.Token) { token.Header["kid"] = "unknown-key" })
	r := adminRouter(m, false, func(c *gin.Context) { c.Status(204) })
	concurrently(40, func() { requestJWT(r, token) })
	require.Equal(t, int32(1), fetches.Load())
	m.certHandler.demandMu.Lock()
	next, failures := m.certHandler.nextOnDemand, m.certHandler.failures
	m.certHandler.demandMu.Unlock()
	require.Equal(t, 1, failures)
	require.WithinDuration(t, time.Now().Add(onDemandRefreshInterval), next, 5*time.Second)
	// Still inside the backoff window: no new outbound request.
	concurrently(40, func() { requestJWT(r, token) })
	require.Equal(t, int32(1), fetches.Load())
	// Once the window passes, a second failure doubles the delay.
	m.certHandler.demandMu.Lock()
	m.certHandler.nextOnDemand = time.Now().Add(-time.Second)
	m.certHandler.demandMu.Unlock()
	requestJWT(r, token)
	require.Equal(t, int32(2), fetches.Load())
	m.certHandler.demandMu.Lock()
	next = m.certHandler.nextOnDemand
	m.certHandler.demandMu.Unlock()
	require.WithinDuration(t, time.Now().Add(2*onDemandRefreshInterval), next, 5*time.Second)
}

func TestConcurrentRequestsAndKeyReads(t *testing.T) {
	m, key := testJWT(t, nil)
	token := signTest(t, key, userClaims("auth0|existing-admin"), nil)
	r := adminRouter(m, false, func(c *gin.Context) { c.Status(204) })
	concurrently(20, func() {
		m.certHandler.refreshMu.Lock()
		m.certHandler.certs = map[string]*rsa.PublicKey{"test-key": &key.PublicKey}
		m.certHandler.lastRefresh = time.Now()
		m.certHandler.refreshMu.Unlock()
		if w := requestJWT(r, token); w.Code != 204 {
			t.Errorf("status %d", w.Code)
		}
	})
}

func TestProtectedOptionsAlwaysVerifiesTheSignature(t *testing.T) {
	m, key := testJWT(t, nil)
	wrongKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	r := gin.New()
	r.Use(Errors())
	r.OPTIONS("/admin/resource", m.Authenticate(false, false), NewAuthMiddleware().CheckPermissions([]string{"admin"}), func(c *gin.Context) { c.Status(204) })
	for _, signer := range []*rsa.PrivateKey{key, wrongKey} {
		req := httptest.NewRequest("OPTIONS", "/admin/resource", nil)
		req.Header.Set("Authorization", "Bearer "+signTest(t, signer, userClaims("auth0|existing-admin"), nil))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if signer == key {
			require.Equal(t, 204, w.Code)
		} else {
			require.Equal(t, 401, w.Code)
		}
	}
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

func TestReverseProxyErrorLogsExcludeRequestSecrets(t *testing.T) {
	core, logs := observer.New(zap.DebugLevel)
	old := log.ZapLogger
	log.ZapLogger = zap.New(core)
	t.Cleanup(func() { log.ZapLogger = old })
	// A closed port makes the proxied request fail and reach ErrorHandler.
	closed := httptest.NewServer(http.NotFoundHandler())
	target := closed.URL
	closed.Close()
	proxy, err := NewReverseProxyMiddleware(target)
	require.NoError(t, err)
	r := gin.New()
	r.Use(Recovery(false), Errors())
	r.GET("/proxy/:id", proxy.ProxyRequest(nil))
	// A real server: ReverseProxy needs a writer with CloseNotify, which the
	// ResponseRecorder behind gin's writer does not implement.
	server := httptest.NewServer(r)
	defer server.Close()
	res, err := http.Get(server.URL + "/proxy/secret-path?token=secret-query")
	require.NoError(t, err)
	require.NoError(t, res.Body.Close())
	require.Equal(t, 500, res.StatusCode)
	require.NotEmpty(t, logs.All())
	for _, entry := range logs.All() {
		data, err := json.Marshal(entry.ContextMap())
		require.NoError(t, err)
		require.NotContains(t, string(data)+entry.Message, "secret-path")
		require.NotContains(t, string(data)+entry.Message, "secret-query")
	}
	require.Contains(t, logs.FilterMessage("failed to reverse proxy request").All()[0].ContextMap()["request"], "/proxy/:id")
}
