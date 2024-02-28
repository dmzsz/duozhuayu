package utils

import (
	"github.com/dmzsz/duozhuayu/internal/configs"
	"github.com/dmzsz/duozhuayu/internal/datasources/sqldb"

	"github.com/jmoiron/sqlx"
)

func SetupPostgresConnection() (*sqlx.DB, error) {
	// var dsn string
	// switch configs.AppConfig.Environment {
	// case constants.EnvironmentDevelopment:
	// 	dsn = configs.AppConfig.DBPostgreDsn
	// case constants.EnvironmentProduction:
	// 	dsn = config.DBPostgreURL
	// }

	// Setup sqlx config of postgreSQL
	config := sqldb.SQLXConfig{
		DriverName:      configs.AppConfig.DatabaseConfig.RDBMS.Env.Driver,
		DbName:          configs.AppConfig.DatabaseConfig.RDBMS.Access.DbName,
		User:            configs.AppConfig.DatabaseConfig.RDBMS.Access.User,
		Password:        configs.AppConfig.DatabaseConfig.RDBMS.Access.Password,
		MaxOpenConns:    configs.AppConfig.DatabaseConfig.RDBMS.Conn.MaxOpenConns,
		MaxIdleConns:    configs.AppConfig.DatabaseConfig.RDBMS.Conn.MaxIdleConns,
		ConnMaxLifetime: configs.AppConfig.DatabaseConfig.RDBMS.Conn.ConnMaxLifetime,
	}

	// Initialize postgreSQL connection with sqlx
	conn, err := config.InitializeSQLXDatabase()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
