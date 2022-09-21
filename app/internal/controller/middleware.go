package controller

import (
	"github.com/Vityalimbaev/Example-Backend/config"
	"github.com/Vityalimbaev/Example-Backend/internal/constants"
	"github.com/Vityalimbaev/Example-Backend/internal/entity"
	"github.com/Vityalimbaev/Example-Backend/internal/service/user"
	"github.com/Vityalimbaev/Example-Backend/pkg/exception"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(ctx *gin.Context) {

	errView := entity.ResponseView{
		IsSuccess: false,
		Data:      exception.UnauthorizedError.Error(),
	}

	authHeader := ctx.GetHeader("Authorization")

	authHeaderParts := strings.Split(authHeader, "Bearer")

	if len(authHeader) == 0 || len(authHeaderParts) < 2 {
		ctx.AbortWithStatusJSON(http.StatusOK, errView)
		return
	}

	accessToken := strings.TrimSpace(authHeaderParts[1])

	tokenKey := config.GetServerConfig().TokenKey

	sub, isValid, isExpired := user.ParseToken(accessToken, tokenKey)

	if !isValid || isExpired {
		ctx.AbortWithStatusJSON(http.StatusOK, errView)
		return
	}

	ctx.Set(constants.CtxUserIdKey, sub[constants.CtxUserIdKey])
	ctx.Set(constants.CtxActiveStatusKey, sub[constants.CtxActiveStatusKey])
	ctx.Set(constants.CtxRoleIdKey, sub[constants.CtxRoleIdKey])

	ctx.Next()
}
