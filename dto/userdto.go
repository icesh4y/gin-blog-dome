package dto

import "essential/models"

type UserDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func UserToDto(user models.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
