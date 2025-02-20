package main

import (
	"go-admin/db"
	"go-admin/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.Connect()

	// db.SeedFakeOrders() * call the seed function

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000", // specify the allowed origin(s)
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
