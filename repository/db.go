package repository

import (
	"context"

	"github.com/google/uuid"
)

const (
	collectionMstApp = "mst_app"
)

func (r *Repository) GetAllMasterAppFromDB(ctx context.Context) ([]MasterApp, error) {
	var data []MasterApp
	err := r.db.NewRef(collectionMstApp).Get(ctx, data)
	return data, err
}

func (r *Repository) GetMasterAppFromDB(ctx context.Context, appID string) (MasterApp, error) {
	var result MasterApp
	err := r.db.NewRef(collectionMstApp).Child(appID).Get(ctx, &result)
	return result, err
}

func (r *Repository) InsertMasterAppToDB(ctx context.Context, master MasterApp) error {
	_, err := r.db.NewRef(collectionMstApp).Push(ctx, master)
	return err
}

func (r *Repository) InsertNewMasterAppToDB(ctx context.Context, name string) (string, error) {
	appID := uuid.NewString()
	master := MasterApp{
		AppID: appID,
		Name:  name,
	}
	_, err := r.db.NewRef(collectionMstApp).Push(ctx, master)
	return appID, err
}

func (r *Repository) UpdateMasterAppOnDB(ctx context.Context, masterApp MasterApp) error {
	return r.db.NewRef(collectionMstApp).Update(ctx,
		map[string]interface{}{
			masterApp.AppID: masterApp,
		})
}
