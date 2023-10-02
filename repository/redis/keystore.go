package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	fieldKeystore = "keystore"

	keyBuildKeystoreStatus = "build_keystore_%s"
)

var (
	errNoBuildKeystoreStatus = errors.New("no build keystore job running")
)

func (r *Repository) GetBuildKeystoreStatus(ctx context.Context, appID string) (BuildStatus, error) {
	key := fmt.Sprintf(keyBuildKeystoreStatus, appID)
	result, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return BuildStatus{
				Status:       BuildStatusEnumFailed,
				ErrorMessage: errNoBuildKeystoreStatus.Error(),
			}, nil
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

	return r.HSetEX(ctx, appID, fieldKeystore, string(marshalled), TTL1Day)
}

func (r *Repository) SetBuildKeystoreStatus(ctx context.Context, param UpdateBuildStatusParam) error {
	key := fmt.Sprintf(keyBuildKeystoreStatus, param.BuildID)
	marshalled, err := json.Marshal(param.BuildStatus)
	if err != nil {
		return err
	}

	return r.redis.SetEx(ctx, key, string(marshalled), TTL1Day).Err()
}
