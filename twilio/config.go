package twilio

type Config struct {
	AccountSID       string `yaml:"account_sid" json:"account_sid" mapstructure:"account_sid"`
	ApiKey           string `yaml:"api_key" json:"api_key" mapstructure:"api_key"`
	ApiSecret        string `yaml:"api_secret" json:"api_secret" mapstructure:"api_secret"`
	VerifyServiceSID string `yaml:"verify_service_sid" json:"verify_service_sid" mapstructure:"verify_service_sid"`
}
