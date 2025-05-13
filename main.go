package goaccount

// The global config variable holds the configuration for the application.
var config = new(Config)

// Config represents the configuration structure for the Socious ID SSO service.
type Config struct {
	Host   string `json:"host" mapstructure:"host"`
	ID     string `json:"id" mapstructure:"id"`
	Secret string `json:"secret" mapstructure:"secret"`
	Jwt    struct {
		Secret   string `json:"secret" mapstructure:"secret"`
		Duration int    `json:"duration" mapstructure:"duration"`
	} `json:"jwt" mapstructure:"jwt"`
}

// Setup initializes the Socious ID SDK with the provided configuration.
func Setup(cfg Config) error {

	// Set the global configuration to the provided config.
	config = &cfg
	return nil // Return nil to indicate successful setup.
}
