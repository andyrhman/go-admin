package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	Id         uuid.UUID   `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	UpdatedAt  string      `json:"updated_at"`
	CreatedAt  string      `json:"created_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrderId      uuid.UUID `json:"order_id"`
	ProductTitle string    `json:"product_title"`
	Price        float32   `json:"price"`
	Quantity     uint      `json:"quantity"`
}

func (order *Order) Count(db *gorm.DB) int64  {
	var total int64
	db.Model(&Order{}).Count(&total)

	return total
}

func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order

	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	return orders
}