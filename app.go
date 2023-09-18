package main

import (
	"context"
	"dynapgen/handler"
	"fmt"
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
	r.Run(":5000")
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
