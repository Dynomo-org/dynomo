package handler

import (
	"dynapgen/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	keystoreFileTypes = map[string]struct{}{
		"jks":      {},
		"keystore": {},
	}
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
