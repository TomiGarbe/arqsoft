package dto

type ReservaDto struct {
	ID             int    `json:"id"`
	
	Nombre   	   string `json:"nombre"`

	Name     	   string `json:"name"`
	LastName 	   string `json:"last_name"`

	FechaInicio    string `json:"fecha_inicio"`
	FechaFinal     string `json:"fecha_final"`
	Dias           int    `json:"dias"`
	Disponibilidad int    `json:"disponibilidad"`
}

type ReservasDto []ReservaDto