package handlers

import (
	"authentication-service/genproto/authentication_service"
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

// @Summary Register the authentication
// @Description Creates new user with the given username and password and other information
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param RegisterRequest body authentication_service.RegisterRequest true "Register the new user with the given username and password"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/auth/register [post]
func (h *authenticationHandlerImpl) Register(ctx *gin.Context) {
	var reReq authentication_service.RegisterRequest
	if err := ctx.ShouldBindJSON(&reReq); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	reRes, err := h.authService.Register(ctx, &reReq)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, reRes)
}

// @Summary Singing in
// @Description Login via password and username or email
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param RegisterRequest body authentication_service.LoginRequest true "Login username and  password"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/auth/login [post]
func (h *authenticationHandlerImpl) Login(ctx *gin.Context) {
	var lrReq authentication_service.LoginRequest
	if err := ctx.ShouldBindJSON(&lrReq); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	lrRes, err := h.authService.Login(ctx, &lrReq)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, lrRes)

}

func (h *authenticationHandlerImpl) Logout(ctx *gin.Context) {
	// Implement logout logic
}

func (h *authenticationHandlerImpl) ResetPassword(ctx *gin.Context) {
	// Implement password reset logic
}

func (h *authenticationHandlerImpl) ChangePassword(ctx *gin.Context) {
	// Implement password change logic
}
