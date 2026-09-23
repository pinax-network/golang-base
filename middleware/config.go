package middleware

type JwtMiddlewareConfig struct {
	Auth0Domain           string   `yaml:"auth0_domain" json:"auth0_domain" mapstructure:"auth0_domain" validate:"required"`                                  // Auth0 domain without protocol and trailing slash, for example login-dev.eosnation.io
	Auth0AllowedAudiences []string `yaml:"auth0_allowed_audiences" json:"auth0_allowed_audiences" mapstructure:"auth0_allowed_audiences" validate:"required"` // Audiences allowed accessing this API
	Namespace             string   `yaml:"namespace" json:"namespace" mapstructure:"namespace" validate:"required"`                                           // The namespace for the token claims
	JwksFile              string   `yaml:"jwks_file" json:"jwks_file" mapstructure:"jwks_file" validate:"omitempty"`
	// Only used by routes explicitly opting into AuthenticateWithServiceClients.
	AdminServiceClients []ServiceClientConfig `yaml:"admin_service_clients" json:"admin_service_clients" mapstructure:"admin_service_clients" validate:"omitempty"`
}

// Keep IDs in values: Viper lowercases map keys, but Auth0 IDs are case-sensitive.
type ServiceClientConfig struct {
	ClientID      string `yaml:"client_id" json:"client_id" mapstructure:"client_id"`
	PrincipalGUID string `yaml:"principal_guid" json:"principal_guid" mapstructure:"principal_guid"`
}
