package repository

import (
	"firebase.google.com/go/v4/db"
	"github.com/google/go-github/v52/github"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	redis  *redis.Client
	db     *db.Client
	github *github.Client
}

func NewRepository(redis *redis.Client, db *db.Client, github *github.Client) *Repository {
	return &Repository{
		redis:  redis,
		db:     db,
		github: github,
	}
}
