package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
}

func (product *Product) Count(db *gorm.DB) int64  {
	var total int64
	db.Model(&Product{}).Count(&total)

	return total
}

func (product *Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []Product

	db.Offset(offset).Limit(limit).Find(&products)

	return products
}