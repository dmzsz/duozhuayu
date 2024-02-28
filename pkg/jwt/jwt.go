package jwt

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dmzsz/duozhuayu/internal/configs"
	V1Domains "github.com/dmzsz/duozhuayu/internal/domains/v1"
	"github.com/dmzsz/duozhuayu/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var jwtConfig JWTConfig

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

// JWTConfig - params to configure JWT
type JWTConfig struct {
	Algorithm     string
	AccessKey     []byte
	AccessKeyTTL  int
	RefreshKey    []byte
	RefreshKeyTTL int
	PrivKeyECDSA  *ecdsa.PrivateKey
	PubKeyECDSA   *ecdsa.PublicKey
	PrivKeyEdDSA  crypto.PrivateKey
	PubKeyEdDSA   crypto.PublicKey
	PrivKeyRSA    *rsa.PrivateKey
	PubKeyRSA     *rsa.PublicKey

	Audience string
	Issuer   string
	AccNbf   int
	RefNbf   int
	Subject  string
}

type JWTClaims struct {
	UserId    string
	Username  string
	Email     string
	RoleIds   []string
	TokenType TokenType

	jwt.RegisteredClaims
}

func init() {
	fmt.Println("jwt init", configs.InitializeAppConfig())
	NewJWT()
}
func NewJWT() *JWTConfig {
	params, err := getParamsJWT()
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{"getParamsJWT": err})
	}
	jwtConfig = params
	return &params
}

// getParamsJWT - read parameters from env
func getParamsJWT() (params JWTConfig, err error) {
	alg := configs.AppConfig.SecurityConfig.JWT.Algorithm
	if alg == "" {
		alg = "HS256" // default algorithm
	}
	// list of accepted algorithms
	// HS256: HMAC-SHA256
	// HS384: HMAC-SHA384
	// HS512: HMAC-SHA512
	// EdDSA: EdDSA Signature with SHA-512
	// ES256: ECDSA Signature with SHA-256
	// ES384: ECDSA Signature with SHA-384
	// ES512: ECDSA Signature with SHA-512
	// PS256: PS256 Signature with SHA-256
	// PS384: PS256 Signature with SHA-384
	// PS512: PS256 Signature with SHA-512
	// RS256: RSA Signature with SHA-256
	// RS384: RSA Signature with SHA-384
	// RS512: RSA Signature with SHA-512
	if alg != "HS256" && alg != "HS384" && alg != "HS512" &&
		alg != "EdDSA" &&
		alg != "ES256" && alg != "ES384" && alg != "ES512" &&
		alg != "PS256" && alg != "PS384" && alg != "PS512" &&
		alg != "RS256" && alg != "RS384" && alg != "RS512" {
		err = errors.New("unsupported algorithm for JWT")
		return
	}
	params.Algorithm = alg
	params.AccessKey = []byte(configs.AppConfig.SecurityConfig.JWT.AccessKey)
	params.AccessKeyTTL = configs.AppConfig.SecurityConfig.JWT.AccessKeyTTL
	params.RefreshKey = []byte(configs.AppConfig.SecurityConfig.JWT.RefreshKey)
	params.RefreshKeyTTL = configs.AppConfig.SecurityConfig.JWT.RefreshKeyTTL

	privateKeyFile := configs.AppConfig.SecurityConfig.JWT.PrivateKeyFile
	if privateKeyFile != "" {
		// load the private key
		privateKeyBytes, err := os.ReadFile(privateKeyFile)
		if err != nil {
			return params, err
		}

		// ECDSA
		if alg == "ES256" || alg == "ES384" || alg == "ES512" {
			privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
			if err != nil {
				return params, err
			}
			params.PrivKeyECDSA = privateKey
		}

		// EdDSA
		if alg == "EdDSA" {
			privateKey, err := jwt.ParseEdPrivateKeyFromPEM(privateKeyBytes)
			if err != nil {
				return params, err
			}
			params.PrivKeyEdDSA = privateKey
		}

		// RSA
		if alg == "RS256" || alg == "RS384" || alg == "RS512" {
			privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
			if err != nil {
				return params, err
			}
			params.PrivKeyRSA = privateKey
		}
	}

	publicKeyFile := configs.AppConfig.SecurityConfig.JWT.PublicKeyFile
	if publicKeyFile != "" {
		// load the public key
		publicKeyBytes, err := os.ReadFile(publicKeyFile)
		if err != nil {
			return params, err
		}

		// ECDSA
		if alg == "ES256" || alg == "ES384" || alg == "ES512" {
			publicKey, err := jwt.ParseECPublicKeyFromPEM(publicKeyBytes)
			if err != nil {
				return params, err
			}
			params.PubKeyECDSA = publicKey
		}

		// EdDSA
		if alg == "EdDSA" {
			publicKey, err := jwt.ParseEdPublicKeyFromPEM(publicKeyBytes)
			if err != nil {
				return params, err
			}
			params.PubKeyEdDSA = publicKey
		}

		// RSA_PSS or RSA
		if alg == "PS256" || alg == "PS384" || alg == "PS512" ||
			alg == "RS256" || alg == "RS384" || alg == "RS512" {
			publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
			if err != nil {
				return params, err
			}
			params.PubKeyRSA = publicKey
		}
	}

	params.Audience = configs.AppConfig.SecurityConfig.JWT.Audience
	params.Issuer = configs.AppConfig.SecurityConfig.JWT.Issuer
	params.AccNbf = configs.AppConfig.SecurityConfig.JWT.AccNbf
	params.RefNbf = configs.AppConfig.SecurityConfig.JWT.RefNbf
	params.Subject = configs.AppConfig.SecurityConfig.JWT.Subject

	return params, nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {

	// 验证签名方法是否为 ECDSA
	if _, ok := token.Method.(*jwt.SigningMethodECDSA); ok {
		return jwtConfig.PubKeyECDSA, nil
	}
	// 验证签名方法是否为 Ed25519
	if _, ok := token.Method.(*jwt.SigningMethodEd25519); ok {
		return jwtConfig.PubKeyEdDSA, nil
	}
	// 验证签名方法是否为 HMAC
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
		if jwtClaims, ok := token.Claims.(*JWTClaims); ok {
			// 获取自定义声明中的字段信息
			tokenType := jwtClaims.TokenType
			if tokenType == AccessToken {
				return jwtConfig.AccessKey, nil
			}
			if tokenType == RefreshToken {
				return jwtConfig.RefreshKey, nil
			}
		}
	}

	// 验证签名方法是否为 RSA-PSS
	if _, ok := token.Method.(*jwt.SigningMethodRSAPSS); ok {
		return jwtConfig.PrivKeyRSA, nil
	}

	// 验证签名方法是否为 RSA
	if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
		return jwtConfig.PubKeyRSA, nil
	}

	// 若签名方法不在预期范围内，则返回错误
	return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
}

