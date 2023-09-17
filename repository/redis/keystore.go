package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	fieldKeystore = "keystore"
	keystoreTTL   = 5 * 60 * time.Second

	keyBuildKeystoreStatus = "app_%s_build_keystore"
)

func (r *Repository) GetKeystoreBuildStatus(ctx context.Context, appID string) (BuildStatus, error) {
	result, err := r.redis.HGet(ctx, appID, fieldKeystore).Result()
	if err != nil {
		if err == redis.Nil {
			return BuildStatus{}, nil
		}

		return BuildStatus{}, err
	}

	var buildStatus BuildStatus
	err = json.Unmarshal([]byte(result), &buildStatus)
	if err != nil {
		return BuildStatus{}, nil
	}

	return buildStatus, nil
}

func (r *Repository) InvalidateKeystoreBuildStatus(ctx context.Context, appID string) error {
	result := r.redis.HDel(ctx, appID, fieldKeystore)
	return result.Err()
}

func (r *Repository) InsertKeystoreBuildStatus(ctx context.Context, appID string, buildStatus BuildStatus) error {
	marshalled, err := json.Marshal(buildStatus)
	if err != nil {
		return err
	}

	return r.HSetEX(ctx, appID, fieldKeystore, string(marshalled), keystoreTTL)
}

func (r *Repository) SetBuildKeystoreStatus(ctx context.Context, param UpdateBuildStatusParam) error {
	key := fmt.Sprintf(keyBuildKeystoreStatus, param.AppID)
	marshalled, err := json.Marshal(param.BuildStatus)
	if err != nil {
		return err
	}

	return r.redis.SetEx(ctx, key, string(marshalled), TTL15Minutes).Err()
}
