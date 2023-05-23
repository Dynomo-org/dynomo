package handler

import (
	"dynapgen/utils/tokenizer"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	errorNoAccess = errors.New("no access")
)

func checkUserAuthorization(ctx *gin.Context) {
	authHeader := ctx.Request.Header["Authorization"]
	if len(authHeader) == 0 {
		WriteJson(ctx, nil, errorNoAccess, http.StatusForbidden)
		ctx.Abort()
		return
	}

	auth := authHeader[0]
	segments := strings.Split(auth, " ")
	if len(segments) < 2 {
		WriteJson(ctx, nil, errorNoAccess, http.StatusForbidden)
		ctx.Abort()
		return
	}

	if strings.ToLower(segments[0]) != "bearer" {
		WriteJson(ctx, nil, errorNoAccess, http.StatusForbidden)
		ctx.Abort()
		return
	}

	token := segments[1]
	claims, err := tokenizer.VerifyAndParseJWTToken(token)
	if err != nil {
		WriteJson(ctx, nil, err, http.StatusInternalServerError)
		ctx.Abort()
		return
	}

	ctx.Set("user_id", claims["id"])
	ctx.Next()
}
