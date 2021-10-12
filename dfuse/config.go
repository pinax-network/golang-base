package dfuse

type Config struct {
	Endpoint string `yaml:"endpoint" json:"endpoint" mapstructure:"endpoint" validate:"required"` // Endpoint including the port, but without protocol. For example eos.dfuse.eosnation.io:9000
	Insecure bool   `yaml:"insecure" json:"insecure" mapstructure:"insecure" validate:"required"` // Whether to use TLS encryption.
	ApiKey   string `yaml:"api_key" json:"api_key" mapstructure:"api_key"`                        // Dfuse API key (optional)
}
