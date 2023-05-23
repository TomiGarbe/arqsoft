package dto

type ClienteDto struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type ClientesDto []ClienteDto