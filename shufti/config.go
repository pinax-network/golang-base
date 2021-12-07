package shufti

type Config struct {
	Host        string `yaml:"host" json:"host" mapstructure:"host" validate:"required"`
	ClientId    string `yaml:"client_id" json:"client_id" mapstructure:"client_id" validate:"required"`
	Secret      string `yaml:"secret" json:"secret" mapstructure:"secret" validate:"required"`
	CallbackUrl string `yaml:"callback_url" json:"callback_url" mapstructure:"callback_url"`
	RedirectUrl string `yaml:"redirect_url" json:"redirect_url" mapstructure:"redirect_url" validate:"required"`
}
