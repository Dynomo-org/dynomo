package main

import (
	"context"
	"dynapgen/handler"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	repoDB "dynapgen/repository/db"
	repoGithub "dynapgen/repository/github"
	repoNSQ "dynapgen/repository/nsq"
	repoRedis "dynapgen/repository/redis"
	"dynapgen/usecase"
	"dynapgen/util/env"
	"dynapgen/util/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v52/github"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	nsq "github.com/nsqio/go-nsq"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

const (
	maxConnectionRetryAttempts = 5
)

type Server struct {
	handler *handler.Handler
}

func NewServer(handler *handler.Handler) *Server {
	return &Server{handler: handler}
}

func main() {
	ctx := context.Background()

	// init config and env
	err := env.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	// init redis connection
	redisAddr := viper.GetString("REDIS_ADDR")
	if !env.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}
	redisOpt := &redis.Options{
		Addr: redisAddr,
	}
	redisConn := initRedis(redisOpt)
	if redisConn == nil {
		log.Fatal(err)
	}
	log.Info("connected to redis")

	// init db
	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
	)
	dbClient, err := initDB("postgres", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("connected to DB")

	// init github connection
	githubClientKey := viper.GetString("GITHUB_CLIENT_KEY")
	githubConn := github.NewTokenClient(ctx, githubClientKey)
	log.Info("connected to github")

	// init NSQ connection
	nsqConfig := nsq.NewConfig()
	nsqConn, err := nsq.NewProducer(viper.GetString("DNM_NSQ_ADDRESS"), nsqConfig)
	if err != nil {
		log.Fatal(err)
	}
	nsqConn.SetLogger(nil, nsq.LogLevelError)

	// init app layers
	repoDB := repoDB.New(dbClient)
	repoRedis := repoRedis.New(redisConn)
	repoGithub := repoGithub.New(githubConn)
	repoNsq := repoNSQ.New(nsqConn)

	usecase := usecase.NewUsecase(repoDB, repoRedis, repoGithub, repoNsq)
	handler := handler.NewHandler(usecase)

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(config))
	handler.RegisterHandler(r)

	server := &http.Server{
		Addr:    ":5000",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully stop the server and its dependencies
	if err := server.Shutdown(ctx); err != nil {
		log.Error(err, "Server Shutdown Error")
	}
	nsqConn.Stop()

	select {
	case <-ctx.Done():
		log.Info("timeout of 5 seconds.")
	default:
		log.Info("Server exiting")
	}
}

// init redis with retry mechanism
func initRedis(redisOpt *redis.Options) *redis.Client {
	for i := 0; i < maxConnectionRetryAttempts; i++ {
		log.Info(fmt.Sprintf("Connecting to redis (%d/%d)\n", i+1, maxConnectionRetryAttempts))
		if client := redis.NewClient(redisOpt); client != nil {
			return client
		}
		time.Sleep(2 * time.Second)
	}

	return nil
}

func initDB(driver, connectionString string) (*sqlx.DB, error) {
	var connectingError error
	for i := 0; i < maxConnectionRetryAttempts; i++ {
		log.Info(fmt.Sprintf("Connecting to DB (%d/%d)\n", i+1, maxConnectionRetryAttempts))
		db, err := sqlx.Connect("postgres", connectionString)
		if err != nil {
			connectingError = err
			time.Sleep(2 * time.Second)
			continue
		}
		return db, nil
	}

	return nil, connectingError
}
