package middlewares

import (
	"go-admin/utils"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("user_session")

	if _, err := utils.ParseJwt(cookie); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	return c.Next()
}
