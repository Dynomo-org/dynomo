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

	router.GET("/ads", checkUserAuthorization, h.HandleGetAppAds)
	router.POST("/ads", checkUserAuthorization, h.HandleCreateNewAds)

	router.GET("/app", checkUserAuthorization, h.HandleGetApp)
	router.GET("/app/full", h.HandleGetFullApp)
	router.POST("/app", checkUserAuthorization, h.HandleCreateNewApp)
	router.POST("/app/build", checkUserAuthorization, h.HandleBuildApp)
	router.PUT("/app", checkUserAuthorization, h.HandleUpdateApp)
	router.DELETE("/app", checkUserAuthorization, h.HandleDeleteApp)
	router.PUT("/app/icon", checkUserAuthorization, h.HandleUpdateAppIcon)
	router.GET("/apps", checkUserAuthorization, h.HandleGetAllApps)

	router.GET("/app/contents", checkUserAuthorization, h.HandleGetAllContents)
	router.GET("/app/content", checkUserAuthorization, h.HandleGetContentDetail)
	router.POST("/app/content", checkUserAuthorization, h.HandleCreateNewContent)
	router.PUT("/app/content", checkUserAuthorization, h.HandleUpdateContent)

	router.GET("/artifacts", checkUserAuthorization, h.HandleGetBuildArtifacts)

	router.GET("/build-status/app", h.HandleGetBuildAppStatus)
	router.GET("/build-status/keystore", h.HandleGetBuildKeystoreStatus)
	router.PUT("/build-status/app", h.HandleUpdateBuildAppStatus)
	router.PUT("/build-status/keystore", h.HandleUpdateBuildKeystoreStatus)

	router.GET("/keystores", checkUserAuthorization, h.HandleGetKeystores)
	router.POST("/keystore", checkUserAuthorization, h.HandleGenerateKeystore)

	router.GET("/user/info", checkUserAuthorization, h.HandleGetUserInfo)
	router.POST("/user/login", h.HandleLoginUser)
	router.POST("/user/register", h.HandleRegisterUser)

	router.GET("/templates", checkUserAuthorization, h.HandleGetAllTemplates)
	router.GET("/template", checkUserAuthorization, h.HandleGetTemplateByID)
	router.POST("/template", checkUserAuthorization, h.HandleCreateTemplate)
	router.PUT("/template", checkUserAuthorization, h.HandleUpdateTemplate)
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
