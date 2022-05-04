package elastic

type Config struct {
	Hosts      []string `yaml:"hosts" json:"hosts" mapstructure:"hosts" validate:"required"`
	User       string   `yaml:"user" json:"user" mapstructure:"user"`
	Password   string   `yaml:"password" json:"password" mapstructure:"password"`
	Index      string   `yaml:"index" json:"index" mapstructure:"index" validate:"required"`
	NumWorkers int      `yaml:"num_workers" json:"num_workers" mapstructure:"num_workers" validate:"required"`
}
