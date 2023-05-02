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
		err = uc.repo.StoreMasterAppToCache(ctx, masterApp)
		if err != nil {
			log.Error(meta, err, "uc.repo.StoreMasterAppToCache() got error - GetMasterApp")
			return MasterApp{}, err
		}
	}

	return convertMasterAppFromRepo(masterApp), nil
}

func (uc *Usecase) NewMasterApp(ctx context.Context, request NewMasterAppRequest) error {
	err := uc.repo.InsertNewMasterAppToDB(ctx, repository.NewMasterAppRequest(request))
	if err != nil {
		log.Error(request, err, "uc.repo.InsertNewMasterAppToDB() got error - NewMasterApp")
		return err
	}

	return nil
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
		AppID:   app.AppID,
		AppName: app.AppName,
		AdsConfig: AdsConfig{
			EnableOpenAd:               app.AdsConfig.EnableOpenAd,
			EnableBannerAd:             app.AdsConfig.EnableBannerAd,
			EnableInterstitialAd:       app.AdsConfig.EnableInterstitialAd,
			EnableRewardAd:             app.AdsConfig.EnableRewardAd,
			EnableNativeAd:             app.AdsConfig.EnableNativeAd,
			PrimaryAd:                  Ad(app.AdsConfig.PrimaryAd),
			SecondaryAd:                Ad(app.AdsConfig.SecondaryAd),
			TertiaryAd:                 Ad(app.AdsConfig.TertiaryAd),
			InterstitialIntervalSecond: app.AdsConfig.InterstitialIntervalSecond,
			TestDevices:                app.AdsConfig.TestDevices,
		},
		AppConfig: AppConfig{
			Strings: AppString(app.AppConfig.Strings),
			Style:   AppStyle(app.AppConfig.Style),
		},
		Contents:       contents,
		Categories:     categories,
		AppPackageName: app.AppPackageName,
		CreatedAt:      app.CreatedAt,
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
		AppID:   app.AppID,
		AppName: app.AppName,
		AdsConfig: repository.AdsConfig{
			EnableOpenAd:               app.AdsConfig.EnableOpenAd,
			EnableBannerAd:             app.AdsConfig.EnableBannerAd,
			EnableInterstitialAd:       app.AdsConfig.EnableInterstitialAd,
			EnableRewardAd:             app.AdsConfig.EnableRewardAd,
			EnableNativeAd:             app.AdsConfig.EnableNativeAd,
			PrimaryAd:                  repository.Ad(app.AdsConfig.PrimaryAd),
			SecondaryAd:                repository.Ad(app.AdsConfig.SecondaryAd),
			TertiaryAd:                 repository.Ad(app.AdsConfig.TertiaryAd),
			InterstitialIntervalSecond: app.AdsConfig.InterstitialIntervalSecond,
			TestDevices:                app.AdsConfig.TestDevices,
		},
		AppConfig: repository.AppConfig{
			Strings: repository.AppString(app.AppConfig.Strings),
			Style:   repository.AppStyle(app.AppConfig.Style),
		},
		Contents:   contents,
		Categories: categories,
	}
}
