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

	router.GET("/_admin/ping", checkUserAuthorization, h.Ping)

	router.GET("/apps", checkUserAuthorization, h.HandleGetAllApps)
	router.GET("/app", checkUserAuthorization, h.HandleGetApp)
	router.GET("/app/full", h.HandleGetFullApp)
	router.POST("/app", checkUserAuthorization, h.HandleCreateNewApp)
	router.POST("/app/build", checkUserAuthorization, h.HandleBuildApp)
	router.PUT("/app", checkUserAuthorization, h.HandleUpdateApp)
	router.DELETE("/app", checkUserAuthorization, h.HandleDeleteApp)

	router.PUT("/app/icon", checkUserAuthorization, h.HandleUpdateAppIcon)

	router.GET("/ads", checkUserAuthorization, h.HandleGetAppAds)
	router.POST("/ads", checkUserAuthorization, h.HandleCreateNewAds)

	router.GET("/build-status/app", h.HandleGetBuildAppStatus)
	router.GET("/build-status/keystore", h.HandleGetBuildKeystoreStatus)
	router.PUT("/build-status/app", h.HandleUpdateBuildAppStatus)
	router.PUT("/build-status/keystore", h.HandleUpdateBuildKeystoreStatus)

	router.GET("/keystores", checkUserAuthorization, h.HandleGetAllKeystores)
	router.POST("/keystore", h.HandleGenerateKeystore)

	router.GET("/user/info", checkUserAuthorization, h.HandleGetUserInfo)
	router.POST("/user/login", h.HandleLoginUser)
	router.POST("/user/register", h.HandleRegisterUser)
}

func WriteJson(ctx *gin.Context, data interface{}, err error, statusCode ...int) {
	payload := map[string]interface{}{
		"is_success": true,
	}
	code := http.StatusOK
	if data != nil {
		payload["data"] = data
	}

	if err != nil {
		code = http.StatusInternalServerError
		payload["is_success"] = false
		payload["error"] = err.Error()
	}

	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	payload["code"] = code

	ctx.JSON(code, payload)
}
