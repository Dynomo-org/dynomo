package handler

import (
	"dynapgen/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleGenerateKeystore(ctx *gin.Context) {
	var param GenerateStoreParam
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		WriteJson(ctx, nil, err, http.StatusBadRequest)
		return
	}

	err = h.usecase.BuildKeystore(ctx, usecase.BuildKeystoreParam(param))
	if err != nil {
		WriteJson(ctx, param, err)
		return
	}

	WriteJson(ctx, nil, nil, http.StatusAccepted)
}
