package main

import (
	"context"
	"dynapgen/handler"
	"dynapgen/repository"
	"dynapgen/usecase"
	"fmt"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"google.golang.org/api/option"
)

const (
	maxRedisConnectRetryAttempts = 5
)

type Server struct {
	handler *handler.Handler
}

func NewServer(handler *handler.Handler) *Server {
	return &Server{handler: handler}
}

func main() {
	env := os.Getenv("ENV")

	// init redis connection
	redisAddr := "localhost:6379"
	if env == "production" {
		redisAddr = "redis:6379"
		gin.SetMode(gin.ReleaseMode)
	}
	redisOpt := &redis.Options{
		Addr: redisAddr,
	}
	redisConn := connectRedis(redisOpt)
	if redisConn == nil {
		log.Fatal("error connecting to redis")
	}
	fmt.Println("connected to redis")

	// init firebase connection
	ctx := context.Background()
	opt := option.WithCredentialsFile("config/adc.json")
	conf := &firebase.Config{
		DatabaseURL: "https://baim-dynamic-app-default-rtdb.firebaseio.com/",
	}
	firebaseConn, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error connecting to firebase")
	}

	// init db connection
	dbConn, err := firebaseConn.Database(ctx)
	if err != nil {
		log.Fatalf("error connecting to database")
	}
	fmt.Println("connected to database")

	// init app layers
	repository := repository.NewRepository(redisConn, dbConn)
	usecase := usecase.NewUsecase(repository)
	handler := handler.NewHandler(usecase)

	r := gin.Default()
	r.Use(cors.Default())
	handler.RegisterHandler(r)
	r.Run(":5000")
}

// init redis with retry mechanism
func connectRedis(redisOpt *redis.Options) *redis.Client {
	for i := 0; i < maxRedisConnectRetryAttempts; i++ {
		fmt.Printf("Trying to connect to redis (%d/%d)\n", i+1, maxRedisConnectRetryAttempts)
		if client := redis.NewClient(redisOpt); client != nil {
			return client
		}
		time.Sleep(2 * time.Second)
	}

	return nil
}
