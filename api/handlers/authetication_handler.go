package handlers

import (
	rediss "authentication-service/api/redis"
	"authentication-service/genproto/authentication_service"
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/services"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AuthenticationHandler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Register(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
	VerifyEmailCode(ctx *gin.Context)
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

	smtpServer := "smtp.gmail.com"
	authEmail := "abdusamatovjavohir@gmail.com"
	authPassword := "xsay zgvy uuvd xven"
	smtpPort := "587"

	code := uuid.NewString()
	// Receiver email address.
	to := []string{reReq.Email}
	rediss.SaveVerificationCode(reReq.Email, code, time.Minute*5)
	// Message.
	subject := "Subject: Verification Code to register Artisant Connect\n"
	body := "Your verification code: " + code
	message := []byte(subject + "\n" + body)

	// Authentication.
	auth := smtp.PlainAuth("", authEmail, authPassword, smtpServer)
	log.Println(reReq.Email)
	// Sending email.
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, authEmail, to, message)
	if err != nil {
		h.log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
	}

	reRes, err := h.authService.Register(ctx, &reReq)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Verification sent successfully", "user": reRes})
}

// @Description Verify Registration code
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param code query string true "verification code"
// @Param email query string true "email"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/auth/verify_email [post]
func (h *authenticationHandlerImpl) VerifyEmailCode(ctx *gin.Context) {
	code := ctx.Query("code")
	email := ctx.Query("email")

	if code == "" || email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing code or email"})
		return
	}

	verificationCode, err := rediss.GetVerificationCode(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get verification code from Redis"})
		return
	}

	if verificationCode != code {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verification code"})
		return
	}

	if verificationCode == code {
		if err := h.authService.VerifyEmailCodeAndUpdateUserInfo(ctx, email); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user status"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "User activated successfully"})
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

// @Summary Change password
// @Description Change password
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param ChangePassword body authentication_service.ChangePasswordRequest true "Change password"
// @security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/auth/change-password [post]
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

// @Summary Reset password
// @Description Reset password via email
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param ResetPassword body authentication_service.ResetPasswordRequest true "Reset password"
// @security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/auth/reset-password [post]
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
