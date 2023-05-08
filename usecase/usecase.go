package usecase

import (
	"context"
	"dynapgen/repository"
)

type repoProvider interface {
	// db provider
	DeleteAppOnDB(ctx context.Context, appID string) error
	GetAllAppFromDB(ctx context.Context) ([]repository.App, error)
	GetAppFromDB(ctx context.Context, appID string) (repository.App, error)
	InsertNewAppToDB(ctx context.Context, request repository.NewAppRequest) error
	InsertAppToDB(ctx context.Context, master repository.App) error
	UpdateAppOnDB(ctx context.Context, App repository.App) error

	// cache provider
	GetAppFromCache(ctx context.Context, appID string) (repository.App, error)
	InvalidateAppOnCache(ctx context.Context, appID string) error
	StoreAppToCache(ctx context.Context, App repository.App) error

	// github provider
	UploadToGithub(ctx context.Context, param repository.UploadFileParam) (string, error)
}

type Usecase struct {
	repo repoProvider
}

func NewUsecase(repo repoProvider) *Usecase {
	return &Usecase{repo: repo}
}
