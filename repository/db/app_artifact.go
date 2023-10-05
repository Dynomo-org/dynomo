package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func (r *Repository) GetAppArtifactsByAppID(ctx context.Context, appID string) ([]AppArtifactInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	query := fmt.Sprintf(queryGetAppArtifacts, "WHERE app.id = $1")
	var result []AppArtifactInfo
	if err := r.db.SelectContext(ctx, &result, query, appID); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) GetAppArtifactsByOwnerID(ctx context.Context, ownerID string) ([]AppArtifactInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	query := fmt.Sprintf(queryGetAppArtifacts, "WHERE artifact.owner_id = $1")
	var result []AppArtifactInfo
	if err := r.db.SelectContext(ctx, &result, query, ownerID); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) UpsertAppArtifact(ctx context.Context, artifact AppArtifact) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeLimit)
	defer cancel()

	artifact.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}
	query, args, err := sqlx.Named(queryUpsertAppArtifact, artifact)
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, r.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}
