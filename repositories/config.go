package base_repositories

type UploadRepositoryConfig struct {
	TempUploadDir string `yaml:"temp_upload_dir" json:"temp_upload_dir" mapstructure:"temp_upload_dir" validate:"required"` // Temporary upload directory before it will be uploaded to a file sink.
}

type Auth0Config struct {
	Client           string `yaml:"client" json:"client" mapstructure:"client" validate:"required"`
	Secret           string `yaml:"secret" json:"secret" mapstructure:"secret" validate:"required"`
	ManagementDomain string `yaml:"management_domain" json:"management_domain" mapstructure:"management_domain" validate:"required"`
	ApiKey           string `yaml:"api_key" json:"api_key" mapstructure:"api_key" validate:"required"`
	CallbackUrl      string `yaml:"callback_url" json:"callback_url" mapstructure:"callback_url" validate:"required,url"`
}
