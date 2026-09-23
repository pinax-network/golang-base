package middleware

import (
	"crypto/rsa"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	base_global "github.com/pinax-network/golang-base/global"
	"github.com/pinax-network/golang-base/helper"
	"github.com/pinax-network/golang-base/log"
	"github.com/pinax-network/golang-base/response"
	base_service "github.com/pinax-network/golang-base/service"
	"go.uber.org/zap"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type JwksMiddleware struct {
	userService   base_service.UserService
	jwtMiddleware *jwtmiddleware.JWTMiddleware
	certHandler   *CertHandler
	config        *JwtMiddlewareConfig
	// certLoader replaces the JWKS URL fetch in tests; nil uses loadCerts.
	certLoader func() (map[string]*rsa.PublicKey, error)
}

type CertHandler struct {
	certs       map[string]*rsa.PublicKey
	lastRefresh time.Time
	refreshMu   *sync.Mutex
	// demandMu serializes refreshes triggered by unknown signing keys, so a
	// burst of such tokens shares a single JWKS request. It is held while
	// fetching; refreshMu only guards the key map and timestamps.
	demandMu     sync.Mutex
	nextOnDemand time.Time
	failures     int
}

const (
	// An unknown key can trigger at most one JWKS fetch per interval...
	onDemandRefreshInterval = time.Minute
	// ...and consecutive failures back off exponentially up to this limit.
	onDemandRefreshMaxBackoff = 10 * time.Minute
)

func NewJwksMiddleware(userService base_service.UserService, config *JwtMiddlewareConfig) (*JwksMiddleware, error) {

	j := &JwksMiddleware{
		config:      config,
		userService: userService,
		certHandler: &CertHandler{
			refreshMu: &sync.Mutex{},
		},
	}

	if config.JwksFile != "" {
		err := j.loadCertsFromFile(config.JwksFile)
		if err != nil {
			return nil, err
		}
		log.Info("loaded Auth0 JWKS from file", zap.String("file", config.JwksFile))
	} else {
		err := j.refreshCerts()
		if err != nil {
			return nil, err
		}
		log.Info("loaded Auth0 JWKS from url", zap.String("url", "https://"+j.config.Auth0Domain+"/.well-known/jwks.json"))

		go j.startRefreshCertTimer()
	}

	j.jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		// This middleware protects handlers, including any explicitly registered
		// OPTIONS handler. Public CORS preflights must be handled before it.
		EnableAuthOnOptions: true,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			// we do not write anything to the ResponseWriter here, this will be done in Authenticate()
		},
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			var audiences []string
			switch aud := token.Claims.(jwt.MapClaims)["aud"].(type) {
			case string:
				audiences = []string{aud}
			case []interface{}:
				for _, value := range aud {
					audience, ok := value.(string)
					if !ok {
						return nil, errors.New("invalid audience")
					}
					audiences = append(audiences, audience)
				}
			default:
				return nil, errors.New("invalid audience")
			}
			checkAud := j.verifyAudience(audiences, true)
			if !checkAud {
				return token, errors.New("invalid audience")
			}

			// Verify 'iss' claim
			iss := "https://" + j.config.Auth0Domain + "/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, true)
			if !checkIss {
				return token, errors.New("invalid issuer")
			}

			kid, ok := token.Header["kid"].(string)
			if !ok || kid == "" {
				return nil, errors.New("missing signing key identifier")
			}
			cert, ok, _ := j.signingKey(kid)
			if !ok && j.config.JwksFile == "" {
				// An unknown key usually means Auth0 rotated its signing keys.
				// Refresh (coalesced and rate limited), then look the key up again.
				j.refreshForUnknownKey()
				cert, ok, _ = j.signingKey(kid)
			}
			if !ok {
				return nil, errors.New("unknown signing key")
			}

			return cert, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	return j, nil
}

func (j *JwksMiddleware) Authenticate(extractUser, allowAnonymous bool) gin.HandlerFunc {
	return j.authenticate(extractUser, allowAnonymous, nil)
}

// AuthenticateWithServiceClients explicitly enables administrative machine access
// on the routes where it is installed. It maps exact Auth0 client IDs to
// server-configured service principal GUIDs. User tokens keep the existing path.
// The supplied user service must reject inactive principals and verify that the
// mapped account belongs to this machine identity when extractUser is true.
func (j *JwksMiddleware) AuthenticateWithServiceClients(extractUser bool, clients []ServiceClientConfig) gin.HandlerFunc {
	authenticate := j.authenticate(extractUser, false, copyServiceClients(clients))
	return func(c *gin.Context) {
		// A credential-free probe can distinguish this safe service-auth path
		// from older servers that reject and log machine tokens as user errors.
		c.Header("X-Pinax-Admin-Service-Auth", "1")
		c.Header("Cache-Control", "no-store")
		authenticate(c)
	}
}