func ParseToken(tokenString string) (claims JWTClaims, err error) {

	token, err := jwt.ParseWithClaims(tokenString, &claims, keyFunc)

	if err != nil {
		log.Fatal(err)
		return JWTClaims{}, errors.New("token is not valid")
	} else if claims, ok := token.Claims.(*JWTClaims); ok {
		return *claims, nil
	} else {
		log.Fatal("unknown claims type, cannot proceed")
		return JWTClaims{}, errors.New("unknown claims type, cannot proceed")
	}
}

func GenerateToken(userId string, username string, email string, roles []V1Domains.RoleDomain, tokenType TokenType) (jwtValue string, err error) {

	var (
		secretKey    []byte
		privKeyECDSA *ecdsa.PrivateKey
		privKeyEdDSA crypto.PrivateKey
		privKeyRSA   *rsa.PrivateKey
		ttl          time.Time
		nbf          int
	)

	var jwtRoles []string

	for _, role := range roles {
		jwtRoles = append(jwtRoles, role.Id)
	}

	if tokenType == AccessToken {
		secretKey = jwtConfig.AccessKey
		ttl = time.Now().Add(time.Hour * time.Duration(jwtConfig.AccessKeyTTL))
		nbf = jwtConfig.AccNbf
	}
	if tokenType == RefreshToken {
		secretKey = jwtConfig.RefreshKey
		ttl = time.Now().Add(time.Hour * 24 * time.Duration(jwtConfig.RefreshKeyTTL))
		nbf = jwtConfig.RefNbf
	}

	claims := &JWTClaims{
		UserId:    userId,
		Username:  username,
		Email:     email,
		RoleIds:   jwtRoles,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ttl),
			ID:        uuid.NewString(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    jwtConfig.Issuer,
			Subject:   jwtConfig.Subject,
		},
	}

	if jwtConfig.Audience != "" {
		claims.Audience = []string{jwtConfig.Audience}
	}
	if nbf > 0 {
		claims.NotBefore = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(nbf)))
	}

	var token *jwt.Token
	alg := jwtConfig.Algorithm

	switch alg {
	case "HS256":
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	case "HS384":
		token = jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	case "HS512":
		token = jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	case "EdDSA":
		privKeyEdDSA = jwtConfig.PrivKeyEdDSA
		token = jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	case "ES256":
		privKeyECDSA = jwtConfig.PrivKeyECDSA
		token = jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	case "ES384":
		privKeyECDSA = jwtConfig.PrivKeyECDSA
		token = jwt.NewWithClaims(jwt.SigningMethodES384, claims)
	case "ES512":
		privKeyECDSA = jwtConfig.PrivKeyECDSA
		token = jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	case "PS256":
		privKeyRSA = jwtConfig.PrivKeyRSA
		token = jwt.NewWithClaims(jwt.SigningMethodPS256, claims)
	case "PS384":
		privKeyRSA = jwtConfig.PrivKeyRSA
		token = jwt.NewWithClaims(jwt.SigningMethodPS384, claims)
	case "PS512":
		privKeyRSA = jwtConfig.PrivKeyRSA
		token = jwt.NewWithClaims(jwt.SigningMethodPS512, claims)
	case "RS256":
		privKeyRSA = jwtConfig.PrivKeyRSA
		token = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	case "RS384":
		privKeyRSA = jwtConfig.PrivKeyRSA
		token = jwt.NewWithClaims(jwt.SigningMethodRS384, claims)
	case "RS512":
		privKeyRSA = jwtConfig.PrivKeyRSA
		token = jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	default:
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	}

	// HMAC
	if alg == "HS256" || alg == "HS384" || alg == "HS512" {
		jwtValue, err = token.SignedString(secretKey)
		if err != nil {
			return
		}
	}

	// EdDSA
	if alg == "EdDSA" {
		jwtValue, err = token.SignedString(privKeyEdDSA)
		if err != nil {
			return
		}
	}
	// ECDSA
	//
	// ES256
	// prime256v1: X9.62/SECG curve over a 256 bit prime field, also known as P-256 or NIST P-256
	// widely used, recommended for general-purpose cryptographic operations
	// openssl ecparam -name prime256v1 -genkey -noout -out private-key.pem
	// openssl ec -in private-key.pem -pubout -out public-key.pem
	//
	// ES384
	// secp384r1: NIST/SECG curve over a 384 bit prime field
	// openssl ecparam -name secp384r1 -genkey -noout -out private-key.pem
	// openssl ec -in private-key.pem -pubout -out public-key.pem
	//
	// ES512
	// secp521r1: NIST/SECG curve over a 521 bit prime field
	// openssl ecparam -name secp521r1 -genkey -noout -out private-key.pem
	// openssl ec -in private-key.pem -pubout -out public-key.pem
	if alg == "ES256" || alg == "ES384" || alg == "ES512" {
		jwtValue, err = token.SignedString(privKeyECDSA)
		if err != nil {
			return
		}
	}

	// RSA-PSS
	// openssl genpkey -algorithm RSA-PSS -out private-key.pem -pkeyopt rsa_keygen_bits:2048
	// openssl rsa -in private-key.pem -pubout -out public-key.pem
	//
	// RS384
	// openssl genpkey -algorithm RSA-PSS -out private-key.pem -pkeyopt rsa_keygen_bits:3072
	// openssl rsa -in private-key.pem -pubout -out public-key.pem
	//
	// RS512
	// openssl genpkey -algorithm RSA-PSS -out private-key.pem -pkeyopt rsa_keygen_bits:4096
	// openssl rsa -in private-key.pem -pubout -out public-key.pem
	if alg == "PS256" || alg == "PS384" || alg == "PS512" {
		jwtValue, err = token.SignedString(privKeyRSA)
		if err != nil {
			return
		}
	}

	// RSA
	//
	// RS256
	// openssl genpkey -algorithm RSA -out private-key.pem -pkeyopt rsa_keygen_bits:2048
	// openssl rsa -in private-key.pem -pubout -out public-key.pem
	//
	// RS384
	// openssl genpkey -algorithm RSA -out private-key.pem -pkeyopt rsa_keygen_bits:3072
	// openssl rsa -in private-key.pem -pubout -out public-key.pem
	//
	// RS512
	// openssl genpkey -algorithm RSA -out private-key.pem -pkeyopt rsa_keygen_bits:4096
	// openssl rsa -in private-key.pem -pubout -out public-key.pem
	if alg == "RS256" || alg == "RS384" || alg == "RS512" {
		jwtValue, err = token.SignedString(privKeyRSA)
		if err != nil {
			return
		}
	}
	return
}
