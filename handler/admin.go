package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Ping(ctx *gin.Context) {
	WriteJson(ctx, map[string]interface{}{
		"pong":    "pong",
		"user_id": ctx.GetString("user_id"),
	}, nil)
}

func (h *Handler) WelcomeMessage(ctx *gin.Context) {
	WriteJson(ctx, map[string]interface{}{
		"message": "Hellow world!",
	}, nil)
}
