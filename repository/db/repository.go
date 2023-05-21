package db

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Rebind(query string) string {
	return sqlx.Rebind(sqlx.DOLLAR, query)
}
