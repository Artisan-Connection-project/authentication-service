package handlers

import (
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenHandler interface {
	RefreshToken(ctx *gin.Context)
	// CancelToken(ctx *gin.Context)
}

type tokenHandlerImpl struct {
	tokenService services.TokenService
}

func NewTokenHandler(tokenSer services.TokenService) TokenHandler {
	return &tokenHandlerImpl{tokenService: tokenSer}
}

func (h *tokenHandlerImpl) GenerateToken(c *gin.Context) {
	var req auth.GenerateTokenRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.tokenService.GenerateToken(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Refresh a token
// @Description Refreshes an existing access token using a refresh token
// @Tags Tokens
// @Accept  json
// @Produce  json
// @Param refreshToken body authentication_service.RefreshTokenRequest true "Refresh Token Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/tokens/refresh-token/{user_id} [post]
func (h *tokenHandlerImpl) RefreshToken(c *gin.Context) {
	userID := c.Param("user_id")

	var req auth.RefreshTokenRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.tokenService.RefreshToken(c, userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// // @Summary Cancel a token
// // @Description Cancels a given access or refresh token
// // @Tags Tokens
// // @Accept  json
// // @Produce  json
// // @Param cancelToken body authentication_service.CancelTokenRequest true "Cancel Token Request"
// // @Success 200 {object} map[string]interface{}
// // @Failure 400 {object} map[string]string
// // @Failure 401 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /revoke/{user_id} [post]
// func (h *tokenHandlerImpl) CancelToken(c *gin.Context) {
// 	userID := c.Param("user_id")

// 	var req auth.CancelTokenRequest
// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	resp, err := h.tokenService.CancelToken(c, userID, &req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel token"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp)
// }
