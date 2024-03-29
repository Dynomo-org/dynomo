package handler

import (
	"dynapgen/constants"
	"dynapgen/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrz1836/go-sanitize"
)

var (
	errorIncompleteInput = errors.New("incomplete input")
)

func (h *Handler) HandleGetUserInfo(ctx *gin.Context) {
	userID := ctx.GetString(constants.ContextKeyUserID)
	result, err := h.usecase.GetUserInfo(ctx, userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == usecase.ErrorUserNotFound {
			statusCode = http.StatusBadRequest
		}
		WriteJson(ctx, nil, err, statusCode)
		return
	}

	WriteJson(ctx, result, nil)
}

func (h *Handler) HandleLoginUser(ctx *gin.Context) {
	var request LoginUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		WriteJson(ctx, nil, err, http.StatusBadRequest)
		return
	}

	if request.Email == "" || request.Password == "" {
		WriteJson(ctx, nil, errorIncompleteInput, http.StatusBadRequest)
		return
	}

	param := usecase.User{
		Email:    sanitize.Email(request.Email, true),
		Password: request.Password,
	}
	result, err := h.usecase.LoginUser(ctx, param)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == usecase.ErrorUserNotFound {
			statusCode = http.StatusUnauthorized
		}
		WriteJson(ctx, nil, err, statusCode)
		return
	}

	WriteJson(ctx, result, nil)
}

func (h *Handler) HandleRegisterUser(ctx *gin.Context) {
	var request RegisterUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		WriteJson(ctx, nil, err, http.StatusBadRequest)
		return
	}

	if request.FullName == "" || request.Email == "" || request.Password == "" {
		WriteJson(ctx, nil, errorIncompleteInput, http.StatusBadRequest)
		return
	}

	param := usecase.User{
		FullName: request.FullName,
		Email:    sanitize.Email(request.Email, true),
		Password: request.Password,
	}
	result, err := h.usecase.RegisterUser(ctx, param)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == usecase.ErrorUserExists {
			statusCode = http.StatusConflict
		}
		WriteJson(ctx, nil, err, statusCode)
		return
	}

	WriteJson(ctx, result, nil)
}
