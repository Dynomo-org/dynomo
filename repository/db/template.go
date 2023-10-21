package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func (r *Repository) GetAllTemplates(ctx context.Context) ([]Template, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result []Template
	if err := r.db.SelectContext(ctx, &result, queryGetAllTemplates); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) GetTemplateByID(ctx context.Context, id string) (Template, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result Template
	if err := r.db.GetContext(ctx, &result, queryGetTemplateByID, id); err != nil {
		return Template{}, err
	}

	return result, nil
}

func (r *Repository) InsertTemplate(ctx context.Context, template Template) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	query, args, err := sqlx.Named(queryInsertTemplate, template)
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, r.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateTemplate(ctx context.Context, template Template) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	query, args, err := sqlx.Named(queryUpdateTemplate, template)
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, r.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}
