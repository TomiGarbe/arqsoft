package dto

type ImagenDto struct {
	Id      int    `json:"id"`
	Url    string `json:"url"`

	HotelId int    `json:"hotel_id"`
}

type ImagenesDto struct {
	Imagenes []ImagenDto `json:"images"`
}