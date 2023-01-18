package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"cool-api/models"
	"cool-api/services"
	"cool-api/types"
	"strconv"
)

func RegisterBasketRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/basket/:userId/addToBasket", func(c *fiber.Ctx) error {

		bodyRequest := &types.AddtoBasketRequest{}

		if err := c.BodyParser(&bodyRequest); err != nil {
			return err
		}

		UserID, err := strconv.Atoi(c.Params("userId"))

		if err != nil {
			return err
		}

		basket := models.Basket{UserID: uint(UserID)}

		if !services.HasBasket(db, UserID, &basket) {
			services.CreateBasket(db, uint(UserID), &basket)
		}
		fmt.Println(basket.ID)
		services.AddToBasket(db, bodyRequest, &basket)

		User := services.GetUser(db, UserID)

		return c.JSON(User)
	})

	app.Get("/basket/:basketId", func(c *fiber.Ctx) error {
		BasketID, err := strconv.Atoi(c.Params("basketId"))

		if err != nil {
			return c.SendStatus(400)
		}

		basket := models.Basket{}

		services.GetBasket(db, uint(BasketID), &basket)

		if basket.ID == 0 {
			return c.SendStatus(404)
		}

		return c.JSON(basket)
	})

}
