package main

import (
	"authentication-service/logger"
)

func main() {
	// Initialize configurations
	// config, err := configs.InitConfig(".")
	// if err != nil {
	// 	log.Fatalf("Error initializing config: %v", err)
	// }

	// Initialize logger
	logger.InitLogger()
	log := logger.GetLogger()

	// Example usage of logger
	log.Info("Starting the application...")
}
