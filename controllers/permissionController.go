package controllers

import (
	"go-admin/db"
	"go-admin/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	db.DB.Find(&permissions)

	return c.JSON(permissions)
}

func CreatePermission(c *fiber.Ctx) error {
	var req models.Permission

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	db.DB.Save(&req)

	return c.JSON(req)
}