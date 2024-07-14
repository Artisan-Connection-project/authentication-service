package handlers

import (
	"authentication-service/services"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Register(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}

type authenticationHandlerImpl struct {
	authService services.AuthenticationService
}

func NewAuthenticationHandler(authService services.AuthenticationService) AuthenticationHandler {
	return &authenticationHandlerImpl{authService: authService}
}

func (h *authenticationHandlerImpl) Login(ctx *gin.Context) {

	h.authService.Login(ctx)
}

func (h *authenticationHandlerImpl) Logout(ctx *gin.Context) {
	// Implement logout logic
}

func (h *authenticationHandlerImpl) Register(ctx *gin.Context) {

}

func (h *authenticationHandlerImpl) ResetPassword(ctx *gin.Context) {
	// Implement password reset logic
}

func (h *authenticationHandlerImpl) ChangePassword(ctx *gin.Context) {
	// Implement password change logic
}
