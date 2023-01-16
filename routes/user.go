package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	s "cool-api/services"
	. "cool-api/types"
)

func RegisterUserRoutes(app *fiber.App, db *gorm.DB) {

	app.Post("/user", func(c *fiber.Ctx) error {

		bodyRequest := &UserCreateRequest{}

		if err := c.BodyParser(&bodyRequest); err != nil {
			return err
		}

		user, err := s.CreateUser(db, bodyRequest)

		if err != nil {
			return c.SendStatus(500)
		}

		return c.JSON(user)
	})

}