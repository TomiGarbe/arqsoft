package dto

type ReservaDto struct {
	ID             int    `json:"id"`
	
	HotelID   	   int 	  `json:"hotel_id"`

	ClienteID 	   int	  `json:"cliente_id"`

	FechaInicio    int 	  `json:"fecha_inicio"`
	FechaFinal     int 	  `json:"fecha_final"`
	Dias           int    `json:"dias"`
	Disponibilidad int    `json:"disponibilidad"`
}

type ReservasDto []ReservaDto