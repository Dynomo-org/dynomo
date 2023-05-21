package usecase

import (
	"dynapgen/repository/db"
	"dynapgen/repository/github"
	"dynapgen/repository/redis"
)

type Usecase struct {
	db     *db.Repository
	cache  *redis.Repository
	github *github.Repository
}

func NewUsecase(db *db.Repository, cache *redis.Repository, github *github.Repository) *Usecase {
	return &Usecase{
		db:     db,
		cache:  cache,
		github: github,
	}
}
