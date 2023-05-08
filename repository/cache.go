package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	redisTTL = 24 * time.Hour
)

func (r *Repository) GetAppFromCache(ctx context.Context, appID string) (App, error) {
	result, err := r.redis.Get(ctx, appID).Result()
	if err != nil {
		if err == redis.Nil {
			return App{}, nil
		}

		return App{}, err
	}

	var app App
	err = json.Unmarshal([]byte(result), &app)
	if err != nil {
		return App{}, nil
	}

	return app, nil
}

func (r *Repository) InvalidateAppOnCache(ctx context.Context, appID string) error {
	result := r.redis.Del(ctx, appID)
	return result.Err()
}

func (r *Repository) StoreAppToCache(ctx context.Context, App App) error {
	marshalled, err := json.Marshal(App)
	if err != nil {
		return err
	}
	result := r.redis.SetEx(ctx, App.AppID, string(marshalled), redisTTL)
	return result.Err()
}
