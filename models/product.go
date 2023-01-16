package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID uint
	Sku string
	Colour string
	Description string
	Price float32
}