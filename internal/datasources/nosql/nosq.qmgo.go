package sqldb

import (
	"context"
	"fmt"
	"time"

	"github.com/dmzsz/duozhuayu/internal/constants"
	"github.com/dmzsz/duozhuayu/pkg/logger"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/event"
	opts "go.mongodb.org/mongo-driver/mongo/options"
)

// NoSQLXConfig holds the configuration for the database instance
type NoSQLXConfig struct {
	Activate string
	AppName  string
	URI      string
	PoolSize uint64
	PoolMon  string
	ConnTTL  int
}

// InitializeNoSQLXDatabase returns a new DBInstance
func (config *NoSQLXConfig) InitializeNoSQLXDatabase() (*qmgo.Client, error) {
	var client *qmgo.Client
	// Connect to the database or cluster
	uri := config.URI

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ConnTTL)*time.Second)
	defer cancel()

	clientConfig := &qmgo.Config{
		Uri:         uri,
		MaxPoolSize: &config.PoolSize,
	}
	serverAPIOptions := opts.ServerAPI(opts.ServerAPIVersion1)

	opt := opts.Client().SetAppName(config.AppName)
	opt.SetServerAPIOptions(serverAPIOptions)

	// for monitoring pool
	if config.PoolMon == "yes" {
		poolMonitor := &event.PoolMonitor{
			Event: func(evt *event.PoolEvent) {
				switch evt.Type {
				case event.GetSucceeded:

					logger.InfoF("DB Mongodb connection successful!", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
				case event.ConnectionReturned:
					logger.FatalF("ConnectionReturned", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
				}
			},
		}
		opt.SetPoolMonitor(poolMonitor)
	}

	client, err := qmgo.NewClient(ctx, clientConfig, options.ClientOptions{ClientOptions: opt})
	if err != nil {
		return client, err
	}

	// Only for debugging
	fmt.Println("MongoDB pool connection successful!")

	return client, nil
}
