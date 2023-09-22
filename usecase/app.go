package usecase

import (
	"context"
	"dynapgen/repository/db"
	"dynapgen/repository/github"
	"dynapgen/util/cmd"
	"dynapgen/util/file"
	"dynapgen/util/log"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/xid"
)

var (
	errorFailedToGetAppFull = errors.New("error getting app full")
)

func (uc *Usecase) GetAllApps(ctx context.Context, param GetAppListParam) (GetAppListResponse, error) {
	apps, err := uc.db.GetAppsByUserID(ctx, param.OwnerID)
	if err != nil {
		log.Error(err, "uc.db.GetApps() got error - GetAllApp", param)
		return GetAppListResponse{}, err
	}

	result := make([]App, 0, len(apps))
	for _, app := range apps {
		result = append(result, convertAppFromDB(app))
	}

	return buildAppListResponse(result, param), nil
}

func (uc *Usecase) GetApp(ctx context.Context, appID string) (App, error) {
	app, err := uc.db.GetApp(ctx, appID)
	if err != nil {
		log.Error(err, "uc.db.GetApp() got error - GetApp", map[string]interface{}{"app_id": appID})
		return App{}, err
	}
	if app.ID == "" {
		return App{}, errorAppNotFound
	}

	return convertAppFromDB(app), nil
}

func (uc *Usecase) GetAppAds(ctx context.Context, appID string) ([]AppAds, error) {
	ads, err := uc.db.GetAppAdsByAppID(ctx, appID)
	if err != nil {
		log.Error(err, "uc.db.GetAppAdsByAppID() got error - GetAppAds", map[string]interface{}{"app_id": appID})
		return nil, err
	}

	result := make([]AppAds, 0, len(ads))
	for _, ad := range ads {
		result = append(result, AppAds(ad))
	}

	return result, nil
}

