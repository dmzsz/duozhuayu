package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dmzsz/duozhuayu/internal/configs"
	"github.com/dmzsz/duozhuayu/internal/constants"
	"github.com/dmzsz/duozhuayu/internal/datasources/caches"
	"github.com/dmzsz/duozhuayu/internal/http/middlewares"
	"github.com/dmzsz/duozhuayu/internal/http/routes"
	"github.com/dmzsz/duozhuayu/internal/utils"
	"github.com/redis/go-redis/v9"

	emailservice "github.com/dmzsz/duozhuayu/internal/services/emailservice/v1"
	"github.com/dmzsz/duozhuayu/pkg/logger"
	"github.com/dmzsz/duozhuayu/pkg/mail"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {
	// setup databases
	conn, err := utils.SetupPostgresConnection()
	if err != nil {
		return nil, err
	}

	// setup router
	router := setupRouter()

	// cache
	var redisCache *caches.RedisCache
	if configs.IsRedis() {
		redisCache = caches.NewRedisCache(&redis.Options{}, time.Duration(0))
	}
	ristrettoCache, err := caches.NewRistrettoCache()
	if err != nil {
		panic(err)
	}

	// emailerService
	emailerService := emailservice.NewEmailService(mail.NewMail())

	// user middleware
	// user with valid basic token can access endpoint
	authMiddleware := middlewares.NewAuthMiddleware(jwtService, []string{constants.UserRoleID})

	// admin middleware
	// only user with valid admin token can access endpoint
	_ = middlewares.NewAuthMiddleware(jwtService, []string{constants.AdminRoleID})

	// API Routes
	api := router.Group("api")
	api.GET("/", routes.RootHandler)
	routes.NewUsersRoute(api, conn, redisCache, ristrettoCache, authMiddleware, emailerService).Routes()

	// we can add web pages if needed
	// web := router.Group("web")
	// ...

	// setup http server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", configs.AppConfig.DatabaseConfig.RDBMS.Env.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &App{
		HttpServer: server,
	}, nil
}

func (a *App) Run() (err error) {
	// Gracefull Shutdown
	go func() {
		logger.InfoF("success to listen and serve on :%d", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, configs.AppConfig.DatabaseConfig.RDBMS.Env.Port)
		if err := a.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit
	logger.Info("shutdown server ...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.HttpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	logger.Info("timeout of 5 seconds.", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	logger.Info("server exiting", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	return
}

func setupRouter() *gin.Engine {
	// set the runtime mode
	var mode = gin.ReleaseMode
	if configs.AppConfig.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	// create a new router instance
	router := gin.New()

	// set up middlewares
	router.Use(middlewares.CORSMiddleware())
	router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	router.Use(gin.Recovery())

	return router
}
