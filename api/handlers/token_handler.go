package handlers

import (
	"authentication-service/services"

	"github.com/gin-gonic/gin"
)

type TokenHandler interface {
	RefreshToken(ctx *gin.Context)
	VerifyToken(ctx *gin.Context)
	CancelToken(ctx *gin.Context)
}

type tokenHandlerImpl struct {
	tokenService services.TokenService
}

func NewTokenHandler(tokenSer services.TokenService) TokenHandler {
	return &tokenHandlerImpl{
		tokenService: tokenSer,
	}
}

func (h *tokenHandlerImpl) RefreshToken(ctx *gin.Context) {
	// Implement token refresh logic
}

func (h *tokenHandlerImpl) VerifyToken(ctx *gin.Context) {
	// Implement token verification logic
}

func (h *tokenHandlerImpl) CancelToken(ctx *gin.Context) {
	// Implement token cancellation logic
}
