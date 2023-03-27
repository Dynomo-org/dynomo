package handler

import (
	"dynapgen/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository *repository.Repository
}

func NewHandler(repository *repository.Repository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, defaultSuccessResponse)
}

func (h *Handler) GetAllCollection(ctx *gin.Context) {
	data, err := h.repository.GetAllCollection(ctx)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
		"err":  err,
	})
}

func (h *Handler) CreateNewMasterApp(ctx *gin.Context) {
	var request NewAppRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if err := h.repository.CreateNewMasterApp(ctx, request.AppName); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, defaultSuccessResponse)
}
