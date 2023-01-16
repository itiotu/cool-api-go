package models

import "gorm.io/gorm"

type BasketProduct struct {
	gorm.Model
	ID uint
	Price float32
	BasketID uint
	ProductID uint
	Product Product
}