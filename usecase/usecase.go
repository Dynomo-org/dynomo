package usecase

import (
	"context"
	"dynapgen/repository"
)

type repoProvider interface {
	// db provider
	GetAllMasterAppFromDB(ctx context.Context) ([]repository.MasterApp, error)
	GetMasterAppFromDB(ctx context.Context, appID string) (repository.MasterApp, error)
	InsertNewMasterAppToDB(ctx context.Context, name string) (string, error)
	InsertMasterAppToDB(ctx context.Context, master repository.MasterApp) error
	UpdateMasterAppOnDB(ctx context.Context, masterApp repository.MasterApp) error

	// cache provider
	GetMasterAppFromCache(ctx context.Context, appID string) (repository.MasterApp, error)
	InvalidateMasterAppOnCache(ctx context.Context, appID string) error
	StoreMasterAppToCache(ctx context.Context, masterApp repository.MasterApp) error
}

type Usecase struct {
	repo repoProvider
}

func NewUsecase(repo repoProvider) *Usecase {
	return &Usecase{repo: repo}
}
