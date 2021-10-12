package middleware

type JwtMiddlewareConfig struct {
	Auth0Domain           string   `yaml:"auth0_domain" json:"auth0_domain" mapstructure:"auth0_domain" validate:"required"`                                  // Auth0 domain without protocol and trailing slash, for example login-dev.eosnation.io
	Auth0AllowedAudiences []string `yaml:"auth0_allowed_audiences" json:"auth0_allowed_audiences" mapstructure:"auth0_allowed_audiences" validate:"required"` // Audiences allowed to access this API
}
