package types

type UserCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
	Sex   string `json:"sex"`
}
