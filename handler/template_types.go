package handler

import "time"

type Template struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	RepositoryURL string            `json:"repository_url"`
	Styles        map[string]string `json:"styles"`
	Strings       map[string]string `json:"strings"`
	Type          int               `json:"type"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at,omitempty"`
}
