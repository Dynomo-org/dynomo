package db

import (
	"database/sql"
	"time"
)

type GetKeystoreParam struct {
	OwnerID     string
	BuildStatus int
}

type Keystore struct {
	Total       int          `db:"total"`
	ID          string       `db:"id"`
	OwnerID     string       `db:"owner_id"`
	Name        string       `db:"name"`
	Alias       string       `db:"alias"`
	Metadata    string       `db:"metadata_encrypted"`
	DownloadURL string       `db:"download_url"`
	BuildStatus int          `db:"build_status"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}
