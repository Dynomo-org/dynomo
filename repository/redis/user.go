package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"dynapgen/constants"

	"github.com/redis/go-redis/v9"
)

const (
	keyUserRole = "user_role_%s"
	keyUserInfo = "user_info_%s"
)

func (r *Repository) SetUserRoleIDMapForUserID(ctx context.Context, userID string, roles map[constants.UserRole]struct{}) error {
	key := fmt.Sprintf(keyUserRole, userID)

	marshaledRoles, err := json.Marshal(roles)
	if err != nil {
		return err
	}

	return r.redis.SetEx(ctx, key, marshaledRoles, TTL1Day).Err()
}

func (r *Repository) GetUserRoleIDMapByUserID(ctx context.Context, userID string) (map[constants.UserRole]struct{}, error) {
	key := fmt.Sprintf(keyUserRole, userID)

	result, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var userRoles map[constants.UserRole]struct{}
	if err := json.Unmarshal([]byte(result), &userRoles); err != nil {
		return nil, err
	}

	return userRoles, nil
}
