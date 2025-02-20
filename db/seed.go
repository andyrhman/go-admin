package db

import (
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"go-admin/models"
)

// SeedFakeOrders creates 20 orders, each with 3 order items, using fake data.
func SeedFakeOrders() {
	// Seed the random generator (you can provide a specific seed for reproducibility)
	gofakeit.Seed(0)

	// Loop 20 times to create 20 orders.
	for i := 0; i < 20; i++ {
		order := models.Order{
			Name:      gofakeit.Name(),
			Email:     gofakeit.Email(),
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}

		// Add 3 fake order items to each order.
		for j := 0; j < 3; j++ {
			item := models.OrderItem{
				ProductTitle: gofakeit.ProductName(),
				Price:        float32(gofakeit.Price(1, 100)), // Price between 1 and 100
				Quantity:     uint(gofakeit.Number(1, 5)),       // Quantity between 1 and 5
			}
			order.OrderItems = append(order.OrderItems, item)
		}

		// Create the order and its associated order items in the database.
		if err := DB.Create(&order).Error; err != nil {
			log.Fatalf("Error seeding order: %v", err)
		}
	}

	log.Println("Successfully seeded 20 orders with 3 order items each!")
}
