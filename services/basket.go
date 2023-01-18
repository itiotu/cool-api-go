package services

import (
	"cool-api/models"
	"cool-api/types"

	"gorm.io/gorm"
)

func AddToBasket(db *gorm.DB, AddtoBasketRequest *types.AddtoBasketRequest, basket *models.Basket) {
	product := models.Product{Sku: &AddtoBasketRequest.Sku}
	db.First(&product)

	basketProduct := models.BasketProduct{Price: AddtoBasketRequest.Price, BasketID: basket.ID, Product: product}

	db.Create(&basketProduct)

	basket.TotalValue += basketProduct.Price
	basket.NumberOfProducts++

	db.Save(&basket)
}

func CreateBasket(db *gorm.DB, UserID uint, basket *models.Basket) (*models.Basket, error) {

	result := db.Create(&basket)

	if result.Error == nil {
		return basket, nil
	}

	return basket, result.Error
}

func HasBasket(db *gorm.DB, UserID int, basket *models.Basket) bool {

	db.First(&basket, "user_id = ?", UserID)

	if basket.ID == 0 {
		return false
	}

	return true
}

func GetBasket(db *gorm.DB, BasketID uint, basket *models.Basket) {
	db.Preload("BasketProduct.Product").First(&basket, BasketID)
}
