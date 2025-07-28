// Go fiber create route group
package routes

import (
	"standard-fiber-go/shared/enum"
	"standard-fiber-go/shared/models"
	registerModels "standard-fiber-go/shared/models/register"

	"github.com/gofiber/fiber/v2"
)

// RegisterUserResp defines the response structure for user registration

// RegisterRoutes function
func RegisterRoutes(route fiber.Router) {
	register := route.Group("/v1/register")

	// Define your routes here
	register.Post("/user", func(c *fiber.Ctx) error {
		// Handle user registration logic
		resp := models.BaseResp[*registerModels.RegisterUserResp]{
			Message: string(enum.Success),
			Data:    nil, // Data is null initially
		}

		// // Registration logic here...
		// // After successful registration, set the user data
		// resp.Data = &RegisterUserResp{
		// 	UserID:   "12345",
		// 	Username: "exampleuser",
		// }

		return c.JSON(resp)
	})
}
