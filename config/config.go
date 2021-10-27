package base_config

type ApplicationMode string

const (
	Release ApplicationMode = "release"
	Debug   ApplicationMode = "debug"
	Test    ApplicationMode = "test"
)

type ApplicationConfig struct {
	GinMode  ApplicationMode `yaml:"gin_mode" json:"gin_mode" mapstructure:"gin_mode" validate:"omitempty,oneof=debug release test"` // Gin mode, can be one of develop, release, test
	Domain   string          `yaml:"domain" json:"domain" mapstructure:"domain"`                                                     // The domain inclusive protocol this app is running on, for example http://localhost:8080
	HttpHost string          `yaml:"http_host" json:"http_host" mapstructure:"http_host"`                                            // HTTP host and port to listen on (for applications serving HTTP)
	GrpcHost string          `yaml:"grpc_host" json:"grpc_host" mapstructure:"grpc_host"`                                            // GRPC host and port to listen on (for applications serving GRPC)
}

func (a *ApplicationConfig) IsDebug() bool {
	return a.GinMode == Debug
}

func (a *ApplicationConfig) IsTest() bool {
	return a.GinMode == Test
}

type SmartContractConfig struct {
	Contract        string `yaml:"contract" json:"contract" mapstructure:"contract" validate:"required"`                            // Contract name
	InitialBlockNum int    `yaml:"initial_block_num" json:"initial_block_num" mapstructure:"initial_block_num" validate:"required"` // Block number on which the contract got deployed initially
}

type ApiConfig struct {
	GrpcAddress string `yaml:"grpc_address" json:"grpc_address" mapstructure:"grpc_address" validate:"required"` // Grpc address (host + port)
	HttpAddress string `yaml:"http_address" json:"http_address" mapstructure:"http_address" validate:"required"` // Http address (host + port)
}

type TemporalConfig struct {
	Host      string `yaml:"host" json:"host" mapstructure:"host" validate:"required"` // Temporal host and port without protocol, for example tempo-node12.mar.eosn.io:7233
	Namespace string `yaml:"namespace" json:"namespace" mapstructure:"namespace" validate:"required"`
}

type ChainConfig struct {
	LookupApi string `yaml:"lookup_api" json:"lookup_api" mapstructure:"lookup_api" validate:"required"` // Endpoint for lookups, for example https://eos.eosn.io
}

type AuditConfig struct {
	FileDir       string `yaml:"file_dir" json:"file_dir" mapstructure:"file_dir"`          // Directory to write logs into
	LogToDatabase bool   `yaml:"to_database" json:"to_database" mapstructure:"to_database"` // Set true to log into database
	LogToConsole  bool   `yaml:"to_console" json:"to_console" mapstructure:"to_console"`    // Set true to log to console
}
