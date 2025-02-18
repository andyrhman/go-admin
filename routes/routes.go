package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/controllers"
)

func Setup(app *fiber.App) {
	// * Health Check
	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(200)
		return c.JSON(fiber.Map{
			"message": "Server status is ok 😁👍",
		})
	})

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
}
