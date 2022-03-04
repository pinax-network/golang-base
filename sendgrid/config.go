package sendgrid

type Config struct {
	Host           string   `yaml:"host" json:"host" mapstructure:"host" validate:"required,url"`             // sendgrid api host, usually https://api.sendgrid.com
	ApiKey         string   `yaml:"api_key" json:"api_key" mapstructure:"api_key" validate:"required"`        // sendgrid api key with appropriate permissions
	ContactListIds []string `yaml:"contact_list_ids" json:"contact_list_ids" mapstructure:"contact_list_ids"` // ids of the contact lists to which users should be added
}
