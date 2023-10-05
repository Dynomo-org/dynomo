package handler

import (
	"dynapgen/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrz1836/go-sanitize"
)

var (
	errorBuildIDEmpty = errors.New("build id is empty")
)

func (h *Handler) HandleGetBuildAppStatus(ctx *gin.Context) {
	buildID := sanitize.AlphaNumeric(ctx.Query("build_id"), false)
	if buildID == "" {
		WriteJson(ctx, nil, errorBuildIDEmpty, http.StatusBadRequest)
		return
	}

	result, err := h.usecase.GetBuildAppStatus(ctx, buildID)
	if err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, result, nil)
}

func (h *Handler) HandleGetBuildKeystoreStatus(ctx *gin.Context) {
	buildID := sanitize.AlphaNumeric(ctx.Query("build_id"), false)
	if buildID == "" {
		WriteJson(ctx, nil, errorBuildIDEmpty, http.StatusBadRequest)
		return
	}

	result, err := h.usecase.GetBuildKeystoreStatus(ctx, buildID)
	if err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, result, nil)
}

func (h *Handler) HandleUpdateBuildAppStatus(ctx *gin.Context) {
	buildID := sanitize.AlphaNumeric(ctx.Query("build_id"), false)
	if buildID == "" {
		WriteJson(ctx, nil, errorBuildIDEmpty, http.StatusBadRequest)
		return
	}

	var body BuildStatusPayload
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		WriteJson(ctx, nil, err, http.StatusBadRequest)
		return
	}

	if err = h.usecase.SetBuildAppStatus(ctx, usecase.UpdateBuildStatusParam{
		BuildID: buildID,
		BuildStatus: usecase.BuildStatus{
			Status:       usecase.BuildStatusEnum(body.Status),
			URL:          body.URL,
			ErrorMessage: body.ErrorMessage,
		},
	}); err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}

func (h *Handler) HandleUpdateBuildKeystoreStatus(ctx *gin.Context) {
	buildID := sanitize.AlphaNumeric(ctx.Query("build_id"), false)
	if buildID == "" {
		return
	}

	var body BuildStatusPayload
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		WriteJson(ctx, nil, err, http.StatusBadRequest)
		return
	}

	if err = h.usecase.SetBuildKeystoreStatus(ctx, usecase.UpdateBuildStatusParam{
		BuildID: buildID,
		BuildStatus: usecase.BuildStatus{
			Status:       usecase.BuildStatusEnum(body.Status),
			URL:          body.URL,
			ErrorMessage: body.ErrorMessage,
		},
	}); err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil)
}
