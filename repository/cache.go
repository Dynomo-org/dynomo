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

func (r *Repository) GetMasterAppFromCache(ctx context.Context, appID string) (MasterApp, error) {
	result, err := r.redis.Get(ctx, appID).Result()
	if err != nil {
		if err == redis.Nil {
			return MasterApp{}, nil
		}

		return MasterApp{}, err
	}

	var masterApp MasterApp
	err = json.Unmarshal([]byte(result), &masterApp)
	if err != nil {
		return MasterApp{}, nil
	}

	return masterApp, nil
}

func (r *Repository) InvalidateMasterAppOnCache(ctx context.Context, appID string) error {
	result := r.redis.Del(ctx, appID)
	return result.Err()
}

func (r *Repository) StoreMasterAppToCache(ctx context.Context, masterApp MasterApp) error {
	marshalled, err := json.Marshal(masterApp)
	if err != nil {
		return err
	}
	result := r.redis.SetEx(ctx, masterApp.AppID, string(marshalled), redisTTL)
	return result.Err()
}
