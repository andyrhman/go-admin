package models

import "github.com/google/uuid"

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
