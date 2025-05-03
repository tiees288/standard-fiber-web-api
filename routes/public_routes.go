// Go fiber create route group
package routes

import (
	healthcheck "webinfo-go/services"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes function
func PublicRoutes(app *fiber.App) {
	// Create a new route group
	public := app.Group("/api/v1/public")

	// Define the routes for the public group
	public.Get("/health", healthcheck.HealthCheck)
}
