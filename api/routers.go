package api

import (
	"authentication-service/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.Engine, h handlers.MainHandler) {

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.Authentication().Login)
			auth.POST("/logout", h.Authentication().Logout)
			auth.POST("/register", h.Authentication().Register)
			auth.POST("/reset-password", h.Authentication().ResetPassword)
			auth.POST("/change-password", h.Authentication().ChangePassword)
		}
		users := api.Group("/users")
		{
			users.GET("/", h.User().GetUsers)
			users.GET("/:username_or_email", h.User().GetUserByUsernameOrEmail)
			users.GET("/profile/:id", h.User().GetUserInfo)
			users.PUT("/profile/:id", h.User().UpdateUserInfo)
			users.PUT("/type/:id", h.User().ChangeUserType)
			users.DELETE("/profile/:id", h.User().DeleteUser)
		}
		tokens := api.Group("/tokens")
		{
			tokens.POST("/refresh/:user_id", h.Token().RefreshToken)
			tokens.GET("/revoke/:user_id", h.Token().CancelToken)
		}
	}
}
