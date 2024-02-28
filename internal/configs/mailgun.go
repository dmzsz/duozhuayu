package configs

// MailgunConfig object
type MailgunConfig struct {
	APIKey string
	// PublicAPIKey string `env:"MAILGUN_PUBLIC_KEY"`
	Domain string
}
