package handler

import (
	"dynapgen/usecase"
	"dynapgen/util/log"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrz1836/go-sanitize"
)

var (
	errorTemplateIDEmpty = errors.New("template id is empty")
)

func (h *Handler) HandleGetAllTemplates(ctx *gin.Context) {
	result, err := h.usecase.GetAllTemplates(ctx)
	if err != nil {
		log.Error(err, "h.usecase.GetAllTemplates() got error - HandleGetAllTemplates")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, result, nil, http.StatusOK)
}

func (h *Handler) HandleGetTemplateByID(ctx *gin.Context) {
	templateID := sanitize.AlphaNumeric(ctx.Query("id"), false)
	if templateID == "" {
		WriteJson(ctx, nil, errorTemplateIDEmpty, http.StatusBadRequest)
		return
	}

	result, err := h.usecase.GetTemplateByID(ctx, templateID)
	if err != nil {
		log.Error(err, "h.usecase.GetTemplateByID() got error - HandleGetTemplateByID")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, result, nil, http.StatusOK)
}

func (h *Handler) HandleCreateTemplate(ctx *gin.Context) {
	var template Template
	if err := ctx.ShouldBindJSON(&template); err != nil {
		log.Error(err, "ctx.ShouldBindJSON() got error - HandleCreateTemplate", nil)
		WriteJson(ctx, nil, err)
		return
	}

	if err := h.usecase.AddTemplate(ctx, usecase.Template(template)); err != nil {
		log.Error(err, "h.usecase.GetTemplateByID() got error - HandleGetTemplateByID")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil, http.StatusAccepted)
}

func (h *Handler) HandleUpdateTemplate(ctx *gin.Context) {
	var template Template
	if err := ctx.ShouldBindJSON(&template); err != nil {
		log.Error(err, "ctx.ShouldBindJSON() got error - HandleUpdateTemplate", nil)
		WriteJson(ctx, nil, err)
		return
	}

	if err := h.usecase.UpdateTemplate(ctx, usecase.Template(template)); err != nil {
		log.Error(err, "h.usecase.GetTemplateByID() got error - HandleUpdateTemplate")
		WriteJson(ctx, nil, err)
		return
	}

	WriteJson(ctx, nil, nil, http.StatusAccepted)
}
