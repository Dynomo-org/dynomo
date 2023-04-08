package usecase

import (
	"context"
	"dynapgen/repository"
	"dynapgen/utils/log"
	"errors"
)

var (
	errorMasterAppNotFound = errors.New("master app not found")
)

func (uc *Usecase) GetAllMasterApp(ctx context.Context) ([]MasterApp, error) {
	apps, err := uc.repo.GetAllMasterAppFromDB(ctx)
	if err != nil {
		log.Error(nil, err, "uc.repo.GetAllMasterAppFromDB() got error - GetAllMasterApp")
		return nil, err
	}

	result := make([]MasterApp, 0, len(apps))
	for _, app := range apps {
		result = append(result, convertMasterAppFromRepo(app))
	}

	return result, nil
}

func (uc *Usecase) GetMasterApp(ctx context.Context, appID string) (MasterApp, error) {
	cached, err := uc.repo.GetMasterAppFromCache(ctx, appID)
	meta := map[string]interface{}{"app_id": appID}
	if err != nil {
		log.Error(meta, err, "uc.repo.GetMasterAppFromCache() got error - GetMasterApp")
		return MasterApp{}, err
	}

	masterApp := cached
	if masterApp.AppID == "" {
		masterAppFromDB, err := uc.repo.GetMasterAppFromDB(ctx, appID)
		if err != nil {
			log.Error(meta, err, "uc.repo.GetMasterAppFromDB() got error - GetMasterApp")
			return MasterApp{}, err
		}
		if masterAppFromDB.AppID == "" {
			log.Error(meta, errorMasterAppNotFound, "master app not found - GetMasterApp")
			return MasterApp{}, errorMasterAppNotFound
		}
		masterApp = masterAppFromDB
	}

	return convertMasterAppFromRepo(masterApp), nil
}

func (uc *Usecase) NewMasterApp(ctx context.Context, name string) (string, error) {
	id, err := uc.repo.InsertNewMasterAppToDB(ctx, name)
	if err != nil {
		log.Error(map[string]interface{}{"name": name}, err, "uc.repo.InsertNewMasterAppToDB() got error - v")
		return "", nil
	}

	return id, nil
}

func (uc *Usecase) SaveMasterApp(ctx context.Context, masterApp MasterApp) error {
	err := uc.repo.InsertMasterAppToDB(ctx, convertMasterAppToRepo(masterApp))
	if err != nil {
		log.Error(masterApp, err, "uc.repo.InsertMasterAppToDB() got error - SaveMasterApp")
		return err
	}

	return nil
}

func (uc *Usecase) UpdateMasterApp(ctx context.Context, masterApp MasterApp) error {
	param := convertMasterAppToRepo(masterApp)
	err := uc.repo.UpdateMasterAppOnDB(ctx, param)
	if err != nil {
		log.Error(param, err, "uc.repo.UpdateMasterAppOnDB() got error - UpdateMasterApp")
		return err
	}

	err = uc.repo.InvalidateMasterAppOnCache(ctx, param.AppID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": param.AppID}, err, "uc.repo.InvalidateMasterAppOnCache() got error - UpdateMasterApp")
		return err
	}

	return nil
}

func convertMasterAppFromRepo(app repository.MasterApp) MasterApp {
	contents := make([]AppContent, 0, len(app.Contents))
	for _, content := range app.Contents {
		contents = append(contents, AppContent(content))
	}

	categories := make([]AppCategory, 0, len(app.Categories))
	for _, category := range app.Categories {
		categories = append(categories, AppCategory(category))
	}

	return MasterApp{
		AppID:     app.AppID,
		Name:      app.Name,
		AdsConfig: AdsConfig(app.AdsConfig),
		AppConfig: AppConfig{
			AppName:           app.AppConfig.AppName,
			ExitPromptMessage: app.AppConfig.ExitPromptMessage,
			Style:             AppStyle(app.AppConfig.Style),
		},
		Contents:   contents,
		Categories: categories,
	}
}

func convertMasterAppToRepo(app MasterApp) repository.MasterApp {
	contents := make([]repository.AppContent, 0, len(app.Contents))
	for _, content := range app.Contents {
		contents = append(contents, repository.AppContent(content))
	}

	categories := make([]repository.AppCategory, 0, len(app.Categories))
	for _, category := range app.Categories {
		categories = append(categories, repository.AppCategory(category))
	}

	return repository.MasterApp{
		AppID:     app.AppID,
		Name:      app.Name,
		AdsConfig: repository.AdsConfig(app.AdsConfig),
		AppConfig: repository.AppConfig{
			AppName:           app.AppConfig.AppName,
			ExitPromptMessage: app.AppConfig.ExitPromptMessage,
			Style:             repository.AppStyle(app.AppConfig.Style),
		},
		Contents:   contents,
		Categories: categories,
	}
}