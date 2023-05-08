package usecase

import (
	"bytes"
	"context"
	"dynapgen/repository"
	"dynapgen/utils/log"
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
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

func (uc *Usecase) SaveApp(ctx context.Context, App App) error {
	err := uc.repo.InsertAppToDB(ctx, convertAppToRepo(App))
	if err != nil {
		log.Error(App, err, "uc.repo.InsertAppToDB() got error - SaveApp")
		return err
	}

	return nil
}

func (uc *Usecase) UpdateApp(ctx context.Context, app App) error {
	param := convertAppToRepo(app)
	err := uc.repo.UpdateAppOnDB(ctx, param)
	if err != nil {
		log.Error(param, err, "uc.repo.UpdateAppOnDB() got error - UpdateApp")
		return err
	}

	err = uc.repo.InvalidateAppOnCache(ctx, param.AppID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": param.AppID}, err, "uc.repo.InvalidateAppOnCache() got error - UpdateApp")
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

	app.IconURL = iconURL
	param := convertAppToRepo(app)
	err = uc.repo.UpdateAppOnDB(ctx, param)
	if err != nil {
		log.Error(param, err, "uc.repo.UpdateAppOnDB() got error - UpdateAppIcon")
		return err
	}

	err = uc.repo.InvalidateAppOnCache(ctx, param.AppID)
	if err != nil {
		log.Error(meta, err, "uc.repo.InvalidateAppOnCache() got error - UpdateApp")
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

func convertAppFromRepo(app repository.App) App {
	contents := make([]AppContent, 0, len(app.Contents))
	for _, content := range app.Contents {
		contents = append(contents, AppContent(content))
	}

	categories := make([]AppCategory, 0, len(app.Categories))
	for _, category := range app.Categories {
		categories = append(categories, AppCategory(category))
	}

	return App{
		AppID:             app.AppID,
		AppName:           app.AppName,
		AppPackageName:    app.AppPackageName,
		VersionCode:       app.VersionCode,
		VersionName:       app.VersionName,
		IconURL:           app.IconURL,
		PrivacyPolicyLink: app.PrivacyPolicyLink,
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
		Contents:   contents,
		Categories: categories,
		CreatedAt:  app.CreatedAt,
	}
}

func convertAppToRepo(app App) repository.App {
	contents := make([]repository.AppContent, 0, len(app.Contents))
	for _, content := range app.Contents {
		contents = append(contents, repository.AppContent(content))
	}

	categories := make([]repository.AppCategory, 0, len(app.Categories))
	for _, category := range app.Categories {
		categories = append(categories, repository.AppCategory(category))
	}

	return repository.App{
		AppID:             app.AppID,
		AppName:           app.AppName,
		AppPackageName:    app.AppPackageName,
		VersionCode:       app.VersionCode,
		VersionName:       app.VersionName,
		IconURL:           app.IconURL,
		PrivacyPolicyLink: app.PrivacyPolicyLink,
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

func sanitizeFileToPNG(fileName, path string) (string, string, error) {
	fileNameSegments := strings.Split(fileName, ".")
	if strings.ToLower(fileNameSegments[1]) == "png" {
		return fileName, path, nil
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}

	img, err := jpeg.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		return "", "", err
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return "", "", err
	}

	os.Remove(path)

	newPath := "./assets/" + fileNameSegments[0] + ".png"
	os.WriteFile(newPath, buf.Bytes(), 0644)

	return fileNameSegments[0] + ".png", newPath, nil
}
