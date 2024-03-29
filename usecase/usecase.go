package usecase

import (
	"dynapgen/repository/db"
	"dynapgen/repository/github"
	"dynapgen/repository/nsq"
	"dynapgen/repository/redis"
	"strings"
)

type Usecase struct {
	db     *db.Repository
	cache  *redis.Repository
	github *github.Repository
	mq     *nsq.Repository
}

func NewUsecase(db *db.Repository, cache *redis.Repository, github *github.Repository, mq *nsq.Repository) *Usecase {
	return &Usecase{
		db:     db,
		cache:  cache,
		github: github,
		mq:     mq,
	}
}

func getRepositoryName(url string) string {
	segments := strings.Split(url, "/")
	return segments[len(segments)-1]
}
