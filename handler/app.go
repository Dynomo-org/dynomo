package handler

import (
	"dynapgen/constants"
	"dynapgen/usecase"
	"dynapgen/utils/log"
	"errors"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	errorAppIDEmpty = errors.New("app_id is empty")

	supportedFileTypes = map[string]struct{}{"png": {}, "jpg": {}, "jpeg": {}}
)

func (h *Handler) HandleCreateNewApp(ctx *gin.Context) {
	var request NewAppRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Error(nil, err, "ctx.BindJSON() got error - HandleCreateNewApp")
		WriteJson(ctx, nil, err)
		return
	}

	ownerID := ctx.GetString(constants.ContextKeyUserID)
	param := usecase.NewAppRequest{
		AppName:     request.AppName,
		PackageName: request.PackageName,
		OwnerID:     ownerID,
	}
	err = h.usecase.NewApp(ctx, param)
	if err != nil {
		log.Error(request, err, "h.usecase.NewApp() got error - HandleCreateNewApp")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}

func (h *Handler) HandleCreateNewAds(ctx *gin.Context) {
	var request NewAppAdsRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Error(nil, err, "ctx.BindJSON() got error - HandleCreateNewAds")
		WriteJson(ctx, nil, err)
		return
	}

	param := usecase.NewAppAdsRequest{
		AppID:            request.AppID,
		Type:             constants.AdType(request.Type),
		OpenAdID:         request.OpenAdID,
		BannerAdID:       request.BannerAdID,
		InterstitialAdID: request.InterstitialAdID,
		RewardAdID:       request.RewardAdID,
		NativeAdID:       request.NativeAdID,
	}
	err = h.usecase.NewAppAds(ctx, param)
	if err != nil {
		log.Error(request, err, "h.usecase.NewAppAds() got error - HandleCreateNewAds")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}

func (h *Handler) HandleGetAllApps(ctx *gin.Context) {
	userID := ctx.GetString(constants.ContextKeyUserID)
	perPageStr := ctx.Query("per_page")
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		perPage = defaultPerPage
	}

	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = defaultPage
	}

	param := usecase.GetAppListParam{
		Page:    page,
		PerPage: perPage,
		OwnerID: userID,
	}
	result, err := h.usecase.GetAllApps(ctx, param)
	if err != nil {
		log.Error(nil, err, "h.usecase.GetAllApps() got error - HandleGetAllApps")
		WriteJson(ctx, nil, err)
		return
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

	WriteJson(ctx, result, nil)
}

func (h *Handler) HandleGetAppAds(ctx *gin.Context) {
	appID := ctx.Query("id")
	if appID == "" {
		WriteJson(ctx, nil, errorAppIDEmpty, http.StatusBadRequest)
		return
	}
	result, err := h.usecase.GetAppAds(ctx, appID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": appID}, err, "h.usecase.GetAppAds() got error - HandleGetAppAds")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, result, nil)
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

	err = h.usecase.UpdateApp(ctx, usecase.App(request))
	if err != nil {
		log.Error(request, err, "h.usecase.UpdateApp() got error - HandleUpdateApp")
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
