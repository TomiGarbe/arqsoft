package dto

type ReservaDto struct {
	ID             int    `json:"id"`
	
	HotelID        int    `json:"hotel_id"`

	ClienteID      int    `json:"cliente_id"`

	FechaInicio    string `json:"fecha_inicio"`
	FechaFinal     string `json:"fecha_final"`
	Dias           int    `json:"dias"`
	Disponibilidad int    `json:"disponibilidad"`
}

type ReservasDto []ReservaDto