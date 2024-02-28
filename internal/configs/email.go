package configs

// EmailConfig - for external email services
type EmailConfig struct {
	Activate     string
	Provider     string
	APIKey       string
	Domain       string
	Host         string
	Port         string
	FromEmail    string
	Password     string
	TrackOpens   bool
	TrackLinks   string
	DeliveryType string

	// for templated email
	EmailVerificationTemplateID int64
	PasswordResetTemplateID     int64
	EmailUpdateVerifyTemplateID int64
	EmailVerificationCodeUUIDv4 bool
	EmailVerificationCodeLength uint64
	PasswordResetCodeUUIDv4     bool
	PasswordResetCodeLength     uint64
	EmailVerificationTag        string
	PasswordResetTag            string
	HTMLModel                   string
	EmailVerifyValidityPeriod   uint64 // in Seconds Default 800Seconds 15 Minute
	PasswordResetValidityPeriod uint64 // in Seconds Default 800Seconds 15 Minute
	NewUserValidityPeriod       uint64 // in Hour Default 24 hour
}
