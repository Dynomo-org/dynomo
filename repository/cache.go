package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	defaultTTL  = 24 * 60 * 60
	keystoreTTL = 5 * 60

	fieldMasterApp = "master_app"
	fieldKeystore  = "keystore"
)

func (r *Repository) GetAppFromCache(ctx context.Context, appID string) (App, error) {
	result, err := r.redis.HGet(ctx, appID, fieldMasterApp).Result()
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

func (r *Repository) GetKeystoreFromCache(ctx context.Context, appID string) (Keystore, error) {
	result, err := r.redis.HGet(ctx, appID, fieldKeystore).Result()
	if err != nil {
		if err == redis.Nil {
			return Keystore{}, nil
		}

		return Keystore{}, err
	}

	var keystore Keystore
	err = json.Unmarshal([]byte(result), &keystore)
	if err != nil {
		return Keystore{}, nil
	}

	return keystore, nil
}

func (r *Repository) InvalidateAppOnCache(ctx context.Context, appID string) error {
	result := r.redis.HDel(ctx, appID, fieldMasterApp)
	return result.Err()
}

func (r *Repository) InvalidateKeystoreOnCache(ctx context.Context, appID string) error {
	result := r.redis.HDel(ctx, appID, fieldKeystore)
	return result.Err()
}

func (r *Repository) StoreAppToCache(ctx context.Context, App App) error {
	marshalled, err := json.Marshal(App)
	if err != nil {
		return err
	}

	return r.HSetEX(ctx, App.AppID, fieldMasterApp, string(marshalled), defaultTTL)
}

func (r *Repository) StoreKeystoreToCache(ctx context.Context, appID string, keystore Keystore) error {
	marshalled, err := json.Marshal(keystore)
	if err != nil {
		return err
	}

	return r.HSetEX(ctx, appID, fieldKeystore, string(marshalled), keystoreTTL)
}

func (r *Repository) HSetEX(ctx context.Context, key, field string, value interface{}, expireSecond int64) error {
	_, err := r.redis.HSet(ctx, key, field, value).Result()
	if err != nil {
		return err
	}
	return r.redis.Expire(ctx, key, time.Duration(expireSecond)*time.Second).Err()
}
