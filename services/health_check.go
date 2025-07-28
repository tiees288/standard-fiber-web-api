package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	// Return a simple JSON response
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Service is up and running",
	})
}