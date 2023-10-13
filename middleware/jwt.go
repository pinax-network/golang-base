package middleware

import (
	"crypto/rsa"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/eosnationftw/eosn-base-api/global"
	"github.com/eosnationftw/eosn-base-api/helper"
	"github.com/eosnationftw/eosn-base-api/log"
	"github.com/eosnationftw/eosn-base-api/response"
	"github.com/eosnationftw/eosn-base-api/service"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"sync"
	"time"
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
}

type CertHandler struct {
	certs       map[string]*rsa.PublicKey
	lastRefresh time.Time
	refreshMu   *sync.Mutex
}

func NewJwksMiddleware(userService base_service.UserService, config *JwtMiddlewareConfig) (*JwksMiddleware, error) {

	j := &JwksMiddleware{
		config:      config,
		userService: userService,
		certHandler: &CertHandler{
			refreshMu: &sync.Mutex{},
		},
	}

	err := j.refreshCerts()
	if err != nil {
		return nil, err
	}

	go j.startRefreshCertTimer()

	j.jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			// we do not write anything to the ResponseWriter here, this will be done in Authenticate()
		},
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			var aud []interface{}

			// we need to manually parse the aud array/string from the token, see https://github.com/form3tech-oss/jwt-go/issues/7
			switch token.Claims.(jwt.MapClaims)["aud"].(type) {
			case []interface{}:
				aud = token.Claims.(jwt.MapClaims)["aud"].([]interface{})
				break
			case interface{}:
				aud = make([]interface{}, 0)
				aud = append(aud, token.Claims.(jwt.MapClaims)["aud"].(interface{}))
				break
			default:
				return token, errors.New("invalid audience")
			}

			s := make([]string, len(aud))
			for i, v := range aud {
				s[i] = fmt.Sprint(v)
			}
			token.Claims.(jwt.MapClaims)["aud"] = s

			checkAud := j.verifyAudience(s, true)
			if !checkAud {
				return token, errors.New("invalid audience")
			}

			// Verify 'iss' claim
			iss := "https://" + j.config.Auth0Domain + "/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("invalid issuer")
			}

			cert, ok := j.certHandler.certs[token.Header["kid"].(string)]

			if !ok {
				// occasionally make sure we still have up to date certs if we receive a "kid not found" issue for a token
				// which could have happened due to signing key rotations
				if time.Since(j.certHandler.lastRefresh) > 1*time.Minute {
					log.Debug("kid not found, refreshing certs", zap.Any("kid", token.Header["kid"]))

					err := j.refreshCerts()
					log.CriticalIfError("failed to reload jwt certs from auth0", err)

					if err == nil {
						cert, ok := j.certHandler.certs[token.Header["kid"].(string)]
						if ok {
							return cert, nil
						}
					}
				}
				return token, errors.New("no cert available for kid " + token.Header["kid"].(string))
			}

			return cert, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	return j, nil
}

func (j *JwksMiddleware) Authenticate(extractUser, allowAnonymous bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		// extract JWT from header
		tokenString, err := jwtmiddleware.FromAuthHeader(c.Request)
		if err != nil {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, err)
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
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, err)
			return
		}

		// extract user information from the token
		claims := jwt.MapClaims{}
		_, _, err = new(jwt.Parser).ParseUnverified(tokenString, &claims)
		if err != nil {
			helper.ReportPrivateErrorAndAbort(c, response.InternalServerError, err)
			return
		}

		// extract and parse auth0 subject
		subject, ok := claims["sub"].(string)
		if !ok {
			helper.ReportPrivateErrorAndAbort(c, response.InternalServerError, fmt.Sprintf("jwt subject expected to be string, instead got: '%T', %v", claims["sub"], claims["sub"]))
			return
		}

		extractAuth0 := strings.Split(subject, "|")
		if len(extractAuth0) < 2 {
			helper.ReportPrivateErrorAndAbort(c, response.InternalServerError, fmt.Sprintf("invalid jwt subject given, needs to be of type 'auth_provider|user_id': %s", tokenString))
			return
		}

		c.Set(base_global.CONTEXT_AUTH0_FULLID, subject)
		c.Set(base_global.CONTEXT_AUTH0_PROVIDER, extractAuth0[0])
		c.Set(base_global.CONTEXT_AUTH0_ID, extractAuth0[1])

		// extract user ID (currently this should always be the EOS Nation ID)
		eosnId, ok := claims[j.getNamespaceClaim("user_id")].(string)
		if !ok || eosnId == "" {
			helper.ReportPrivateErrorAndAbort(c, response.InternalServerError, fmt.Sprintf("missing claim for the user id (%q): %s", j.getNamespaceClaim("user_id"), tokenString))
			return
		}

		c.Set(base_global.CONTEXT_USER_EOSN_ID, claims[j.getNamespaceClaim("user_id")])
		c.Set(base_global.CONTEXT_USER_EMAIL, claims[j.getNamespaceClaim("email")])
		c.Set(base_global.CONTEXT_USER_EMAIL_VERIFIED, claims[j.getNamespaceClaim("email_verified")])

		c.Set(base_global.CONTEXT_USER_PERMISIONS, claims["permissions"])

		// get the corresponding user from the database if requested
		if extractUser {
			user, apiErr := j.userService.ExtractUserByGUID(c, eosnId)
			if apiErr != nil {
				helper.ReportPrivateErrorAndAbort(c, apiErr, nil)
				return
			}

			userEmail, ok := claims[j.getNamespaceClaim("email")].(string)
			if ok {
				user.Email = userEmail
			}
			userEmailVerified, ok := claims[j.getNamespaceClaim("email_verified")].(bool)
			if ok {
				user.EmailVerified = userEmailVerified
			}

			// convert permission list to string array
			permissions, ok := claims["permissions"].([]interface{})
			if ok {
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

func (j *JwksMiddleware) refreshCerts() error {
	certs, err := j.loadCerts()

	if err != nil {
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
	resp, err := http.Get(certsUrl)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return nil, err
	}

	if len(jwks.Keys) == 0 {
		return nil, errors.New("no keys available at " + certsUrl)
	}

	for _, key := range jwks.Keys {
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
