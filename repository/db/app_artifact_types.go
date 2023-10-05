package db

import (
	"database/sql"
	"time"
)

type AppArtifactInfo struct {
	Total       int       `db:"total"`
	ID          string    `db:"id"`
	Name        string    `db:"artifact_name"`
	AppName     string    `db:"app_name"`
	AppID       string    `db:"app_id"`
	IconURL     string    `db:"icon_url"`
	DownloadURL string    `db:"download_url"`
	BuildStatus int       `db:"build_status"`
	CreatedAt   time.Time `db:"created_at"`
}

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
