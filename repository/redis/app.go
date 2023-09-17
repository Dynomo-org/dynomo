package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	TTL1Day      = 24 * 60 * 60 * time.Second
	TTL15Minutes = 15 * 60 * time.Second

	keyApp            = "app_%s"
	keyBuildAppStatus = "app_%s_build_app"
)

var (
	errNoBuildAppStatus = errors.New("no build app job running")
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

func (r *Repository) GetBuildAppStatus(ctx context.Context, appID string) (BuildStatus, error) {
	key := fmt.Sprintf(keyBuildAppStatus, appID)
	result, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return BuildStatus{
				Status:       BuildStatusEnumFailed,
				ErrorMessage: errNoBuildAppStatus.Error(),
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

	return r.redis.SetEx(ctx, key, string(marshalled), TTL1Day).Err()
}

func (r *Repository) SetBuildAppStatus(ctx context.Context, param UpdateBuildStatusParam) error {
	key := fmt.Sprintf(keyBuildAppStatus, param.AppID)
	marshalled, err := json.Marshal(param.BuildStatus)
	if err != nil {
		return err
	}

	return r.redis.SetEx(ctx, key, string(marshalled), TTL15Minutes).Err()
}
