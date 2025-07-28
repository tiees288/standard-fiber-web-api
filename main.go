package main

import (
	"log"

	"standard-fiber-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()

	// Add your routes and middleware here
	app.Get("/dashboard", monitor.New())

	// Register all routes
	routes.MapRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