func (j *JwksMiddleware) authenticate(extractUser, allowAnonymous bool, clients map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// extract JWT from header
		tokenString, err := jwtmiddleware.FromAuthHeader(c.Request)
		if err != nil {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "invalid authorization header")
			return
		}

		// allow anonymous access only if we don't have a token, otherwise it has to be valid
		if tokenString == "" && allowAnonymous {
			c.Next()
			return
		}

		// validate JWT
		err = j.jwtMiddleware.CheckJWT(c.Writer, c.Request)
		if err != nil {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "invalid access token")
			return
		}

		// extract user information from the token
		claims := jwt.MapClaims{}
		_, _, err = new(jwt.Parser).ParseUnverified(tokenString, &claims)
		if err != nil {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "invalid access token")
			return
		}

		// extract and parse auth0 subject
		subject, ok := claims["sub"].(string)
		if !ok {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "invalid token subject")
			return
		}
		if strings.HasSuffix(subject, "@clients") || claims["gty"] == "client-credentials" {
			j.authenticateService(c, claims, subject, extractUser, clients)
			return
		}

		// Split at the first separator only: SAML and custom-database subjects
		// such as "samlp|connection|user" keep the identity suffix intact.
		extractAuth0 := strings.SplitN(subject, "|", 2)
		if len(extractAuth0) != 2 || extractAuth0[0] == "" || extractAuth0[1] == "" {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "invalid user subject")
			return
		}

		c.Set(base_global.CONTEXT_AUTH0_FULLID, subject)
		c.Set(base_global.CONTEXT_AUTH0_PROVIDER, extractAuth0[0])
		c.Set(base_global.CONTEXT_AUTH0_ID, extractAuth0[1])

		// extract user ID (currently this should always be the EOS Nation ID)
		eosnId, ok := claims[j.getNamespaceClaim("user_id")].(string)
		if !ok || eosnId == "" {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "missing user identifier")
			return
		}

		c.Set(base_global.CONTEXT_USER_GUID, claims[j.getNamespaceClaim("user_id")])
		c.Set(base_global.CONTEXT_USER_EMAIL, claims[j.getNamespaceClaim("email")])
		c.Set(base_global.CONTEXT_USER_EMAIL_VERIFIED, claims[j.getNamespaceClaim("email_verified")])

		permissions, valid := permissionClaims(claims["permissions"])
		if !valid {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "invalid token permissions")
			return
		}
		c.Set(base_global.CONTEXT_USER_PERMISIONS, permissions)

		// set the Github namespaces if available
		if githubId, ok := claims[j.getNamespaceClaim("github_id")]; ok {
			c.Set(base_global.CONTEXT_USER_GITHUB_ID, githubId)

			if githubUsername, ok := claims[j.getNamespaceClaim("github_username")]; ok {
				c.Set(base_global.CONTEXT_USER_GITHUB_USERNAME, githubUsername)
			}
		}

		// get the corresponding user from the database if requested
		if extractUser {
			if j.userService == nil {
				helper.ReportPrivateErrorAndAbort(c, response.InternalServerError, "user lookup unavailable")
				return
			}
			user, apiErr := j.userService.ExtractUserByGUID(c, eosnId)
			if apiErr != nil {
				helper.ReportPrivateErrorAndAbort(c, apiErr, nil)
				return
			}
			if user == nil {
				helper.ReportPublicErrorAndAbort(c, response.Forbidden, nil)
				return
			}
			userCopy := *user
			user = &userCopy

			userEmail, ok := claims[j.getNamespaceClaim("email")].(string)
			if ok {
				user.Email = userEmail
			}
			userEmailVerified, ok := claims[j.getNamespaceClaim("email_verified")].(bool)
			if ok {
				user.EmailVerified = userEmailVerified
			}

			// convert permission list to string array
			if permissions != nil {
				permissionStrings := make([]string, len(permissions))
				for i, p := range permissions {
					permissionStrings[i] = p.(string)
				}
				user.Permissions = permissionStrings
			}

			c.Set(base_global.CONTEXT_USER, user)
		}

		c.Next()
	}
}

func (j *JwksMiddleware) signingKey(kid string) (*rsa.PublicKey, bool, time.Time) {
	j.certHandler.refreshMu.Lock()
	defer j.certHandler.refreshMu.Unlock()
	key, ok := j.certHandler.certs[kid]
	return key, ok, j.certHandler.lastRefresh
}

