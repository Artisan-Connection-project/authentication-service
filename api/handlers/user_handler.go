package handlers

import (
	auth "authentication-service/genproto/authentication_service"
	"authentication-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler interface {
	GetUserInfo(ctx *gin.Context)
	UpdateUserInfo(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	ChangeUserType(ctx *gin.Context)
}

type userHandlerImpl struct {
	userService services.UserManagementService
	log         *logrus.Logger
}

func NewUserHandler(userService services.UserManagementService, log *logrus.Logger) UserHandler {
	return &userHandlerImpl{
		userService: userService,
		log:         log,
	}
}

// @Summary Getting user by its id
// @Description Get user by its id
// @Tags User Management
// @Accept  json
// @Produce  json
// @Param user_id path string true "user id"
// @security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/profile/{user_id} [get]
func (h *userHandlerImpl) GetUserInfo(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	req := &auth.GetUserInfoRequest{Id: userId}
	res, err := h.userService.GetUserInfo(ctx, req)
	if err != nil {
		if err.Error() == "user does not exist" {
			h.log.Error("user does not exist:" + err.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user does not exist"})
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": res.User})
}

// @Summary Update user by the updated fields
// @Description Update user by the provided fields
// @Tags User Management
// @Accept  json
// @Produce  json
// @Param updatedUser body authentication_service.UpdateUserInfoRequest true "update user information"
// @security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/profile/ [put]
func (h *userHandlerImpl) UpdateUserInfo(ctx *gin.Context) {
	req := &auth.UpdateUserInfoRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.userService.UpdateUserInfo(ctx, req)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": res})
}

// @Summary Delete user info by the provided id
// @Description Delete user info by the provided id
// @Tags User Management
// @Accept  json
// @Produce  json
// @Param id path string true "user id"
// @security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/profile/{id} [delete]
func (h *userHandlerImpl) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	req := &auth.DeleteUserInfoRequest{Id: userId}
	_, err := h.userService.DeleteUserInfo(ctx, req)
	if err != nil {
		if err.Error() == "user does not exist" {
			h.log.Error(err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user does not exist"})
		}
		h.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// @Summary Get all users with page and sorting
// @Description Retrieve a list of users with pagination and sorting
// @Tags User Management
// @Accept  json
// @Produce  json
// @Param limit query int true "Number of users per page"
// @Param page query int true "Page number"
// @Param order_by query string false "Field to order by (e.g., 'name', 'email')"
// @security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/all [get]
func (h *userHandlerImpl) GetUsers(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")      // Default limit if not provided
	page := ctx.DefaultQuery("page", "1")         // Default page if not provided
	orderBy := ctx.DefaultQuery("order_by", "id") // Default ordering field if not provided

	req := &auth.GetUsersInfoRequest{
		Limit:   limit,
		Page:    page,
		OrderBy: orderBy,
	}

	res, err := h.userService.GetUsersInfo(ctx, req)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": res.Users})
}

// @Summary Change user type
// @Description Change user's type to admin, user, artisan or other
// @Tags User Management
// @Accept  json
// @Produce  json
// @Param user_id query string true "user id"
// @Param user_type query string true "user type"
// @security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/type/ [put]
func (h *userHandlerImpl) ChangeUserType(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	userType := ctx.Query("user_type")

	if userId == "" || userType == "" {
		h.log.Error("user_id and user_type cannot be empty")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id and user_type are required"})
		return
	}

	getReq := auth.GetUserInfoRequest{Id: userId}
	res, err := h.userService.GetUserInfo(ctx, &getReq)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	upReq := auth.UpdateUserInfoRequest{
		Id:       res.User.Id,
		Username: res.User.Username,
		FullName: res.User.FullName,
		Email:    res.User.Email,
		UserType: userType,
		Bio:      res.User.Bio,
	}

	upRes, err := h.userService.UpdateUserInfo(ctx, &upReq)

	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": upRes})
}
