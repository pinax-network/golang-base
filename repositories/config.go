package base_repositories

type UploadRepositoryConfig struct {
	TempUploadDir string `yaml:"temp_upload_dir" json:"temp_upload_dir" mapstructure:"temp_upload_dir" validate:"required"` // Temporary upload directory before it will be uploaded to a file sink.
}
