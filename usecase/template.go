package usecase

import (
	"context"
	"dynapgen/repository/redis"
)

func (uc *Usecase) GetTemplateByID(ctx context.Context, id string) (Template, error) {
	cached, err := uc.cache.GetTemplate(ctx, id)
	if err != nil {
		return Template{}, err
	}
	if cached.ID != "" {
		return Template(cached), nil
	}

	template, err := uc.db.GetTemplateByID(ctx, id)
	if err != nil {
		return Template{}, err
	}

	err = uc.cache.InsertTemplate(ctx, redis.Template(template))
	if err != nil {
		return Template{}, err
	}

	return Template(template), nil
}
