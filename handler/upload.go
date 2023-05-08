package handler

import (
	"errors"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) HandleUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		WriteJson(ctx, nil, err)
		return
	}

	if file.Size >= 1*1024*1024 {
		WriteJson(ctx, nil, errors.New("file size is too large"), http.StatusBadRequest)
		return
	}

	filenameSegments := strings.Split(file.Filename, ".")
	if len(filenameSegments) < 2 {
		WriteJson(ctx, nil, errors.New("file extension is not supported"), http.StatusBadRequest)
		return
	}

	file.Filename = strings.Replace(uuid.NewString(), "-", "", -1)[:10] + "." + filenameSegments[len(filenameSegments)-1]
	dst := filepath.Join("./assets/" + file.Filename)
	ctx.SaveUploadedFile(file, dst)

	WriteJson(ctx, file, nil)
}
