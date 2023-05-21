package usecase

import (
	"context"
	"dynapgen/repository/db"
	"dynapgen/repository/github"
	"dynapgen/utils/log"
	"os"
	"time"

	"github.com/google/uuid"
)

func (uc *Usecase) GetAllApps(ctx context.Context, param GetAppListParam) (GetAppListResponse, error) {
	apps, err := uc.db.GetAllApps(ctx)
	if err != nil {
		log.Error(nil, err, "uc.db.GetApps() got error - GetAllApp")
		return GetAppListResponse{}, err
	}

	result := make([]App, 0, len(apps))
	for _, app := range apps {
		result = append(result, App{
			ID:          app.ID,
			Name:        app.Name,
			IconURL:     app.IconURL,
			PackageName: app.PackageName,
			CreatedAt:   app.CreatedAt,
			UpdatedAt:   app.UpdatedAt,
		})
	}

	return buildAppListResponse(result, param), nil
}

func (uc *Usecase) GetApp(ctx context.Context, appID string) (App, error) {
	app, err := uc.db.GetApp(ctx, appID)
	if err != nil {
		return App{}, err
	}
	if app.ID == "" {
		return App{}, errorAppNotFound
	}

	return App(app), nil
}

func (uc *Usecase) NewApp(ctx context.Context, request NewAppRequest) error {
	app := db.App{
		ID:                         uuid.NewString(),
		OwnerID:                    request.OwnerID,
		Name:                       request.AppName,
		PackageName:                request.PackageName,
		Type:                       AppTypeUnset,
		Version:                    1,
		VersionCode:                "1.0.0",
		IconURL:                    "https://raw.githubusercontent.com/Dynapgen/master-storage-1/main/assets/default-icon.png",
		ColorPrimary:               "#FFBB86FC",
		ColorPrimaryVariant:        "#FF3700B3",
		ColorOnPrimary:             "#FFFFFFFF",
		InterstitialIntervalSecond: 30,
	}

	err := uc.db.InsertApp(ctx, app)
	if err != nil {
		log.Error(request, err, "uc.repo.InsertApp() got error - NewApp")
		return err
	}

	return nil
}

func (uc *Usecase) UpdateApp(ctx context.Context, request App) error {
	app, err := uc.GetApp(ctx, request.ID)
	if err != nil {
		log.Error(app, err, "uc.Get() got error - UpdateApp")
	}

	app.updateWith(request)
	timeNow := time.Now()
	param := db.App(app)
	param.UpdatedAt = &timeNow
	err = uc.db.UpdateApp(ctx, param)
	if err != nil {
		log.Error(param, err, "uc.db.UpdateApp() got error - UpdateApp")
		return err
	}

	return nil
}

func (uc *Usecase) UpdateAppIcon(ctx context.Context, appID string, iconName, localPath string) error {
	meta := map[string]interface{}{
		"app_id":     appID,
		"local_path": localPath,
		"icon_name":  iconName,
	}

	app, err := uc.GetApp(ctx, appID)
	if err != nil {
		log.Error(meta, err, "uc.GetApp() got error - UpdateAppIcon")
		return err
	}

	iconURL, err := uc.github.Upload(ctx, github.UploadFileParam{
		FilePathLocal:         localPath,
		FileName:              iconName,
		DestinationFolderPath: appID + "/",
		ReplaceIfNameExists:   true,
	})
	if err != nil {
		log.Error(meta, err, "uc.repo.UploadToGithub() got error - UpdateAppIcon")
		return err
	}
	os.Remove(localPath)

	timeNow := time.Now()
	app.IconURL = iconURL
	app.UpdatedAt = &timeNow
	param := db.App(app)
	err = uc.db.UpdateApp(ctx, param)
	if err != nil {
		log.Error(param, err, "uc.repo.UpdateAppOnDB() got error - UpdateAppIcon")
		return err
	}

	return nil
}

func (uc *Usecase) DeleteApp(ctx context.Context, appID string) error {
	err := uc.db.DeleteApp(ctx, appID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": appID}, err, "uc.repo.DeleteApp() - DeleteApp")
		return err
	}

	return nil
}

func buildAppListResponse(apps []App, param GetAppListParam) GetAppListResponse {
	if len(apps) == 0 {
		return GetAppListResponse{
			Apps: []App{},
		}
	}

	return GetAppListResponse{
		Apps:      apps,
		TotalPage: apps[0].Total / param.PerPage,
		Page:      param.Page,
	}
}
