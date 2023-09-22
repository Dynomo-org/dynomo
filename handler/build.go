package handler

import (
	"dynapgen/usecase"
	"dynapgen/util/assets"
	"errors"
	"fmt"
	"net/http"
	"strings"

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

	file, err := ctx.FormFile("keystore")
	if err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	if file.Size >= 3*1024*1024 {
		WriteJson(ctx, nil, errors.New("file size is too large"), http.StatusBadRequest)
		return
	}

	filenameSegments := strings.Split(file.Filename, ".")
	if len(filenameSegments) < 2 {
		WriteJson(ctx, nil, errors.New("file extension not found"), http.StatusBadRequest)
		return
	}

	fileExt := filenameSegments[len(filenameSegments)-1]
	if _, ok := keystoreFileTypes[strings.ToLower(fileExt)]; !ok {
		WriteJson(ctx, nil, errors.New("file extension not supported"), http.StatusBadRequest)
		return
	}

	// TODO: uniform keystore naming
	// upload to github with same folder
	dst := assets.GenerateWorkFilePath(fmt.Sprintf("%s.%s", file.Filename, fileExt))
	ctx.SaveUploadedFile(file, dst)

	err = h.usecase.BuildApp(ctx, usecase.BuildAppParam{
		AppID:        appID,
		KeystorePath: dst,
	})
	if err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil, http.StatusAccepted)
}
