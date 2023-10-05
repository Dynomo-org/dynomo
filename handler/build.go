package handler

import (
	"dynapgen/constants"
	"dynapgen/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleBuildApp(ctx *gin.Context) {
	appID := ctx.PostForm("app_id")
	if appID == "" {
		WriteJson(ctx, nil, errorAppIDEmpty, http.StatusBadRequest)
		return
	}

	var param BuildAppParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		WriteJson(ctx, nil, err, http.StatusBadRequest)
		return
	}

	err := h.usecase.BuildApp(ctx, usecase.BuildAppParam{
		AppID:      appID,
		KeystoreID: param.KeystoreID,
	})
	if err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil, http.StatusAccepted)
}

func (h *Handler) HandleGetBuildArtifacts(ctx *gin.Context) {
	appID := ctx.Query("app_id")
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

	param := usecase.GetBuildArtifactsParam{
		Page:    page,
		PerPage: perPage,
		AppID:   appID,
		OwnerID: userID,
	}

	fn := h.usecase.GetBuildArtifactsByAppID
	if appID != "" {
		fn = h.usecase.GetBuildArtifactsByOwnerID
	}

	result, err := fn(ctx, param)
	if err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, result, nil)
}
