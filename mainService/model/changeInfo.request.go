package model

type ChangeInfoRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Birthday string `json:"birthday"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone_number"`
}
