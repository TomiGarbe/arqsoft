package dto

type TelefonoDto struct {
	ID      int    `json:"id"`
	Codigo  string `json:"codigo"`
	Numero  string `json:"numero"`
	HotelID int    `json:"hotel_id"`
}

type TelefonosDto []TelefonoDto