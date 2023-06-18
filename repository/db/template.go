package db

import (
	"context"
)

func (repo *Repository) GetTemplateByID(ctx context.Context, id string) (Template, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result Template
	if err := repo.db.GetContext(ctx, &result, queryGetTemplateByID, id); err != nil {
		return Template{}, err
	}

	return result, nil
}
