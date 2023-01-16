package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	routes "cool-api/routes"
)

func registerRoute(db *gorm.DB) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.RegisterUserRoutes(app, db)

	app.Listen(":3000")
}
