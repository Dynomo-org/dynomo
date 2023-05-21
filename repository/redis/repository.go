package redis

import "github.com/redis/go-redis/v9"

type Repository struct {
	redis *redis.Client
}

func New(redis *redis.Client) *Repository {
	return &Repository{redis}
}
