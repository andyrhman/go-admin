package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
	limit := 5
	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)
	
	total := entity.Count(db)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": (total + int64(limit) - 1) / int64(limit),
		},
	}
}
