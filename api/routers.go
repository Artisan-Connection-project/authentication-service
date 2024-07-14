package api

import (
	_ "authentication-service/api/docs"
	"authentication-service/api/handlers"
	"authentication-service/api/middleware"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Testing auth API
// @version         1.0
// @description     This is a sample server.
// @host            localhost:8081
// @BasePath        /
// @schemes         http
// @securityDefinitions.apiKey BearerAuth
// @in              header
// @name            Authorization
func SetupRouters(r *gin.Engine, h handlers.MainHandler) {
	r.GET("swagger/*any", ginSwagger.WrapHandler(files.Handler))

	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.Authentication().Login)
			auth.POST("/register", h.Authentication().Register)
		}
		r.Use(middleware.AuthMiddleware())

		auth.POST("/logout", h.Authentication().Logout)
		auth.POST("/reset-password", h.Authentication().ResetPassword)
		auth.POST("/change-password", h.Authentication().ChangePassword)

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
