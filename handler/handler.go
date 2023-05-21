package handler

import (
	"dynapgen/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20

	router.GET("/", h.WelcomeMessage)
	router.GET("/_admin/ping", h.Ping)
	router.GET("/_admin/apps", h.HandleGetAllApps)
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
