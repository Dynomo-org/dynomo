package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func (r *Repository) GetAppContentsByAppID(ctx context.Context, appID string) ([]AppContent, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result []AppContent
	if err := r.db.SelectContext(ctx, &result, queryGetAppContentsByAppID, appID); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) GetAppContentByID(ctx context.Context, id string) (AppContent, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result AppContent
	if err := r.db.GetContext(ctx, &result, queryGetAppContentByID, id); err != nil {
		return AppContent{}, err
	}

	return result, nil
}

func (r *Repository) InsertAppContent(ctx context.Context, content AppContent) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	query, args, err := sqlx.Named(queryInsertAppContent, content)
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, r.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateAppContent(ctx context.Context, content AppContent) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	query, args, err := sqlx.Named(queryUpdateAppContent, content)
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, r.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}
