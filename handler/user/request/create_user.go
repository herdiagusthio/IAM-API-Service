package request

import "iam-api-service/service/user"

type CreateUserRequest struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required, email"`
	Phone_number string `json:"phone_number" validate:"required, number"`
	Password     string `json:"password" validate:"required"`
}

func (req *CreateUserRequest) ConvertToUserData() *user.CreateUserData {
	var data user.CreateUserData

	data.Name = req.Name
	data.Email = req.Email
	data.Phone_number = req.Phone_number
	data.Password = req.Password

	return &data
}
