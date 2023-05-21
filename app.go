package main

import (
	"context"
	"dynapgen/handler"
	repoDB "dynapgen/repository/db"
	repoGithub "dynapgen/repository/github"
	repoRedis "dynapgen/repository/redis"
	"dynapgen/usecase"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v52/github"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read config, err:", err.Error())
	}

	// init redis connection
	redisAddr := viper.GetString("REDIS_ADDR")
	if viper.GetString("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	redisOpt := &redis.Options{
		Addr: redisAddr,
	}
	redisConn := initRedis(redisOpt)
	if redisConn == nil {
		log.Fatal("error connecting to redis - ", err)
	}
	fmt.Println("connected to redis")

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
		log.Fatal("error connecting to DB - ", err)
	}
	fmt.Println("connected to DB")

	// init github connection
	githubClientKey := viper.GetString("GITHUB_CLIENT_KEY")
	githubConn := github.NewTokenClient(ctx, githubClientKey)
	fmt.Println("connected to github")

	// init app layers
	repoDB := repoDB.New(dbClient)
	repoRedis := repoRedis.New(redisConn)
	repoGithub := repoGithub.New(githubConn)
	usecase := usecase.NewUsecase(repoDB, repoRedis, repoGithub)
	handler := handler.NewHandler(usecase)

	r := gin.Default()
	r.Use(cors.Default())
	handler.RegisterHandler(r)
	r.Run(":5000")
}

// init redis with retry mechanism
func initRedis(redisOpt *redis.Options) *redis.Client {
	for i := 0; i < maxConnectionRetryAttempts; i++ {
		fmt.Printf("Connecting to redis (%d/%d)\n", i+1, maxConnectionRetryAttempts)
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
		fmt.Printf("Connecting to DB (%d/%d)\n", i+1, maxConnectionRetryAttempts)
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
