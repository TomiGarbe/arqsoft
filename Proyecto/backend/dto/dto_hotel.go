package dto

type HotelDto struct {
	ID       	int    	`json:"id"`
	Nombre   	string 	`json:"nombre"`
	Descripcion	string 	`json:"descripcion"`
	Email    	string 	`json:"email"`
	Imagenes  []string  `json:"imagenes"`
	Cant_Hab 	int    	`json:"cant_hab"`
	Amenities 	string 	`json:"amenities"`

	TelefonosDto TelefonosDto `json:"telefonos,omitempty"`
}


type HotelesDto []HotelDto