package handler

import (
	"dynapgen/constants"
	"dynapgen/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleGenerateKeystore(ctx *gin.Context) {
	var param GenerateStoreParam
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		WriteJson(ctx, nil, err, http.StatusBadRequest)
		return
	}

	userID := ctx.GetString(constants.ContextKeyUserID)
	err = h.usecase.BuildKeystore(ctx, usecase.BuildKeystoreParam{
		OwnerID:       userID,
		KeystoreName:  param.KeystoreName,
		FullName:      param.FullName,
		Organization:  param.Organization,
		Country:       param.Country,
		Alias:         param.Alias,
		KeyPassword:   param.KeyPassword,
		StorePassword: param.StorePassword,
	})
	if err != nil {
		WriteJson(ctx, param, err)
		return
	}

	WriteJson(ctx, nil, nil, http.StatusAccepted)
}

func (h *Handler) HandleGetAllKeystores(ctx *gin.Context) {
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

	param := usecase.GetKeystoreListParam{
		Page:    page,
		PerPage: perPage,
		OwnerID: userID,
	}
	result, err := h.usecase.GetKeystoreList(ctx, param)
	if err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, result, nil)
}
