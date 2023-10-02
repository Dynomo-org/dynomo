package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

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
