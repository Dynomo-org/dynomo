package repository

import (
	"firebase.google.com/go/v4/db"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	redis *redis.Client
	db    *db.Client
}

func NewRepository(redis *redis.Client, db *db.Client) *Repository {
	return &Repository{
		redis: redis,
		db:    db,
	}
}
