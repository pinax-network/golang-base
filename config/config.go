package config

type ApplicationConfig struct {
	HttpHost string `yaml:"http_host" json:"http_host" mapstructure:"http_host"` // HTTP host and port to listen on (for applications serving HTTP)
	GrpcHost string `yaml:"grpc_host" json:"grpc_host" mapstructure:"grpc_host"` // GRPC host and port to listen on (for applications serving GRPC)
}

type SmartContractConfig struct {
	Contract        string `yaml:"contract" json:"contract" mapstructure:"contract" validate:"required"`                            // Contract name
	InitialBlockNum int    `yaml:"initial_block_num" json:"initial_block_num" mapstructure:"initial_block_num" validate:"required"` // Block number on which the contract got deployed initially
}

type GrpcConfig struct {
	Address string `yaml:"grpc_address" json:"grpc_address" mapstructure:"grpc_address" validate:"required"` // Grpc address (host + port)
}
