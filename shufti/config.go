package shufti

type Config struct {
	Host               string `yaml:"host" json:"host" mapstructure:"host" validate:"required"`
	ClientId           string `yaml:"client_id" json:"client_id" mapstructure:"client_id" validate:"required"`
	Secret             string `yaml:"secret" json:"secret" mapstructure:"secret" validate:"required"`
	CallbackUrl        string `yaml:"callback_url" json:"callback_url" mapstructure:"callback_url"`
	VerificationUrlTtl int    `yaml:"verification_url_ttl" json:"verification_url_ttl" mapstructure:"verification_url_ttl" validate:"required"` // The ttl for the verificaiton url in minutes
}
