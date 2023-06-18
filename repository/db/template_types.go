package db

type Template struct {
	ID            string `db:"id"`
	Name          string `db:"name"`
	RepositoryURL string `db:"repository_url"`
	Type          int    `db:"type"`
}
