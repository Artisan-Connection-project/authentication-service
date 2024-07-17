package api

import (
	_ "authentication-service/api/docs"
	"authentication-service/api/handlers"
	"authentication-service/configs"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Testing auth API
// @version         1.0
// @description     This is a sample server.
// @host            localhost:8081
// @BasePath        /
// @schemes         http
// @securityDefinitions.apiKey BearerAuth
// @in              header
// @name            Authorization
func SetupRouters(r *gin.Engine, h handlers.MainHandler, config *configs.Config) {
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

			auth.POST("/logout", h.Authentication().Logout)

			auth.POST("/reset-password", h.Authentication().ResetPassword)

			auth.POST("/change-password", h.Authentication().ChangePassword)
		}

		users := api.Group("/users")
		{
			users.GET("/all", h.User().GetUsers)

			users.GET("/profile/:user_id", h.User().GetUserInfo)

			users.PUT("/profile/", h.User().UpdateUserInfo)
			users.PUT("/type/", h.User().ChangeUserType)
			users.DELETE("/profile/:id", h.User().DeleteUser)
		}

		tokens := api.Group("/tokens")
		{
			tokens.POST("/refresh-token/:user_id", h.Token().RefreshToken)
			// tokens.GET("/revoke/:user_id", h.Token().CancelToken)
		}
	}
}
