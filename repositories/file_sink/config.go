package file_sink

type CephFileSinkConfig struct {
	Host      string `yaml:"host" json:"host" mapstructure:"host" validate:"required"`                   // Host including the port, but without protocol. For example ceph-gw.service.march.consul.eosn.io:7480
	Secure    bool   `yaml:"secure" json:"secure" mapstructure:"secure" validate:"required"`             // Whether to use TLS
	Bucket    string `yaml:"bucket" json:"bucket" mapstructure:"bucket" validate:"required"`             // Ceph bucket name
	AccessKey string `yaml:"access_key" json:"access_key" mapstructure:"access_key" validate:"required"` // Ceph access key
	Secret    string `yaml:"secret" json:"secret" mapstructure:"secret" validate:"required"`             // Ceph secret
	BaseUrl   string `yaml:"base_url" json:"base_url" mapstructure:"base_url" validate:"required"`       // Base url from which the uploaded files are being served, for example https://cdn.eosnation.io/test/. This is prefixed to the bucket so the full url will be https://cdn.eosnation.io/test/_bucket_/filename.xyz
}

type LocalFileSinkConfig struct {
	BaseUrl   string `yaml:"base_url" json:"base_url" mapstructure:"base_url" validate:"required"`       // Base url from which the uploaded files are being served, for example https://static.eosnation.io/test/
	UploadDir string `yaml:"upload_dir" json:"upload_dir" mapstructure:"upload_dir" validate:"required"` // File directory to upload static files to, for example /var/www/static
}
