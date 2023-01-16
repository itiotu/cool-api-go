package database

import (
	model "cool-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "root:parola@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&model.User{},
		&model.Basket{},
		&model.BasketProduct{},
		&model.Product{},
	)

	return db, nil
}
