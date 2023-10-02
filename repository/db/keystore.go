package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

func (r *Repository) GetKeystoreByID(ctx context.Context, id string) (Keystore, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result Keystore
	if err := r.db.GetContext(ctx, &result, queryGetKeystoreByID, id); err != nil && err != sql.ErrNoRows {
		return Keystore{}, err
	}

	return result, nil
}

func (r *Repository) GetKeystoresByOwnerID(ctx context.Context, ownerID string) ([]Keystore, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	var result []Keystore
	if err := r.db.SelectContext(ctx, &result, queryGetKeystoreByOwnerID, ownerID); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) UpsertKeystore(ctx context.Context, keystore Keystore) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	keystore.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}
	query, args, err := sqlx.Named(queryUpsertKeystore, keystore)
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, r.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}
