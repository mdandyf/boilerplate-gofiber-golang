package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdandyf/go-boilerplate/pkg/jwt"
	"github.com/mdandyf/go-boilerplate/pkg/logger"
	"github.com/mdandyf/go-boilerplate/pkg/redis"
)

// App represents the application.
type App struct {
	Router *fiber.App
	Redis  *redis.Client
	JWT    *jwt.Manager
}

// NewApp creates a new instance of the application.
func NewApp() *App {
	app := &App{
		Router: fiber.New(),
		Redis:  redis.NewClient(),
		JWT:    jwt.NewManager(),
	}

	// Initialize routes
	app.initRoutes()

	return app
}

// initRoutes initializes the application routes.
func (a *App) initRoutes() {
	// Define your routes here
}

// Start starts the application.
func (a *App) Start() error {
	logger.Info("Starting the application...")
	return a.Router.Listen(":3000")
}
