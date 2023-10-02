package db

import (
	"database/sql"
	"time"
)

type AppArtifact struct {
	ID          string       `db:"id"`
	AppID       string       `db:"app_id"`
	OwnerID     string       `db:"owner_id"`
	Name        string       `db:"name"`
	Metadata    string       `db:"metadata_encrypted"`
	DownloadURL string       `db:"download_url"`
	BuildStatus int          `db:"build_status"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}
