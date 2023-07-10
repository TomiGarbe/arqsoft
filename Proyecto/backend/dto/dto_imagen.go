package dto

type ImagenDto struct {
	ID      int    `json:"id"`
	Url    string `json:"url"`

	HotelID int    `json:"hotel_id"`
}

type ImagenesDto []ImagenDto