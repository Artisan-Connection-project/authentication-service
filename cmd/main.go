package main

import (
	"authentication-service/api"
	"authentication-service/api/handlers"
	"authentication-service/configs"
	"authentication-service/logger"
	"authentication-service/services"
	"authentication-service/storage/postgres"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.InitLogger()

	log := logger.GetLogger()
	log.WithFields(logrus.Fields{
		"TestLogger": "test-logger",
	})

	config, err := configs.InitConfig(".")
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	db, err := postgres.ConnectDB(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepository(db)
	hasher := postgres.NewBcryptHasher(10)
	authRepo := postgres.NewAuthenticationRepository(userRepo, hasher, db)
	tokenRepo := postgres.NewTokenRepository(db)

	tokenService := services.NewTokenService(tokenRepo, config.SecretKey)
	userService := services.NewUserManagementService(userRepo)
	authService := services.NewAuthenticationService(userRepo, authRepo, tokenService)

	mainHandler := handlers.NewMainHandler(authService, userService, tokenService)

	router := gin.Default()

	api.SetupRouters(router, mainHandler, config)

	err = router.Run(":" + config.GinServerPort)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
