package handler

import (
	"dynapgen/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleGenerateKeystore(ctx *gin.Context) {
	var param GenerateStoreParam
	err := ctx.BindJSON(&param)
	if err != nil {
		WriteJson(ctx, nil, err, http.StatusBadRequest)
	}

	err = h.usecase.GenerateKeystore(ctx, usecase.GenerateStoreParam(param))
	if err != nil {
		WriteJson(ctx, param, err)
	}

	WriteJson(ctx, nil, nil)
}

func (h *Handler) HandleGetGenerateKeystoreStatus(ctx *gin.Context) {
	appID := ctx.Query("app_id")
	if appID == "" {
		WriteJson(ctx, nil, errorAppIDEmpty, http.StatusBadRequest)
	}

	status, err := h.usecase.GetGenerateKeystoreStatus(ctx, appID)
	if err != nil {
		WriteJson(ctx, map[string]interface{}{"app_id": appID}, err)
	}

	WriteJson(ctx, Keystore(status), nil)
}
