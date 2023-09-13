package router

import (
	"logger/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/health", handlers.HealthCheck)

	v1.Post("/log", handlers.PostLog)
	v1.Get("/logs", handlers.GetLogs)
}