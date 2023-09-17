package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	redis *redis.Client
}

func New(redis *redis.Client) *Repository {
	return &Repository{redis}
}

func (r *Repository) HSetEX(ctx context.Context, key, field string, value interface{}, expire time.Duration) error {
	_, err := r.redis.HSet(ctx, key, field, value).Result()
	if err != nil {
		return err
	}
	return r.redis.Expire(ctx, key, expire).Err()
}
