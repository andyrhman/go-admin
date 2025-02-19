package controllers

import (
	"go-admin/db"
	"go-admin/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	db.DB.Preload("Role").Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user.SetPassword("123123")

	db.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	user := models.User{
		Id: uid,
	}

	db.DB.Preload("Role").First(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	// * store the data parsed from the request body
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	// * store the existing user data fetched from the database
	var existingUser models.User
	if err := db.DB.Where("id = ?", uid).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	if input.FullName != "" {
		existingUser.FullName = input.FullName
	}

	if input.Email != "" && input.Email != existingUser.Email {
		var existingUserByEmail models.User
		if err := db.DB.Where("email = ?", input.Email).First(&existingUserByEmail).Error; err == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Email already exists",
			})
		}
		existingUser.Email = input.Email
	}

	if input.Username != "" && input.Username != existingUser.Username {
		var existingUserByUsername models.User
		if err := db.DB.Where("username = ?", input.Username).First(&existingUserByUsername).Error; err == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Username already exists",
			})
		}
		existingUser.Username = input.Username
	}

	if input.RoleId != 0 {
        var role models.Role
        if err := db.DB.Where("id = ?", input.RoleId).First(&role).Error; err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "message": "Role not found",
            })
        }
        existingUser.RoleId = input.RoleId
    }

	db.DB.Save(&existingUser)

	return c.JSON(existingUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	db.DB.Where("id = ?", uid).Delete(&models.User{})

	return c.SendStatus(fiber.StatusNoContent)
}

func TestUpdate(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	db.DB.Where("id = ?", uid).Updates(&input)

	return c.JSON(input)
}
