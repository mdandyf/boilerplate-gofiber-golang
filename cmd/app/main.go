package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdandyf/go-boilerplate/internal/app"
	"github.com/mdandyf/go-boilerplate/pkg/jwt"
	"github.com/mdandyf/go-boilerplate/pkg/logger"
	"github.com/mdandyf/go-boilerplate/pkg/redis"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Initialize logger
	logger := logger.NewLogger()

	// Initialize Redis client
	redisClient := redis.NewClient()

	// Initialize JWT manager
	jwtManager := jwt.NewManager()

	// Initialize the application
	application := app.NewApp(app, logger, redisClient, jwtManager)

	// Setup routes
	application.SetupRoutes()

	// Start the server
	err := application.StartServer()
	if err != nil {
		logger.Fatal("Failed to start server:", err)
	}
}
