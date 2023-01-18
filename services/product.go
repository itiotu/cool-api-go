package services

import (
	"cool-api/models"
	"cool-api/types"

	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, ProductCreateRequest types.ProductCreateRequest) (models.Product, error) {
	product := models.Product{
		Sku:         &ProductCreateRequest.Sku,
		Colour:      &ProductCreateRequest.Colour,
		Description: &ProductCreateRequest.Description,
		Price:       &ProductCreateRequest.Price,
	}

	result := db.Create(&product)

	if result.Error == nil {
		return product, nil
	}

	return product, result.Error

}

func GetProduct(db *gorm.DB, ProductID int) models.Product {
	product := new(models.Product)

	db.First(&product, ProductID)

	return *product
}

func GetProductBySku(db *gorm.DB, Sku string) models.Product {
	product := models.Product{Sku: &Sku}

	db.First(&product)

	return product
}
