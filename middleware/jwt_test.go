package middleware

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	base_global "github.com/pinax-network/golang-base/global"
	"github.com/pinax-network/golang-base/log"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func userClaims(subject string) jwt.MapClaims {
	claims := serviceClaims()
	delete(claims, "gty")
	claims["sub"] = subject
	claims["https://pinax.example/user_id"] = "existing-guid"
	claims["permissions"] = []string{"admin"}
	return claims
}

// Both entry points share the user-subject parser: the plain user middleware
// and the admin routes that also accept service clients.
func userRouters(m *JwksMiddleware, handler gin.HandlerFunc) map[string]*gin.Engine {
	plain := gin.New()
	plain.Use(Recovery(false), Errors())
	plain.GET("/admin/auth-check", m.Authenticate(false, false), handler)
	return map[string]*gin.Engine{
		"Authenticate":                   plain,
		"AuthenticateWithServiceClients": serviceRouter(m, false, allowedClients(), handler),
	}
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
		for name, r := range userRouters(m, func(c *gin.Context) {
			got.full = c.GetString(base_global.CONTEXT_AUTH0_FULLID)
			got.provider = c.GetString(base_global.CONTEXT_AUTH0_PROVIDER)
			got.id = c.GetString(base_global.CONTEXT_AUTH0_ID)
			c.Status(204)
		}) {
			got = struct{ full, provider, id string }{}
			w := requestJWT(r, signTest(t, key, userClaims(subject), nil))
			require.Equal(t, 204, w.Code, "%s %s", name, subject)
			require.Equal(t, subject, got.full, name)
			require.Equal(t, want.provider, got.provider, name)
			require.Equal(t, want.id, got.id, name)
		}
	}
}

func TestMalformedAuth0SubjectsAreStillRejected(t *testing.T) {
	m, key := testJWT(t, nil)
	for _, subject := range []string{"", "no-separator", "auth0|", "|alice"} {
		for name, r := range userRouters(m, func(c *gin.Context) { c.Status(204) }) {
			w := requestJWT(r, signTest(t, key, userClaims(subject), nil))
			require.Equal(t, 401, w.Code, "%s %q", name, subject)
		}
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
	token := signTest(t, rotated, serviceClaims(), func(token *jwt.Token) { token.Header["kid"] = "rotated-key" })
	r := serviceRouter(m, false, allowedClients(), func(c *gin.Context) { c.Status(204) })
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
	missing := signTest(t, rotated, serviceClaims(), func(token *jwt.Token) { token.Header["kid"] = "missing-key" })
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
	token := signTest(t, key, serviceClaims(), func(token *jwt.Token) { token.Header["kid"] = "unknown-key" })
	r := serviceRouter(m, false, allowedClients(), func(c *gin.Context) { c.Status(204) })
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
