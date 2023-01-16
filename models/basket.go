package models

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	ID uint
	TotalValue float32
	UserID uint
	BasketProduct []BasketProduct
}