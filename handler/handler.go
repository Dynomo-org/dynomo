package handler

import (
	"context"
	"dynapgen/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usecaseProvider interface {
	DeleteApp(ctx context.Context, appID string) error
	GetAllApp(ctx context.Context) ([]usecase.App, error)
	GetApp(ctx context.Context, appID string) (usecase.App, error)
	GetGenerateKeystoreStatus(ctx context.Context, appID string) (usecase.Keystore, error)
	GenerateKeystore(ctx context.Context, param usecase.GenerateStoreParam) error
	NewApp(ctx context.Context, request usecase.NewAppRequest) error
	SaveApp(ctx context.Context, App usecase.App) error
	UpdateApp(ctx context.Context, App usecase.App) error
	UpdateAppIcon(ctx context.Context, appID string, iconName, localPath string) error
}

type Handler struct {
	usecase usecaseProvider
}

func NewHandler(usecase usecaseProvider) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20

	router.GET("/", h.WelcomeMessage)
	router.GET("/_admin/ping", h.Ping)
	router.GET("/_admin/apps", h.HandleGetAllApp)
	router.GET("/app", h.HandleGetApp)
	router.GET("/keystore", h.HandleGetGenerateKeystoreStatus)

	router.POST("/app", h.HandleCreateNewApp)
	router.POST("/keystore", h.HandleGenerateKeystore)

	router.PUT("/app", h.HandleUpdateApp)
	router.PUT("/app/icon", h.HandleUpdateAppIcon)

	router.DELETE("/app", h.HandleDeleteApp)
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
	}
	payload["code"] = code

	ctx.JSON(code, payload)
}
