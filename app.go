package main

import (
	"context"
	"dynapgen/handler"
	"dynapgen/repository"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"google.golang.org/api/option"
)

type Server struct {
	handler *handler.Handler
}

func NewServer(handler *handler.Handler) *Server {
	return &Server{handler: handler}
}

func (s *Server) RegisterHandler(r *gin.Engine) {
	r.GET("/ping", s.handler.Ping)

	// Admin router
	r.GET("/_admin/collections", s.handler.GetAllCollection)

	// Public Router
	r.POST("/app/createapp", s.handler.CreateNewMasterApp)
}

func main() {
	env := os.Getenv("ENV")

	// init redis connection
	redisAddr := "localhost:6379"
	if env == "production" {
		redisAddr = "redis:6379"
		gin.SetMode(gin.ReleaseMode)
	}
	redisConn := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
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
	handler := handler.NewHandler(repository)
	server := NewServer(handler)

	r := gin.Default()
	server.RegisterHandler(r)
	r.Run(":5000")
}
