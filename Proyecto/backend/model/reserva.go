package model

type Reserva struct {
	ID        		int			`gorm:"primaryKey"`

	Hotel     		Hotel 		`gorm:"foreignkey:HotelId"`
	HotelID   		int

	Cliente   		Cliente 	`gorm:"foreignkey:ClienteId"`
	ClienteID 		int

	FechaInicio     int			`gorm:"not null"`
	FechaFinal      int			`gorm:"not null"`
	Dias     		int			`gorm:"not null"`
}

type Reservas []Reserva