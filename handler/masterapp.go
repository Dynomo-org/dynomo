package handler

import (
	"dynapgen/usecase"
	"dynapgen/utils/log"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	errorAppIDEmpty = errors.New("app_id is empty")
)

func (h *Handler) HandleCreateNewMasterApp(ctx *gin.Context) {
	var request NewMasterAppRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Error(nil, err, "ctx.BindJSON() got error - HandleCreateNewMasterApp")
		WriteJson(ctx, nil, err)
		return
	}

	err = h.usecase.NewMasterApp(ctx, usecase.NewMasterAppRequest(request))
	if err != nil {
		log.Error(request, err, "h.usecase.NewMasterApp() got error - HandleCreateNewMasterApp")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}

func (h *Handler) HandleGetAllMasterApp(ctx *gin.Context) {
	apps, err := h.usecase.GetAllMasterApp(ctx)
	if err != nil {
		log.Error(nil, err, "h.usecase.GetAllMasterApp() got error - HandleGetAllMasterApp")
		WriteJson(ctx, nil, err)
		return
	}

	result := make([]MasterApp, 0, len(apps))
	for _, app := range apps {
		result = append(result, convertMasterAppFromUsecase(app))
	}

	WriteJson(ctx, result, nil)
}

func (h *Handler) HandleGetMasterApp(ctx *gin.Context) {
	appID := ctx.Query("id")
	if appID == "" {
		WriteJson(ctx, nil, errorAppIDEmpty, http.StatusBadRequest)
		return
	}
	result, err := h.usecase.GetMasterApp(ctx, appID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": appID}, err, "h.usecase.GetMasterApp() got error - HandleGetMasterApp")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, convertMasterAppFromUsecase(result), nil)
}

func (h *Handler) HandleUpdateMasterApp(ctx *gin.Context) {
	var request MasterApp
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Error(nil, err, "ctx.BindJSON() got error - HandleUpdateMasterApp")
		WriteJson(ctx, nil, err)
		return
	}

	param := convertMasterAppToUsecase(request)
	err = h.usecase.UpdateMasterApp(ctx, param)
	if err != nil {
		log.Error(param, err, "h.usecase.UpdateMasterApp() got error - HandleUpdateMasterApp")
		WriteJson(ctx, nil, err)
		return
	}
}

func convertMasterAppFromUsecase(app usecase.MasterApp) MasterApp {
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

func convertMasterAppToUsecase(app MasterApp) usecase.MasterApp {
	contents := make([]usecase.AppContent, 0, len(app.Contents))
	for _, content := range app.Contents {
		contents = append(contents, usecase.AppContent(content))
	}

	categories := make([]usecase.AppCategory, 0, len(app.Categories))
	for _, category := range app.Categories {
		categories = append(categories, usecase.AppCategory(category))
	}

	return usecase.MasterApp{
		AppID:   app.AppID,
		AppName: app.AppName,
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
