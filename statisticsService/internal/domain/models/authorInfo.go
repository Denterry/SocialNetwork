package models

type MSResponse struct {
	Message string     `json:"message"`
	Data    AuthorInfo `json:"data"`
}

type AuthorInfo struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
