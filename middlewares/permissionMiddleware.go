package middlewares

import (
	"errors"
	"go-admin/db"
	"go-admin/models"
	"go-admin/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("user_session")

	Id, err := utils.ParseJwt(cookie);

	uid, _ := uuid.Parse(Id)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	user := models.User{
		Id: uid,
	}

	db.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	db.DB.Preload("Permissions").Find(&user)

	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")
}