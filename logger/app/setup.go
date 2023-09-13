package app

import (
	"os"
	"logger/database"
	"logger/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func Setup() error {
	err := godotenv.Load()
	if err!=nil {
		return err
	}
	database.Connect();

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))
	router.SetupRoutes(app)
	// gets the port and start
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	
	app.Listen(host + ":" + port)

	return nil
}