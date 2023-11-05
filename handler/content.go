package handler

import (
	// golang package
	"dynapgen/usecase"
	"dynapgen/util/log"
	"net/http"

	// external package
	"github.com/gin-gonic/gin"
	"github.com/mrz1836/go-sanitize"
)

func (h *Handler) HandleCreateNewContent(ctx *gin.Context) {
	var payload AppContent
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Error(err, "ctx.ShouldBindJSON() got error - HandleCreateNewContent")
		WriteJson(ctx, nil, err)
		return
	}

	if err := h.usecase.NewAppContent(ctx, usecase.AppContent(payload)); err != nil {
		log.Error(err, "h.usecase.NewAppContent() got error - HandleCreateNewContent")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}

func (h *Handler) HandleGetAllContents(ctx *gin.Context) {
	appID := sanitize.AlphaNumeric(ctx.Query("app_id"), false)
	if appID == "" {
		WriteJson(ctx, nil, errorAppIDEmpty, http.StatusBadRequest)
		return
	}
	result, err := h.usecase.GetAppContentsByAppID(ctx, appID)
	if err != nil {
		log.Error(err, "h.usecase.GetApp() got error - HandleGetAllContents", map[string]interface{}{"app_id": appID})
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, result, nil)
}

func (h *Handler) HandleGetContentDetail(ctx *gin.Context) {
	contentID := sanitize.AlphaNumeric(ctx.Query("content_id"), false)
	if contentID == "" {
		WriteJson(ctx, nil, errorContentIDEmpty, http.StatusBadRequest)
		return
	}
	result, err := h.usecase.GetAppContentByID(ctx, contentID)
	if err != nil {
		log.Error(err, "h.usecase.GetAppContentByID() got error - HandleGetContentDetail", map[string]interface{}{"content_id": contentID})
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, result, nil)
}

func (h *Handler) HandleUpdateContent(ctx *gin.Context) {
	var payload AppContent
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Error(err, "ctx.ShouldBindJSON() got error - HandleCreateNewContent")
		WriteJson(ctx, nil, err)
		return
	}

	if err := h.usecase.UpdateAppContent(ctx, usecase.AppContent(payload)); err != nil {
		log.Error(err, "h.usecase.UpdateAppContent() got error - HandleCreateNewContent")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}
