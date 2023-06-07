package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func (r *Repository) GetAppAdsByID(ctx context.Context, id string) (AppAds, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result AppAds
	if err := r.db.GetContext(ctx, &result, queryGetAppAdsByID, id); err != nil && err != sql.ErrNoRows {
		return AppAds{}, err
	}

	return result, nil
}

func (r *Repository) GetAppAdsByAppID(ctx context.Context, appID string) ([]AppAds, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result []AppAds
	if err := r.db.SelectContext(ctx, &result, queryGetAppAdsByAppID, appID); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) InsertAppAds(ctx context.Context, appAds AppAds) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	query, args, err := sqlx.Named(queryInsertAppAds, appAds)
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, r.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateAds(ctx context.Context, ads AppAds) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	query, args, err := sqlx.Named(queryUpdateAppAds, ads)
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, r.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}
