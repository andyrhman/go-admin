package controllers

import (
	"go-admin/db"
	"go-admin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	db.DB.Find(&roles)

	return c.JSON(roles)
}

type CreateRoleRequest struct {
	Name        string `json:"name" validate:"required"`
	Permissions []uint `json:"permissions" validate:"required,dive,gt=0"`
}

func CreateRole(c *fiber.Ctx) error {
	var req CreateRoleRequest

	// Parse the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Create a Role instance with the provided name.
	role := models.Role{
		Name: req.Name,
	}

	// Map the permission IDs to Permission structs.
	for _, permID := range req.Permissions {
		role.Permissions = append(role.Permissions, models.Permission{
			Id: permID,
		})
	}

	// Save the new role to the database using GORM.
	if err := db.DB.Create(&role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create role",
		})
	}

	// Return the created role with status 201.
	return c.Status(fiber.StatusCreated).JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var roles models.Role
	
	if err := db.DB.Where("id = ?", id).First(&roles).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error not found",
		})
	}

	return c.JSON(roles)
}

type UpdateRoleRequest struct {
	Name        string `json:"name" validate:"required"`
	Permissions []uint `json:"permissions" validate:"required,dive,gt=0"`
}

func UpdateRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var req UpdateRoleRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	// Fetch the existing role.
	var role models.Role
	if err := db.DB.Preload("Permissions").First(&role, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Update the role's basic field.
	role.Name = req.Name
	if err := db.DB.Model(&role).Update("name", req.Name).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update role name"})
	}

	var newPermissions []models.Permission
	for _, permID := range req.Permissions {
		newPermissions = append(newPermissions, models.Permission{Id: permID})
	}

	if err := db.DB.Model(&role).Association("Permissions").Replace(newPermissions); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update permissions"})
	}

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	db.DB.Where("id = ?", id).Delete(&models.Role{})

	return c.SendStatus(fiber.StatusNoContent)
}