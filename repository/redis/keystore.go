package redis

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

const (
	fieldKeystore = "keystore"
	keystoreTTL   = 5 * 60
)

func (r *Repository) GetKeystore(ctx context.Context, appID string) (Keystore, error) {
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

func (r *Repository) InvalidateKeystore(ctx context.Context, appID string) error {
	result := r.redis.HDel(ctx, appID, fieldKeystore)
	return result.Err()
}
func (r *Repository) InsertKeystore(ctx context.Context, appID string, keystore Keystore) error {
	marshalled, err := json.Marshal(keystore)
	if err != nil {
		return err
	}

	return r.HSetEX(ctx, appID, fieldKeystore, string(marshalled), keystoreTTL)
}
