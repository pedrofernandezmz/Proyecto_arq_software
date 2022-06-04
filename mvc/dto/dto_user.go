package dto

type UserDto struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
}

type UsersDto []UserDto
