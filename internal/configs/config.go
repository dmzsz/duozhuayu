package configs

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/dmzsz/duozhuayu/internal/constants"
	"github.com/spf13/viper"
)

const Activated string = "yes"

type Config struct {
	// Port        int    `mapstructure:"PORT"`
	// Environment string `mapstructure:"ENVIRONMENT"`
	// Debug       bool   `mapstructure:"DEBUG"`

	// DbActivate        string        `mapstructure:"ACTIVATE_RDBMS"`
	// DbDriver          string        `mapstructure:"DB_DRIVER"`
	// DbHost            string        `mapstructure:"DB_HOST"`
	// DbPort            string        `mapstructure:"DB_PORT"`
	// DbTimeZone        string        `mapstructure:"DB_TIMEZONE"`
	// DbName            string        `mapstructure:"DB_NAME"`
	// DbUser            string        `mapstructure:"DB_USER"`
	// DbPassword        string        `mapstructure:"DB_PASS"`
	// DbSslmode         string        `mapstructure:"DB_SSLMODE"`
	// DbMinTLS          string        `mapstructure:"DB_SSL_MIN_TLS"`
	// DbRootCA          string        `mapstructure:"DB_SSL_ROOT_CA"`
	// DbServerCert      string        `mapstructure:"DB_SSL_SERVER_CERT"`
	// DbClientCert      string        `mapstructure:"DB_SSL_CLIENT_CERT"`
	// DbClientKey       string        `mapstructure:"DB_SSL_CLIENT_KEY"`
	// DbMaxIdleConns    int           `mapstructure:"DB_MAX_IDLE_CONNS"`
	// DbMaxOpenConns    int           `mapstructure:"DB_MAX_OPEN_CONNS"`
	// DbConnMaxLifetime time.Duration `mapstructure:"DB_CONN_MAX_LIFETIME"`
	// DbLogLevel        int           `mapstructure:"DB_LOG_LEVEL"`

	// RedisActivate string `mapstructure:"ACTIVATE_REDIS"`
	// RedisHost     string `mapstructure:"REDIS_HOST"`
	// RedisPort     string `mapstructure:"REDIS_PORT"`
	// RedisPassword string `mapstructure:"REDIS_PASS"`
	// RedisExpired  int    `mapstructure:"REDIS_EXPIRED"`
	// RedisPoolSize int    `mapstructure:"REDIS_POOL_SIZE"`
	// RedisConnTTL  int    `mapstructure:"REDIS_CONN_TTL"`

	// MongoActivate string `mapstructure:"ACTIVATE_MONGO"`
	// MongoAppName  string `mapstructure:"MONGO_APP"`
	// MongoURI      string `mapstructure:"MONGO_URI"`
	// MongoPoolSize uint64 `mapstructure:"MONGO_POOLSIZE"`
	// MongoPoolMon  string `mapstructure:"MONGO_MONITOR_POOL"`
	// MongoConnTTL  int    `mapstructure:"MONGO_CONNTTL"`

	// JWTSecret string `mapstructure:"JWT_SECRET"`
	// JWTExpired int    `mapstructure:"JWT_EXPIRED"`
	// JWTIssuer  string `mapstructure:"JWT_ISSUER"`

	// OTPEmail    string `mapstructure:"OTP_EMAIL"`
	// OTPPassword string `mapstructure:"OTP_PASSWORD"`

	DatabaseConfig DatabaseConfig
	EmailConfig    EmailConfig
	LoggerConfig   LoggerConfig
	MailgunConfig  MailgunConfig
	SecurityConfig SecurityConfig
	ServerConfig   ServerConfig
	ViewConfig     ViewConfig
}

var AppConfig Config

func InitializeAppConfig() error {
	cwd, error := getCallerDir()
	if error != nil {
		return constants.ErrLoadConfig
	}
	fmt.Println(filepath.Join(cwd, "config.yaml"))
	viper.SetConfigFile(filepath.Join(cwd, "config.yaml")) // 指定配置文件路径
	// viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	// viper.SetConfigName("config") // 配置文件名称(无扩展名)
	// viper.SetConfigName(".env")          // allow directly reading from .env file
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("internal/config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	error = viper.ReadInConfig()
	if error != nil {
		return constants.ErrLoadConfig
	}

	error = viper.Unmarshal(&AppConfig)
	// fmt.Println("AppConfig", AppConfig.SecurityConfig.JWT)
	if error != nil {
		return constants.ErrParseConfig
	}
	// fmt.Println(AppConfig.SecurityConfig.JWT)

	// if AppConfig.Port == 0 || AppConfig.Environment == "" || AppConfig.JWTSecret == "" ||
	// 	AppConfig.JWTExpired == 0 || AppConfig.JWTIssuer == "" || AppConfig.OTPEmail == "" ||
	// 	AppConfig.OTPPassword == "" ||
	// 	// (AppConfig.Database.REDIS.Activate == "yes"
	// 	// &&
	// 	// 	(
	// 	// 	AppConfig.Database.REDIS.Env.Host == "" ||
	// 	// 	AppConfig.Database.REDIS.Env.Host == "" ||
	// 	// 	AppConfig.REDISExpired == 0
	// 	// 	)

	// 	// )||
	// 	(AppConfig.Database.RDBMS.Activate == "yes" && AppConfig.Database.RDBMS.Env.Driver == "") {
	// 	return constants.ErrEmptyVar
	// }

	// switch AppConfig.Environment {
	// case constants.EnvironmentDevelopment:
	// 	if AppConfig.DBPostgreDsn == "" {
	// 		return constants.ErrEmptyVar
	// 	}
	// case constants.EnvironmentProduction:
	// 	if AppConfig.DBPostgreURL == "" {
	// 		return constants.ErrEmptyVar
	// 	}
	// }

	return nil
}

func getCallerDir() (string, error) {
	_, filename, _, ok := runtime.Caller(1) // 获取调用栈信息
	if !ok {
		return "", errors.New("failed to get caller information")
	}
	// 获取文件所在的目录
	dir := filepath.Dir(filename)
	return dir, nil
}
