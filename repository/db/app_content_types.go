package db

import (
	"database/sql"
	"time"
)

type AppContent struct {
	ID           string       `db:"id"`
	AppID        string       `db:"app_id"`
	Title        string       `db:"title"`
	CategoryID   string       `db:"category_id"`
	Description  string       `db:"description"`
	Content      string       `db:"content"`
	ThumbnailURL string       `db:"thumbnail_url"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
}
