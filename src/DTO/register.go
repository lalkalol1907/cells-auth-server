package DTO

type RegisterDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`

	Surname string `json:"surname"`
}
