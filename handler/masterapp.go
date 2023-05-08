package handler

import (
	"dynapgen/usecase"
	"dynapgen/utils/log"
	"errors"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	errorAppIDEmpty = errors.New("app_id is empty")

	supportedFileTypes = map[string]struct{}{"png": struct{}{}, "jpg": struct{}{}, "jpeg": struct{}{}}
)

func (h *Handler) HandleCreateNewApp(ctx *gin.Context) {
	var request NewAppRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Error(nil, err, "ctx.BindJSON() got error - HandleCreateNewApp")
		WriteJson(ctx, nil, err)
		return
	}

	err = h.usecase.NewApp(ctx, usecase.NewAppRequest(request))
	if err != nil {
		log.Error(request, err, "h.usecase.NewApp() got error - HandleCreateNewApp")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}

func (h *Handler) HandleGetAllApp(ctx *gin.Context) {
	apps, err := h.usecase.GetAllApp(ctx)
	if err != nil {
		log.Error(nil, err, "h.usecase.GetAllApp() got error - HandleGetAllApp")
		WriteJson(ctx, nil, err)
		return
	}

	result := make([]App, 0, len(apps))
	for _, app := range apps {
		result = append(result, convertAppFromUsecase(app))
	}

	WriteJson(ctx, result, nil)
}

func (h *Handler) HandleGetApp(ctx *gin.Context) {
	appID := ctx.Query("id")
	if appID == "" {
		WriteJson(ctx, nil, errorAppIDEmpty, http.StatusBadRequest)
		return
	}
	result, err := h.usecase.GetApp(ctx, appID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": appID}, err, "h.usecase.GetApp() got error - HandleGetApp")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, convertAppFromUsecase(result), nil)
}

func (h *Handler) HandleDeleteApp(ctx *gin.Context) {
	appID := ctx.Query("id")
	if appID == "" {
		WriteJson(ctx, nil, errorAppIDEmpty, http.StatusBadRequest)
		return
	}
	err := h.usecase.DeleteApp(ctx, appID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": appID}, err, "h.usecase.DeleteApp() got error - HandleDeleteApp")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}

func (h *Handler) HandleUpdateApp(ctx *gin.Context) {
	var request App
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Error(nil, err, "ctx.BindJSON() got error - HandleUpdateApp")
		WriteJson(ctx, nil, err)
		return
	}

	param := convertAppToUsecase(request)
	err = h.usecase.UpdateApp(ctx, param)
	if err != nil {
		log.Error(param, err, "h.usecase.UpdateApp() got error - HandleUpdateApp")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}

func (h *Handler) HandleUpdateAppIcon(ctx *gin.Context) {
	appID := ctx.PostForm("app_id")
	if appID == "" {
		WriteJson(ctx, nil, errorAppIDEmpty, http.StatusBadRequest)
		return
	}

	file, err := ctx.FormFile("icon")
	if err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	if file.Size >= 1*1024*1024 {
		WriteJson(ctx, nil, errors.New("file size is too large"), http.StatusBadRequest)
		return
	}

	filenameSegments := strings.Split(file.Filename, ".")
	if len(filenameSegments) < 2 {
		WriteJson(ctx, nil, errors.New("file extension not found"), http.StatusBadRequest)
		return
	}

	fileExt := filenameSegments[len(filenameSegments)-1]
	if _, ok := supportedFileTypes[strings.ToLower(fileExt)]; !ok {
		WriteJson(ctx, nil, errors.New("file extension not supported"), http.StatusBadRequest)
		return
	}

	file.Filename = "app_icon." + filenameSegments[len(filenameSegments)-1]
	dst := filepath.Join("./assets/" + file.Filename)
	ctx.SaveUploadedFile(file, dst)

	err = h.usecase.UpdateAppIcon(ctx, appID, file.Filename, dst)
	if err != nil {
		log.Error(nil, err, "h.usecase.UpdateAppIcon() got error - HandleUpdateAppIcon")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}

func convertAppFromUsecase(app usecase.App) App {
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

func convertAppToUsecase(app App) usecase.App {
	contents := make([]usecase.AppContent, 0, len(app.Contents))
	for _, content := range app.Contents {
		contents = append(contents, usecase.AppContent(content))
	}

	categories := make([]usecase.AppCategory, 0, len(app.Categories))
	for _, category := range app.Categories {
		categories = append(categories, usecase.AppCategory(category))
	}

	return usecase.App{
		AppID:             app.AppID,
		AppName:           app.AppName,
		AppPackageName:    app.AppPackageName,
		VersionCode:       app.VersionCode,
		VersionName:       app.VersionName,
		IconURL:           app.IconURL,
		PrivacyPolicyLink: app.PrivacyPolicyLink,
		AdsConfig: usecase.AdsConfig{
			EnableOpenAd:               app.AdsConfig.EnableOpenAd,
			EnableBannerAd:             app.AdsConfig.EnableBannerAd,
			EnableInterstitialAd:       app.AdsConfig.EnableInterstitialAd,
			EnableRewardAd:             app.AdsConfig.EnableRewardAd,
			EnableNativeAd:             app.AdsConfig.EnableNativeAd,
			PrimaryAd:                  usecase.Ad(app.AdsConfig.PrimaryAd),
			SecondaryAd:                usecase.Ad(app.AdsConfig.SecondaryAd),
			TertiaryAd:                 usecase.Ad(app.AdsConfig.TertiaryAd),
			InterstitialIntervalSecond: app.AdsConfig.InterstitialIntervalSecond,
			TestDevices:                app.AdsConfig.TestDevices,
		},
		AppConfig: usecase.AppConfig{
			Strings: usecase.AppString(app.AppConfig.Strings),
			Style:   usecase.AppStyle(app.AppConfig.Style),
		},
		Contents:   contents,
		Categories: categories,
	}
}
