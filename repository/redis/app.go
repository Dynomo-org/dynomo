package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	defaultTTL = 24 * 60 * 60 * time.Second

	keyApp = "app_%s"
)

func (r *Repository) GetAppFullByID(ctx context.Context, appID string) (AppFull, error) {
	key := fmt.Sprintf(keyApp, appID)
	result, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return AppFull{}, nil
		}

		return AppFull{}, err
	}

	var app AppFull
	err = json.Unmarshal([]byte(result), &app)
	if err != nil {
		return AppFull{}, nil
	}

	return app, nil
}

func (r *Repository) InvalidateAppFull(ctx context.Context, appID string) error {
	key := fmt.Sprintf(keyApp, appID)
	result := r.redis.Del(ctx, key)
	return result.Err()
}

func (r *Repository) InsertAppFull(ctx context.Context, app AppFull) error {
	key := fmt.Sprintf(keyApp, app.ID)
	marshalled, err := json.Marshal(app)
	if err != nil {
		return err
	}

	return r.redis.SetEx(ctx, key, string(marshalled), defaultTTL).Err()
}
