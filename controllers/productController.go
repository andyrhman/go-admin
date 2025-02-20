package controllers

import (
	"go-admin/db"
	"go-admin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(db.DB, &models.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	db.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	product := models.Product{
		Id: uid,
	}

	db.DB.First(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	// * store the data parsed from the request body
	product := models.Product{
		Id: uid,
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	db.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	product := models.Product{
		Id: uid,
	}

	db.DB.Model(&product).Delete(product)

	return c.SendStatus(fiber.StatusNoContent)
}
