// Go fiber create route group
package routes

import (
	healthcheck "webinfo-go/services"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes function
func PublicRoutes(route fiber.Router) {
	// Create a new route group
	public := route.Group("/v1")

	// Define the routes for the public group
	public.Get("/health", healthcheck.HealthCheck)
}
