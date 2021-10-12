package consul

type Config struct {
	Host       string `yaml:"host" json:"host" mapstructure:"host" validate:"required"`                   // Consul host inclusive protocol and port, for example http://consul11.mar.eosn.io:8500
	Datacenter string `yaml:"datacenter" json:"datacenter" mapstructure:"datacenter" validate:"required"` // Consul datacenter, for example "march"
	Folder     string `yaml:"folder" json:"folder" mapstructure:"folder" validate:"required"`             // Consul folder for KV pairs, for example "test"
}
