package models

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Phone    string `json:"phone_number"`
}
