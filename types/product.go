package types

import (
	"cool-api/models"
	"errors"

	"gorm.io/gorm"
)

type ProductCreateRequest struct {
	Sku         string  `json:"sku"`
	Colour      string  `json:"colour"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func (productCreateRequest ProductCreateRequest) Validate(db *gorm.DB) []error {
	var errorArray []error

	skuLen := len([]rune(productCreateRequest.Sku))

	if skuLen <= 1 || skuLen >= 15 {
		errorArray = append(errorArray, errors.New("Sku must be between 1 and 15."))
	}

	product := models.Product{}
	db.First(&product, "sku = ?", productCreateRequest.Sku)

	if product.ID != 0 {
		errorArray = append(errorArray, errors.New("Sku already exists in the system."))
	}

	return errorArray
}
