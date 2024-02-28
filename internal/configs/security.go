package configs

import (
	"crypto"
	"net/http"
	// "github.com/dmzsz/duozhuayu/lib"
	// "github.com/dmzsz/duozhuayu/lib/middleware"
)

// SecurityConfig ...
type SecurityConfig struct {
	UserPassMinLength int

	MustBasicAuth string
	BasicAuth     struct {
		Username string
		Password string
	}

	MustJWT string
	JWT     struct {
		Algorithm      string `default:"HS256"`
		AccessKey      string
		AccessKeyTTL   int
		RefreshKey     string
		RefreshKeyTTL  int
		PrivateKeyFile string
		PublicKeyFile  string

		Audience string
		Issuer   string
		AccNbf   int
		RefNbf   int
		Subject  string
	}

	InvalidateJWT string // when user logs off, invalidate the tokens

	AuthCookieActivate bool
	AuthCookiePath     string
	AuthCookieDomain   string
	AuthCookieSecure   bool
	AuthCookieHTTPOnly bool
	AuthCookieSameSite http.SameSite
	ServeJwtAsResBody  bool

	MustHash string
	HashPass HashPassConfig
	HashSec  string // optional secret for argon2id hashing

	// data encryption at rest
	MustCipher bool
	CipherKey  string // for 256-bit ChaCha20-Poly1305
	Blake2bSec string // optional secret for blake2b hashing

	VerifyEmail bool
	RecoverPass bool

	MustFW   string
	Firewall struct {
		ListType string
		IP       string
	}

	MustCORS string
	// CORS     []middleware.CORSPolicy

	TrustedPlatform string

	Must2FA string
	TwoFA   struct {
		Issuer string
		Crypto crypto.Hash
		Digits int

		Status Status2FA
		PathQR string

		DoubleHash bool
	}
}

// Status2FA - user's 2FA statuses
type Status2FA struct {
	Verified string
	On       string
	Off      string
	Invalid  string
}

// HashPassConfig - params for argon2id
type HashPassConfig struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

type JWTParameters struct {
	Algorithm      string `default:"HS256"`
	AccessKey      string
	AccessKeyTTL   int
	RefreshKey     string
	RefreshKeyTTL  int
	PrivateKeyFile string
	PublicKeyFile  string

	Audience string
	Issuer   string
	AccNbf   int
	RefNbf   int
	Subject  string
}
