package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	keyTemplate = "template_%s"
)

func (r *Repository) GetTemplate(ctx context.Context, id string) (Template, error) {
	key := fmt.Sprintf(keyTemplate, id)

	result, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return Template{}, nil
		}

		return Template{}, err
	}

	var template Template
	if err = json.Unmarshal([]byte(result), &template); err != nil {
		return Template{}, err
	}

	return template, nil
}

func (r *Repository) InsertTemplate(ctx context.Context, template Template) error {
	key := fmt.Sprintf(keyTemplate, template.ID)
	marshalled, err := json.Marshal(template)
	if err != nil {
		return err
	}

	return r.redis.SetEx(ctx, key, string(marshalled), TTL1Day).Err()
}
