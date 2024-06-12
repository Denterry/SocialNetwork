package model

type FullUser struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
