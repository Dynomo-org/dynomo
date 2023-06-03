package usecase

import (
	"context"
	"dynapgen/repository/db"
	"dynapgen/repository/github"
	"dynapgen/utils/log"
	"errors"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	errorFailedToGetAppFull = errors.New("error getting app full")
)

func (uc *Usecase) GetAllApps(ctx context.Context, param GetAppListParam) (GetAppListResponse, error) {
	apps, err := uc.db.GetAppsByUserID(ctx, param.OwnerID)
	if err != nil {
		log.Error(param, err, "uc.db.GetApps() got error - GetAllApp")
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
		log.Error(map[string]interface{}{"app_id": appID}, err, "uc.db.GetApp() got error - GetApp")
		return App{}, err
	}
	if app.ID == "" {
		return App{}, errorAppNotFound
	}

	return App(app), nil
}

func (uc *Usecase) GetAppAds(ctx context.Context, appID string) ([]AppAds, error) {
	ads, err := uc.db.GetAppAdsByAppID(ctx, appID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": appID}, err, "uc.db.GetAppAdsByAppID() got error - GetAppAds")
		return nil, err
	}

	result := make([]AppAds, 0, len(ads))
	for _, ad := range ads {
		result = append(result, AppAds(ad))
	}

	return result, nil
}

func (uc *Usecase) GetAppFull(ctx context.Context, appID string) (AppFull, error) {
	cachedApp, err := uc.cache.GetAppFullByID(ctx, appID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": appID}, err, "uc.cache.GetUserRoleIDMapByUserID() got error - GetAppFull")
	}
	if cachedApp.ID != "" {
		return convertAppFullFromCache(cachedApp), nil
	}

	var (
		wg sync.WaitGroup

		errMsgs     []string
		app         db.App
		appAds      []db.AppAds
		appContents []db.AppContent
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		app, err = uc.db.GetApp(ctx, appID)
		if err != nil {
			log.Error(map[string]interface{}{"app_id": appID}, err, "uc.db.GetApp() got error - GetAppFull")
			errMsgs = append(errMsgs, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		appAds, err = uc.db.GetAppAdsByAppID(ctx, appID)
		if err != nil {
			log.Error(map[string]interface{}{"app_id": appID}, err, "uc.db.GetAppAdsByAppID() got error - GetAppFull")
			errMsgs = append(errMsgs, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		appContents, err = uc.db.GetAppContentsByAppID(ctx, appID)
		if err != nil {
			log.Error(map[string]interface{}{"app_id": appID}, err, "uc.db.GetAppContentsByAppID() got error - GetAppFull")
			errMsgs = append(errMsgs, err.Error())
		}
	}()

	wg.Wait()
	if len(errMsgs) > 0 {
		log.Error(map[string]interface{}{"app_id": appID}, errors.New(strings.Join(errMsgs, ", ")), "error getting app full - GetAppFull")
		return AppFull{}, errorFailedToGetAppFull
	}

	return buildAppFull(app, appAds, appContents), nil
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
		CreatedAt:                  time.Now(),
	}

	err := uc.db.InsertApp(ctx, app)
	if err != nil {
		log.Error(request, err, "uc.repo.InsertApp() got error - NewApp")
		return err
	}

	return nil
}

func (uc *Usecase) NewAppAds(ctx context.Context, request NewAppAdsRequest) error {
	ads := db.AppAds{
		ID:               uuid.NewString(),
		AppID:            request.AppID,
		Type:             request.Type,
		OpenAdID:         request.OpenAdID,
		BannerAdID:       request.BannerAdID,
		InterstitialAdID: request.InterstitialAdID,
		RewardAdID:       request.RewardAdID,
		NativeAdID:       request.NativeAdID,
		CreatedAt:        time.Now(),
	}

	err := uc.db.InsertAppAds(ctx, ads)
	if err != nil {
		log.Error(request, err, "uc.db.InsertAppAds() got error - NewAppAds")
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
