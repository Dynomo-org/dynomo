package handler

import (
	"context"
	"dynapgen/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usecaseProvider interface {
	GetAllMasterApp(ctx context.Context) ([]usecase.MasterApp, error)
	GetMasterApp(ctx context.Context, appID string) (usecase.MasterApp, error)
	NewMasterApp(ctx context.Context, name string) (string, error)
	SaveMasterApp(ctx context.Context, masterApp usecase.MasterApp) error
	UpdateMasterApp(ctx context.Context, masterApp usecase.MasterApp) error
}

type Handler struct {
	usecase usecaseProvider
}

func NewHandler(usecase usecaseProvider) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET("/", h.WelcomeMessage)
	router.GET("/_admin/ping", h.Ping)
	router.GET("/_admin/apps", h.HandleGetAllMasterApp)

	router.GET("/app", h.HandleGetMasterApp)
	router.POST("/app", h.HandleCreateNewMasterApp)
	router.PUT("/app", h.HandleUpdateMasterApp)
}

func WriteJson(ctx *gin.Context, data interface{}, err error, statusCode ...int) {
	payload := map[string]interface{}{
		"success": true,
	}
	code := http.StatusOK
	if data != nil {
		code = http.StatusOK
		payload["data"] = data
	}

	if err != nil {
		code = http.StatusInternalServerError
		payload["success"] = false
		payload["error"] = err.Error()
	}

	if len(statusCode) > 0 {
		code = statusCode[0]
		payload["code"] = code
	}

	ctx.JSON(code, payload)
}
