package usecase

import (
	"context"
	"dynapgen/repository"
	"dynapgen/utils/log"
	"errors"
	"os"
	"time"
)

var (
	errorAppNotFound = errors.New("master app not found")
)

func (uc *Usecase) GetAllApp(ctx context.Context) ([]App, error) {
	apps, err := uc.repo.GetAllAppFromDB(ctx)
	if err != nil {
		log.Error(nil, err, "uc.repo.GetAllAppFromDB() got error - GetAllApp")
		return nil, err
	}

	result := make([]App, 0, len(apps))
	for _, app := range apps {
		result = append(result, convertAppFromRepo(app))
	}

	return result, nil
}

func (uc *Usecase) GetApp(ctx context.Context, appID string) (App, error) {
	cached, err := uc.repo.GetAppFromCache(ctx, appID)
	meta := map[string]interface{}{"app_id": appID}
	if err != nil {
		log.Error(meta, err, "uc.repo.GetAppFromCache() got error - GetApp")
		return App{}, err
	}

	app := cached
	if app.AppID == "" {
		appFromDB, err := uc.repo.GetAppFromDB(ctx, appID)
		if err != nil {
			log.Error(meta, err, "uc.repo.GetAppFromDB() got error - GetApp")
			return App{}, err
		}
		if appFromDB.AppID == "" {
			log.Error(meta, errorAppNotFound, "master app not found - GetApp")
			return App{}, errorAppNotFound
		}

		app = appFromDB
		err = uc.repo.StoreAppToCache(ctx, app)
		if err != nil {
			log.Error(meta, err, "uc.repo.StoreAppToCache() got error - GetApp")
			return App{}, err
		}
	}

	return convertAppFromRepo(app), nil
}

func (uc *Usecase) NewApp(ctx context.Context, request NewAppRequest) error {
	err := uc.repo.InsertNewAppToDB(ctx, repository.NewAppRequest(request))
	if err != nil {
		log.Error(request, err, "uc.repo.InsertNewAppToDB() got error - NewApp")
		return err
	}

	return nil
}

func (uc *Usecase) SaveApp(ctx context.Context, app App) error {
	err := uc.repo.InsertAppToDB(ctx, app.convertAppToRepo())
	if err != nil {
		log.Error(app, err, "uc.repo.InsertAppToDB() got error - SaveApp")
		return err
	}

	return nil
}

func (uc *Usecase) UpdateApp(ctx context.Context, request App) error {
	app, err := uc.GetApp(ctx, request.AppID)
	if err != nil {
		log.Error(app, err, "uc.Get() got error - UpdateApp")
	}

	app.updateWith(request)
	timeNow := time.Now()
	param := app.convertAppToRepo()
	param.UpdatedAt = &timeNow
	err = uc.repo.UpdateAppOnDB(ctx, param)
	if err != nil {
		log.Error(param, err, "uc.repo.UpdateAppOnDB() got error - UpdateApp")
		return err
	}

	err = uc.repo.StoreAppToCache(ctx, param)
	if err != nil {
		log.Error(param, err, "uc.repo.StoreAppToCache() got error - UpdateApp")
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

	iconURL, err := uc.repo.UploadToGithub(ctx, repository.UploadFileParam{
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
	param := app.convertAppToRepo()
	err = uc.repo.UpdateAppOnDB(ctx, param)
	if err != nil {
		log.Error(param, err, "uc.repo.UpdateAppOnDB() got error - UpdateAppIcon")
		return err
	}

	err = uc.repo.StoreAppToCache(ctx, param)
	if err != nil {
		log.Error(param, err, "uc.repo.StoreAppToCache() got error - UpdateApp")
		return err
	}

	return nil
}

func (uc *Usecase) DeleteApp(ctx context.Context, appID string) error {
	err := uc.repo.DeleteAppOnDB(ctx, appID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": appID}, err, "uc.repo.DeleteAppOnDB() - DeleteApp")
		return err
	}

	return nil
}
