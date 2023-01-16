package services

import (
	. "cool-api/models"
	. "cool-api/types"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, UserCreateRequest *UserCreateRequest) (User, error) {
	user := User{Name: &UserCreateRequest.Name, Email: &UserCreateRequest.Email, Age: &UserCreateRequest.Age, Sex: &UserCreateRequest.Sex}

	result := db.Create(&user)

	if result.Error == nil {
		return user, nil
	}

	return user, result.Error

}