// Get App Full for mobile client usage. Aggregating app - ads - content
func (uc *Usecase) GetAppFull(ctx context.Context, appID string) (AppFull, error) {
	cachedApp, err := uc.cache.GetAppFullByID(ctx, appID)
	if err != nil {
		log.Error(err, "uc.cache.GetUserRoleIDMapByUserID() got error - GetAppFull", map[string]interface{}{"app_id": appID})
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
			log.Error(err, "uc.db.GetApp() got error - GetAppFull", map[string]interface{}{"app_id": appID})
			errMsgs = append(errMsgs, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		appAds, err = uc.db.GetAppAdsByAppID(ctx, appID)
		if err != nil {
			log.Error(err, "uc.db.GetAppAdsByAppID() got error - GetAppFull", map[string]interface{}{"app_id": appID})
			errMsgs = append(errMsgs, err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		appContents, err = uc.db.GetAppContentsByAppID(ctx, appID)
		if err != nil {
			log.Error(err, "uc.db.GetAppContentsByAppID() got error - GetAppFull", map[string]interface{}{"app_id": appID})
			errMsgs = append(errMsgs, err.Error())
		}
	}()

	wg.Wait()
	if len(errMsgs) > 0 {
		log.Error(errors.New(strings.Join(errMsgs, ", ")), "error getting app full - GetAppFull", map[string]interface{}{"app_id": appID})
		return AppFull{}, errorFailedToGetAppFull
	}

	result := buildAppFull(app, appAds, appContents)
	err = uc.cache.InsertAppFull(ctx, convertAppFullToCache(result))
	if err != nil {
		log.Error(err, "uc.cache.InsertAppFull() got error - InsertAppFull")
	}

	return result, nil
}

func (uc *Usecase) NewApp(ctx context.Context, param NewAppRequest) error {
	app := App{
		ID:                         xid.New().String(),
		OwnerID:                    param.OwnerID,
		Name:                       param.AppName,
		PackageName:                param.PackageName,
		Version:                    1,
		VersionCode:                "1.0.0",
		IconURL:                    "https://raw.githubusercontent.com/Dynapgen/master-storage-1/main/assets/default-icon.png",
		TemplateID:                 param.TemplateID,
		InterstitialIntervalSecond: 120,
		CreatedAt:                  time.Now(),
	}

	appStrings, err := generateAppStrings(param.TemplateID)
	if err != nil {
		log.Error(err, "generateAppStrings got error - New App", param)
		return err
	}

	appStyles, err := generateAppStyles(param.TemplateID)
	if err != nil {
		log.Error(err, "generateAppStyle got error - New App", param)
		return err
	}

	app.Strings = appStrings
	app.Styles = appStyles

	if err = uc.db.InsertApp(ctx, app.convertAppToDB()); err != nil {
		log.Error(err, "uc.repo.InsertApp() got error - NewApp", param)
		return err
	}

	return nil
}

func (uc *Usecase) NewAppAds(ctx context.Context, request NewAppAdsRequest) error {
	ads := db.AppAds{
		ID:               xid.New().String(),
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
		log.Error(err, "uc.db.InsertAppAds() got error - NewAppAds", request)
		return err
	}

	err = uc.cache.InvalidateAppFull(ctx, request.AppID)
	if err != nil {
		log.Error(err, "uc.cache.InvalidateApp() got error - NewAppAds", map[string]interface{}{})
	}

	return nil
}

func (uc *Usecase) UpdateApp(ctx context.Context, request App) error {
	app, err := uc.GetApp(ctx, request.ID)
	if err != nil {
		log.Error(err, "uc.Get() got error - UpdateApp", app)
	}

	app.updateWith(request)
	param := app.convertAppToDB()
	if err = uc.db.UpdateApp(ctx, param); err != nil {
		log.Error(err, "uc.db.UpdateApp() got error - UpdateApp", param)
		return err
	}

	if err = uc.cache.InvalidateAppFull(ctx, app.ID); err != nil {
		log.Error(err, "uc.cache.InvalidateApp() got error - UpdateApp", app)
	}

	return nil
}

func (uc *Usecase) UpdateAppIcon(ctx context.Context, appID string, iconName, localPath string) error {
	defer func() {
		os.Remove(localPath)
	}()

	meta := map[string]interface{}{
		"app_id":     appID,
		"local_path": localPath,
		"icon_name":  iconName,
	}

	app, err := uc.GetApp(ctx, appID)
	if err != nil {
		log.Error(err, "uc.GetApp() got error - UpdateAppIcon", meta)
		return err
	}

	iconURL, err := uc.github.Upload(ctx, github.UploadFileParam{
		FilePathLocal:       localPath,
		FilePathRemote:      file.GenerateUniqueGithubFilePath(file.GithubFolderIcon, iconName),
		ReplaceIfNameExists: true,
	})
	if err != nil {
		log.Error(err, "uc.repo.UploadToGithub() got error - UpdateAppIcon", meta)
		return err
	}

	timeNow := time.Now()
	app.IconURL = iconURL
	app.UpdatedAt = &timeNow
	param := app.convertAppToDB()
	err = uc.db.UpdateApp(ctx, param)
	if err != nil {
		log.Error(err, "uc.repo.UpdateAppOnDB() got error - UpdateAppIcon", param)
		return err
	}

	cleanupCommand := fmt.Sprintf("rm %s", localPath)
	if err := cmd.ExecCommandWithContext(ctx, cleanupCommand); err != nil {
		log.Error(err, "error cleaning up updated app icon - UpdateAppIcon", cleanupCommand)
		return err
	}

	return nil
}

func (uc *Usecase) DeleteApp(ctx context.Context, appID string) error {
	err := uc.db.DeleteApp(ctx, appID)
	if err != nil {
		log.Error(err, "uc.repo.DeleteApp() - DeleteApp", map[string]interface{}{"app_id": appID})
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

func generateAppStrings(templateID string) (map[string]string, error) {
	if value, found := templateIDStringMap[templateID]; found {
		return value, nil
	}

	return nil, errors.New("template ID not found")
}

func generateAppStyles(templateID string) (map[string]string, error) {
	if value, found := templateIDStyleMap[templateID]; found {
		return value, nil
	}

	return nil, errors.New("template ID not found")
}
