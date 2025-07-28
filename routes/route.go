package routes

import (
	v1 "webinfo-go/routes/v1"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes function to register public routes
func MapRoutes(app fiber.Router) {
	api := app.Group("/api")
	// Map V1 routes
	v1.PublicRoutes(api)
	v1.RegisterRoutes(api)
}
