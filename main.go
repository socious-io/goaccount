package sociousid

// The global config variable holds the configuration for the application.
var config = new(Config)

// Config represents the configuration structure for the payment service.
type Config struct {
	Host   string // Chains represents the blockchain networks supported by the service.
	ID     string // Fiats represents the supported fiat services (e.g., Stripe).
	Secret string // Prefix is used for table name prefix or query prefix (database-related).
}

// Setup initializes the Socious D SDK with the provided configuration.
func Setup(cfg Config) error {

	// Set the global configuration to the provided config.
	config = &cfg
	return nil // Return nil to indicate successful setup.
}
