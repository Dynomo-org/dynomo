package db

import (
	"database/sql"
	"time"
)

type Template struct {
	ID            string       `db:"id"`
	Name          string       `db:"name"`
	RepositoryURL string       `db:"repository_url"`
	Styles        string       `db:"styles"`  // jsonb type
	Strings       string       `db:"strings"` // jsonb type
	Type          int          `db:"type"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
}
