package database

type ClusterConfig struct {
	IsGaleraCluster bool                `yaml:"is_galera_cluster" json:"is_galera_cluster" mapstructure:"is_galera_cluster" validate:"required"`
	User            string              `yaml:"user" json:"user" mapstructure:"user" validate:"required"`
	Password        string              `yaml:"password" json:"password" mapstructure:"password" validate:"required"`
	Database        string              `yaml:"database" json:"database" mapstructure:"database" validate:"required"`
	Connections     []*ConnectionConfig `yaml:"connections" json:"connections" mapstructure:"connections" validate:"gte=1,dive,required"`
}

type ConnectionConfig struct {
	Host string `yaml:"host" json:"host" mapstructure:"host" validate:"required"`
	Port int    `yaml:"port" json:"port" mapstructure:"port" validate:"required"`
}

type MigrationConfig struct {
	MigrationDir string `yaml:"migration_dir" json:"migration_dir" mapstructure:"migration_dir" validate:"required"` // Directory for SQL migration files
}
