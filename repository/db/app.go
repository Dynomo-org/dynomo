package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	queryTimeLimit = 5 * time.Second
)

func (r *Repository) DeleteApp(ctx context.Context, appID string) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	if _, err := r.db.QueryContext(ctx, queryDeleteApp, appID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAppsByUserID(ctx context.Context, userID string) ([]App, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result []App
	if err := r.db.SelectContext(ctx, &result, queryGetAppsByUserID, userID); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) GetApp(ctx context.Context, appID string) (App, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result App
	if err := r.db.GetContext(ctx, &result, queryGetApp, appID); err != nil && err != sql.ErrNoRows {
		return App{}, err
	}

	return result, nil
}

func (r *Repository) InsertApp(ctx context.Context, app App) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	query, args, err := sqlx.Named(queryInsertApp, app)
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, r.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateApp(ctx context.Context, app App) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	query, args, err := sqlx.Named(queryUpdateApp, app)
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, r.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}
