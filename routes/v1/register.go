// Go fiber create route group
package routes

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes function
func RegisterRoutes(route fiber.Router) {
	register := route.Group("/v1/register")

	// Define your routes here
	register.Post("/user", func(c *fiber.Ctx) error {
		// Handle user registration logic
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User registered successfully",
		})
	})
}
