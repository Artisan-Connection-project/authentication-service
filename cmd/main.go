package main

import (
	"authentication-service/api"
	"authentication-service/api/handlers"
	"authentication-service/configs"
	"authentication-service/genproto/authentication_service"
	"authentication-service/logger"
	"authentication-service/services"
	"authentication-service/storage/postgres"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

	// redisClient := cache.NewRedisCache("localhost:6379", "", 0)

	userRepo := postgres.NewUserRepository(db)
	hasher := postgres.NewBcryptHasher(10)
	authRepo := postgres.NewAuthenticationRepository(userRepo, hasher, db)
	tokenRepo := postgres.NewTokenRepository(db)

	emailService := services.NewEmailService("abdusamatovjavohir@gmail.com", "", "", "")
	tokenService := services.NewTokenService(tokenRepo, config.SecretKey)
	userService := services.NewUserManagementService(userRepo)
	authService := services.NewAuthenticationService(authRepo, tokenService, emailService)

	mainHandler := handlers.NewMainHandler(authService, userService, tokenService, log)

	go func() {
		router := gin.Default()
		api.SetupRouters(router, mainHandler, config)

		if err := router.Run(":" + config.GinServerPort); err != nil {
			log.Fatalf("Failed to run Gin server: %v", err)
		}
	}()

	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen on port 50051: %v", err)
		}

		grpcServer := grpc.NewServer()
		mainService := services.NewMainService(tokenService, authService, userService)

		authentication_service.RegisterAuthenticationServiceServer(grpcServer, mainService.(authentication_service.AuthenticationServiceServer))

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
