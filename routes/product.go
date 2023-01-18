package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"cool-api/services"
	"cool-api/types"
	"strconv"
)

func RegisterProductRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/product", func(c *fiber.Ctx) error {

		bodyRequest := types.ProductCreateRequest{}

		if err := c.BodyParser(&bodyRequest); err != nil {
			return err
		}

		validationErrors := bodyRequest.Validate(db)

		if len(validationErrors) >= 1 {

			response := struct {
				Errors []string
			}{}

			for _, err := range validationErrors {
				response.Errors = append(response.Errors, err.Error())
			}

			return c.JSON(response)
		}

		product, err := services.CreateProduct(db, bodyRequest)

		if err != nil {
			return c.SendStatus(500)
		}

		return c.JSON(product)
	})

	app.Get("/product/:id", func(c *fiber.Ctx) error {

		ProductID, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.SendStatus(400)
		}

		product := services.GetProduct(db, ProductID)

		if product.ID == 0 {
			return c.SendStatus(404)
		}

		return c.JSON(product)
	})
}
