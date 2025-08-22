package healthcheck

import (
	"standard-fiber-go/shared/database"

	"github.com/gofiber/fiber/v2"
)

// DatabaseHealthCheck returns database health status
func DatabaseHealthCheck(c *fiber.Ctx) error {
	healthStatus := database.HealthCheck()
	
	if !healthStatus["database_connected"].(bool) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status":  "error",
			"message": "Database connection failed",
			"details": healthStatus,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Database is healthy",
		"details": healthStatus,
	})
}