// refreshForUnknownKey fetches the JWKS at most once per interval however many
// concurrent requests carry unknown signing keys. Waiting callers re-check after
// acquiring the guard, so they reuse a refresh that finished meanwhile, and a
// failed fetch backs off instead of retrying on every unauthenticated request.
func (j *JwksMiddleware) refreshForUnknownKey() {
	h := j.certHandler
	h.demandMu.Lock()
	defer h.demandMu.Unlock()
	_, _, lastRefresh := j.signingKey("")
	now := time.Now()
	if now.Sub(lastRefresh) <= onDemandRefreshInterval || now.Before(h.nextOnDemand) {
		return
	}
	log.Debug("signing key not found, refreshing certs")
	if err := j.refreshCerts(); err != nil {
		h.failures++
		backoff := onDemandRefreshInterval
		for i := 1; i < h.failures && backoff < onDemandRefreshMaxBackoff; i++ {
			backoff *= 2
		}
		if backoff > onDemandRefreshMaxBackoff {
			backoff = onDemandRefreshMaxBackoff
		}
		h.nextOnDemand = now.Add(backoff)
		log.CriticalIfError("failed to reload jwt certs from auth0", err, zap.Duration("retry_after", backoff))
		return
	}
	h.failures = 0
	h.nextOnDemand = time.Time{}
}

func (j *JwksMiddleware) refreshCerts() error {
	load := j.loadCerts
	if j.certLoader != nil {
		load = j.certLoader
	}
	certs, err := load()

	if err != nil {
		return err
	}

	j.certHandler.refreshMu.Lock()
	j.certHandler.certs = certs
	j.certHandler.lastRefresh = time.Now()
	j.certHandler.refreshMu.Unlock()

	return nil
}

func (j *JwksMiddleware) loadCertsFromFile(path string) error {
	certs := make(map[string]*rsa.PublicKey)

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(file).Decode(&jwks)
	if err != nil {
		return err
	}

	if len(jwks.Keys) == 0 {
		return errors.New("no keys available at " + path)
	}

	for _, key := range jwks.Keys {
		if len(key.X5c) == 0 {
			return errors.New("signing key has no certificate")
		}
		certs[key.Kid], err = jwt.ParseRSAPublicKeyFromPEM([]byte("-----BEGIN CERTIFICATE-----\n" + key.X5c[0] + "\n-----END CERTIFICATE-----"))
		log.Debug("loaded cert", zap.String("kid", key.Kid))
		if err != nil {
			return err
		}
	}

	if len(certs) == 0 {
		err := errors.New("unable to find appropriate key")
		return err
	}

	j.certHandler.refreshMu.Lock()
	j.certHandler.certs = certs
	j.certHandler.lastRefresh = time.Now()
	j.certHandler.refreshMu.Unlock()

	return nil
}

func (j *JwksMiddleware) startRefreshCertTimer() {

	ticker := time.NewTicker(10 * time.Minute)

	for {
		select {
		case <-ticker.C:
			err := j.refreshCerts()
			log.CriticalIfError("failed to refresh certs", err)
		}
	}
}

func (j *JwksMiddleware) loadCerts() (map[string]*rsa.PublicKey, error) {
	certs := make(map[string]*rsa.PublicKey)

	certsUrl := "https://" + j.config.Auth0Domain + "/.well-known/jwks.json"
	client := &http.Client{Timeout: 5 * time.Second, CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
		return errors.New("JWKS redirects are not allowed")
	}}
	resp, err := client.Get(certsUrl)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("JWKS request failed")
	}

	var jwks = Jwks{}
	err = json.NewDecoder(io.LimitReader(resp.Body, 1024*1024)).Decode(&jwks)
	if err != nil {
		return nil, err
	}

	if len(jwks.Keys) == 0 {
		return nil, errors.New("no keys available at " + certsUrl)
	}

	for _, key := range jwks.Keys {
		if len(key.X5c) == 0 {
			return nil, errors.New("signing key has no certificate")
		}
		certs[key.Kid], err = jwt.ParseRSAPublicKeyFromPEM([]byte("-----BEGIN CERTIFICATE-----\n" + key.X5c[0] + "\n-----END CERTIFICATE-----"))
		log.Debug("loaded cert", zap.String("kid", key.Kid))

		if err != nil {
			return nil, err
		}
	}

	if len(certs) == 0 {
		err := errors.New("unable to find appropriate key")
		return nil, err
	}

	return certs, nil
}

func (j *JwksMiddleware) verifyAudience(auds []string, req bool) bool {

	if len(auds) == 0 {
		return !req
	}

	for _, allowedAud := range j.config.Auth0AllowedAudiences {
		for _, aud := range auds {
			if subtle.ConstantTimeCompare([]byte(aud), []byte(allowedAud)) != 0 {
				return true
			}
		}
	}

	return false
}

func (j *JwksMiddleware) getNamespaceClaim(claim string) string {

	namespace := j.config.Namespace
	if !strings.HasSuffix(namespace, "/") {
		namespace = namespace + "/"
	}

	return namespace + claim
}
