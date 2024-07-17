package handlers

import (
	"authentication-service/genproto/authentication_service"
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthenticationHandler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Register(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}

type authenticationHandlerImpl struct {
	log         *logrus.Logger
	authService services.AuthenticationService
}

func NewAuthenticationHandler(authService services.AuthenticationService, log *logrus.Logger) AuthenticationHandler {
	return &authenticationHandlerImpl{authService: authService, log: log}
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
		h.log.Error(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	reRes, err := h.authService.Register(ctx, &reReq)
	if err != nil {
		h.log.Error(err)
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
	var req auth.LogoutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.authService.Logout(ctx.Request.Context(), &req)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (h *authenticationHandlerImpl) ChangePassword(ctx *gin.Context) {
	var req auth.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.authService.ChangePassword(ctx.Request.Context(), &req)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

func (h *authenticationHandlerImpl) ResetPassword(ctx *gin.Context) {
	var req auth.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.authService.ResetPassword(ctx.Request.Context(), &req)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset email sent successfully"})
}

// func (h *authenticationHandlerImpl) VerifyResetCode(ctx *gin.Context) {
// 	var req auth.VerifyResetCodeRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	_, err := h.authService.VerifyResetCode(ctx.Request.Context(), &req)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

//		ctx.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
//	}
//
