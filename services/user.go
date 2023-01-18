package services

import (
	"cool-api/models"
	"cool-api/types"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, UserCreateRequest *types.UserCreateRequest) (models.User, error) {
	user := models.User{Name: &UserCreateRequest.Name, Email: &UserCreateRequest.Email, Age: &UserCreateRequest.Age, Sex: &UserCreateRequest.Sex}

	result := db.Create(&user)

	if result.Error == nil {
		return user, nil
	}

	return user, result.Error

}

func GetUser(db *gorm.DB, UserID int) models.User {
	user := new(models.User)

	db.Model(&models.User{}).Preload("Basket.BasketProduct.Product").First(&user, UserID)

	return *user
}
