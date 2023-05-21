package github

import "github.com/google/go-github/v52/github"

type Repository struct {
	github *github.Client
}

func New(github *github.Client) *Repository {
	return &Repository{github}
}
