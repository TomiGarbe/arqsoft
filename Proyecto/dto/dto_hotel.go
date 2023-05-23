package dto

type HotelDto struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Telefono string `json:"telefono"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Cant_Hab int    `json:"cant_hab"`
}


type HotelesDto []HotelDto