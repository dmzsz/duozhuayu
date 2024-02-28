package sqldb

import (
	"fmt"
	"time"

	"github.com/dmzsz/duozhuayu/internal/constants"
	"github.com/dmzsz/duozhuayu/pkg/logger"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/sirupsen/logrus"
)

// SQLXConfig holds the configuration for the database instance
type SQLXConfig struct {
	DriverName      string
	DBDriver        string
	Host            string
	Port            string
	TimeZone        string
	DbName          string
	User            string
	Password        string
	Sslmode         string
	MinTLS          string
	RootCA          string
	ServerCert      string
	ClientCert      string
	ClientKey       string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// InitializeSQLXDatabase returns a new DBInstance
func (config *SQLXConfig) InitializeSQLXDatabase() (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	driver := config.DriverName
	user := config.User
	password := config.Password
	dbname := config.DbName
	host := config.Host
	port := config.Port
	sslmode := config.Sslmode
	timeZone := config.TimeZone
	maxIdleConns := config.MaxIdleConns
	maxOpenConns := config.MaxOpenConns
	connMaxLifetime := config.ConnMaxLifetime

	switch driver {
	case "mysql":
		address := "host=" + host
		if port != "" {
			address += " port=" + port
		}
		dsn := address + " user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=" + sslmode + " TimeZone=" + timeZone
		db, err = sqlx.Open(driver, dsn)

		if err != nil {
			return nil, fmt.Errorf("error opening database: %v", err)
		}

		// set maximum number of open connections to database
		logger.Info(fmt.Sprintf("Setting maximum number of open connections to %d", maxOpenConns), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
		db.SetMaxOpenConns(maxOpenConns)

		// set maximum number of idle connections in the pool
		logger.Info(fmt.Sprintf("Setting maximum number of idle connections to %d", maxIdleConns), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
		db.SetMaxIdleConns(maxIdleConns)

		// set maximum time to wait for new connection
		logger.Info(fmt.Sprintf("Setting connect maximum lifetime for a connection to %s", connMaxLifetime), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
		db.SetConnMaxLifetime(connMaxLifetime)

		if err = db.Ping(); err != nil {
			return nil, fmt.Errorf("error pinging database: %v", err)
		}
	case "postgres":
		address := "host=" + config.Host
		if config.Port != "" {
			address += " port=" + config.Port
		}
		dsn := address + " user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=" + sslmode + " TimeZone=" + timeZone

		db, err := sqlx.Open(driver, dsn)
		if err != nil {
			return nil, fmt.Errorf("error opening database: %v", err)
		}

		// set maximum number of open connections to database
		logger.Info(fmt.Sprintf("Setting maximum number of open connections to %d", maxOpenConns), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
		db.SetMaxOpenConns(maxOpenConns)

		// set maximum number of idle connections in the pool
		logger.Info(fmt.Sprintf("Setting maximum number of idle connections to %d", maxIdleConns), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
		db.SetMaxIdleConns(maxIdleConns)

		// set maximum time to wait for new connection
		logger.Info(fmt.Sprintf("Setting connect maximum lifetime for a connection to %s", connMaxLifetime), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
		db.SetConnMaxLifetime(connMaxLifetime)

		if err = db.Ping(); err != nil {
			return nil, fmt.Errorf("error pinging database: %v", err)
		}
	case "sqlite3":
		db, err = sqlx.Open(driver, dbname)
		if err != nil {
			return nil, fmt.Errorf("error opening database: %v", err)
		}

		if err == nil {
			logger.Info(fmt.Sprintf("DB %v connection successful!", driver), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})

		}
	default:
		logger.Fatal(fmt.Sprintf("The driver %v is not implemented yet", driver), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
	}

	return db, nil
}
