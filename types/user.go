package types

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

type UserCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
	Sex   string `json:"sex"`
}

func (userCreateRequest UserCreateRequest) Validate() []error {
	var errorArray []error

	nameLength := len([]rune(userCreateRequest.Name))

	if nameLength <= 1 || nameLength >= 30 {
		errorArray = append(errorArray, errors.New("Name must be between 1 and 30."))
	}

	_, err := mail.ParseAddress(userCreateRequest.Email)
	if err != nil {
		errorArray = append(errorArray, errors.New("Email is not a valid email address."))
	}

	if userCreateRequest.Age < 18 {
		errorArray = append(errorArray, errors.New("Minimum signup age is 18."))
	}

	allowedSexes := []string{"male", "female"}

	if !contains(allowedSexes, userCreateRequest.Sex) {
		errorArray = append(errorArray, errors.New(fmt.Sprintf("Sex can only have the following values: %s", strings.Join(allowedSexes, ", "))))
	}

	return errorArray
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
