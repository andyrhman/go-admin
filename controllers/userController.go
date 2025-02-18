package controllers

import (
	"go-admin/db"
	"go-admin/models"
	"go-admin/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	db.DB.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	hashedPassword := utils.HashPassword("123123")


	user.Password = []byte(hashedPassword)

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

	db.DB.Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id); 
	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}
	user := models.User{
		Id: uid,
	}

	// ! Always use BodyParser
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	db.DB.Model(&user).Updates(user)

	return c.JSON(user)
}