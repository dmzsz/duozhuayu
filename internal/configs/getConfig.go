package configs

// IsProd returns true when app is running in production mode
func IsProd() bool {
	return AppConfig.ServerConfig.ServerEnv == "production"
}

// IsSentry returns true when sentry logger is enabled in env config
func IsSentry() bool {
	return AppConfig.LoggerConfig.Activate == Activated
}

// IsBasicAuth returns true when basic auth is enabled in env config
func IsBasicAuth() bool {
	return AppConfig.SecurityConfig.MustBasicAuth == Activated
}

// IsJWT returns true when JWT is enabled in env config
func IsJWT() bool {
	return AppConfig.SecurityConfig.MustJWT == Activated
}

// InvalidateJWT returns true when this feature is enabled in env config
func InvalidateJWT() bool {
	return AppConfig.SecurityConfig.InvalidateJWT == Activated
}

// IsAuthCookie returns true when auth cookie is enabled in env config
func IsAuthCookie() bool {
	return AppConfig.SecurityConfig.AuthCookieActivate
}

// IsHashPass returns true when password hashing is enabled in env config
func IsHashPass() bool {
	return AppConfig.SecurityConfig.MustHash == Activated
}

// IsCipher returns true when encryption at rest is enabled in env config
func IsCipher() bool {
	return AppConfig.SecurityConfig.MustCipher
}

// Is2FA returns true when two-factor authentication is enabled in env config
func Is2FA() bool {
	return AppConfig.SecurityConfig.Must2FA == Activated
}

// Is2FADoubleHash returns true when double hashing is enabled in env config
func Is2FADoubleHash() bool {
	return AppConfig.SecurityConfig.TwoFA.DoubleHash
}

// IsWAF returns true when app firewall is enabled in env config
func IsWAF() bool {
	return AppConfig.SecurityConfig.MustFW == Activated
}

// IsCORS returns true when CORS is enabled in env config
func IsCORS() bool {
	return AppConfig.SecurityConfig.MustCORS == Activated
}

// IsTemplatingEngine returns true when serving HTML is enabled in env config
func IsTemplatingEngine() bool {
	return AppConfig.ViewConfig.Activate == Activated
}

// IsRDBMS returns true when RDBMS is enabled in env config
func IsRDBMS() bool {
	return AppConfig.DatabaseConfig.RDBMS.Activate == Activated
}

// IsRedis returns true when Redis is enabled in env config
func IsRedis() bool {
	return AppConfig.DatabaseConfig.REDIS.Activate == Activated
}

// IsMongo returns true when Mongo is enabled in env config
func IsMongo() bool {
	return AppConfig.DatabaseConfig.MongoDB.Activate == Activated
}

// IsEmailService returns true when email service is enabled in env config
func IsEmailService() bool {
	return AppConfig.EmailConfig.Activate == Activated
}

// IsEmailVerificationService returns true when it is enabled in env config
func IsEmailVerificationService() bool {
	return AppConfig.SecurityConfig.VerifyEmail
}

// IsPassRecoveryService returns true when it is enabled in env config
func IsPassRecoveryService() bool {
	return AppConfig.SecurityConfig.RecoverPass
}

// IsEmailVerificationCodeUUIDv4 returns true when it is enabled in env config
func IsEmailVerificationCodeUUIDv4() bool {
	return AppConfig.EmailConfig.EmailVerificationCodeUUIDv4
}

// IsPasswordRecoverCodeUUIDv4 returns true when it is enabled in env config
func IsPasswordRecoverCodeUUIDv4() bool {
	return AppConfig.EmailConfig.PasswordResetCodeUUIDv4
}
