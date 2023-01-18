package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID     uint
	Name   *string
	Email  *string
	Age    *uint8
	Sex    *string
	Basket Basket
}

func (user User) Validate(db *gorm.DB) {
	if *user.Age <= 18 {
		db.AddError(errors.New("age need to be 18+"))
	}
}
