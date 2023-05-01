package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const (
	collectionMstApp = "mst_app"
)

func (r *Repository) GetAllMasterAppFromDB(ctx context.Context) ([]MasterApp, error) {
	var result map[string]interface{}
	err := r.db.NewRef(collectionMstApp).Get(ctx, &result)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	slice := make([]interface{}, 0, len(result))
	for _, app := range result {
		slice = append(slice, app)
	}
	sliceStr, _ := json.Marshal(slice)

	var data []MasterApp
	err = json.Unmarshal([]byte(sliceStr), &data)
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

func (r *Repository) InsertNewMasterAppToDB(ctx context.Context, request NewMasterAppRequest) error {
	appID := uuid.NewString()
	master := MasterApp{
		AppID:          appID,
		AppName:        request.AppName,
		AppPackageName: request.PackageName,
		CreatedAt:      time.Now(),
	}
	err := r.db.NewRef(collectionMstApp).Child(appID).Set(ctx, master)
	return err
}

func (r *Repository) UpdateMasterAppOnDB(ctx context.Context, masterApp MasterApp) error {
	return r.db.NewRef(collectionMstApp).Update(ctx,
		map[string]interface{}{
			masterApp.AppID: masterApp,
		})
}
