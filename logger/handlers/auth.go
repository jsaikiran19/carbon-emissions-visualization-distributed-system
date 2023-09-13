package handlers

import (
	"time"

	"net/http"
	"logger/models"
	"logger/database"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func PostLog(c *fiber.Ctx) error {
	var data map[string]string
	
	if err := c.BodyParser(&data); err != nil {
		return err
	} 
	message := data["message"]
	username := data["username"]
	if message == "" {
		response := models.BuildResponse(http.StatusText(401), "Log cannot be empty", models.APIResponse{}, "Error")
		return c.Status(401).JSON(response)
	}
	db := database.DB
	logInfo := models.Log{
		Message: message,
		Username: username,
		Timestamp: time.Now(),
	}

	q := db.Create(&logInfo);
	if q.Error != nil {
		statusCode := fiber.StatusInternalServerError
		response := models.BuildResponse(http.StatusText(statusCode), "Failed to create new user", nil, q.Error.Error())
		return c.Status(statusCode).JSON(response)
	}
	response := models.BuildResponse(http.StatusText(fiber.StatusCreated), "Log posted successfully", logInfo, "")
	return c.Status(fiber.StatusCreated).JSON(response)
}

func GetLogs(c *fiber.Ctx) error {
	var logs []models.Log
	db := database.DB
	query := db.Find(&logs)

	if query.Error != nil {
		statusCode := fiber.StatusInternalServerError
		response := models.BuildResponse(http.StatusText(statusCode), "Failed to get logs", nil, query.Error.Error())
		return c.Status(statusCode).JSON(response)
	}
	return c.Status(200).JSON(logs)  // return json response with status code of OK (200), and the list of log entries as

}