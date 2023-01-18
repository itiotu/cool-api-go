package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint
	Sku         *string `gorm:"index:idx_name,unique,length:20"`
	Colour      *string
	Description *string
	Price       *float32
}
