package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/eosnationftw/eosn-base-api/helper"
	"github.com/eosnationftw/eosn-base-api/log"
	"github.com/eosnationftw/eosn-base-api/response"
	"github.com/eosnationftw/eosn-base-api/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"time"
)

type JwksMiddleware struct {
	userService  base_service.UserService
	jwtValidator *validator.Validator
	config       *JwtMiddlewareConfig
}

func NewJwksMiddleware(userService base_service.UserService, config *JwtMiddlewareConfig) (*JwksMiddleware, error) {

	issuerURL, err := url.Parse(fmt.Sprintf("https://%s/", config.Auth0Domain))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		config.Auth0AllowedAudiences,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to set up the validator: %v", err)
	}

	log.Debug("allowed audiences", zap.Any("allowed", config.Auth0AllowedAudiences))

	return &JwksMiddleware{
		userService:  userService,
		jwtValidator: jwtValidator,
		config:       config,
	}, nil
}

func (j *JwksMiddleware) Authenticate(extractUser, allowAnonymous bool) gin.HandlerFunc {

	return func(c *gin.Context) {

		log.Debug("executing authenticate handler")

		middleware := jwtmiddleware.New(
			j.jwtValidator.ValidateToken,
			jwtmiddleware.WithCredentialsOptional(allowAnonymous),
			jwtmiddleware.WithErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) {

				log.Debug("jwt middleware error handler", zap.Error(err))

				log.Debug("token", zap.Any("token", r.Header.Get("Authorization")))

				switch {
				case errors.Is(err, jwtmiddleware.ErrJWTMissing):
					helper.ReportPublicErrorAndAbort(c, response.NewApiErrorBadRequest(response.BAD_REQUEST_HEADER), err)
					return
				case errors.Is(err, jwtmiddleware.ErrJWTInvalid):
					helper.ReportPublicErrorAndAbort(c, response.Unauthorized, err)
					return
				default:
					helper.ReportPrivateErrorAndAbort(c, response.InternalServerError, err)
					return
				}
			}),
		)

		middleware.CheckJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			claims := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

			log.Debug("token claims", zap.Any("claims", claims))

			payload, err := json.Marshal(claims)
			if err != nil {
				helper.ReportPrivateErrorAndAbort(c, response.InternalServerError, err)
				return
			}

			log.Debug("parsed token claims", zap.Any("payload", payload))
		})).ServeHTTP(c.Writer, c.Request)

		/*
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
			eosnId, ok := claims["https://account.eosnation.io/user_id"].(string)
			if !ok || eosnId == "" {
				helper.ReportPrivateErrorAndAbort(c, response.InternalServerError, fmt.Sprintf("missing claim for the user id ('https://account.eosnation.io/user_id'): %s", tokenString))
				return
			}

			c.Set(base_global.CONTEXT_USER_EOSN_ID, claims["https://account.eosnation.io/user_id"])
			c.Set(base_global.CONTEXT_USER_EMAIL, claims["https://account.eosnation.io/email"])
			c.Set(base_global.CONTEXT_USER_EMAIL_VERIFIED, claims["https://account.eosnation.io/email_verified"])

			c.Set(base_global.CONTEXT_USER_PERMISIONS, claims["permissions"])

			// get the corresponding user from the database if requested
			if extractUser {
				user, apiErr := j.userService.ExtractUserByEosnId(c, eosnId)
				if apiErr != nil {
					helper.ReportPrivateErrorAndAbort(c, apiErr, nil)
					return
				}

				userEmail, ok := claims["https://account.eosnation.io/email"].(string)
				if ok {
					user.Email = userEmail
				}
				userEmailVerified, ok := claims["https://account.eosnation.io/email_verified"].(bool)
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
		*/

		c.Next()
	}

}
