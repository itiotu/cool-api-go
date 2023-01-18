package routes

import (
	"cool-api/services"
	"cool-api/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(app *fiber.App, db *gorm.DB) {

	app.Post("/user", func(c *fiber.Ctx) error {

		bodyRequest := &types.UserCreateRequest{}

		if err := c.BodyParser(&bodyRequest); err != nil {
			return c.SendStatus(400)
		}

		validationErrors := bodyRequest.Validate()

		if len(validationErrors) >= 1 {

			response := struct {
				Errors []string
			}{}

			for _, err := range validationErrors {
				response.Errors = append(response.Errors, err.Error())
			}

			return c.JSON(response)
		}

		user, err := services.CreateUser(db, bodyRequest)

		if err != nil {
			return c.SendStatus(500)
		}

		return c.JSON(user)
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {

		UserID, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.SendStatus(400)
		}

		user := services.GetUser(db, UserID)

		if user.ID == 0 {
			return c.SendStatus(404)
		}

		return c.JSON(user)
	})
}
