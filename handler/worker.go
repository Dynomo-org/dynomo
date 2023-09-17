package handler

import (
	"dynapgen/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleUpdateBuildAppStatus(ctx *gin.Context) {
	appID := ctx.Query("app_id")
	if appID == "" {
		WriteJson(ctx, nil, errorAppIDEmpty, http.StatusBadRequest)
		return
	}

	var body BuildStatusPayload
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		WriteJson(ctx, nil, err, http.StatusBadRequest)
		return
	}

	if err = h.usecase.SetBuildAppStatus(ctx, usecase.UpdateBuildStatusParam{
		AppID: appID,
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
	appID := ctx.Query("app_id")
	if appID == "" {
		WriteJson(ctx, nil, errorAppIDEmpty, http.StatusBadRequest)
		return
	}

	var body BuildStatusPayload
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		WriteJson(ctx, nil, err, http.StatusBadRequest)
		return
	}

	if err = h.usecase.SetBuildKeystoreStatus(ctx, usecase.UpdateBuildStatusParam{
		AppID: appID,
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
