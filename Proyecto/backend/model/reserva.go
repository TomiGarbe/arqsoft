package model

type Reserva struct {
	ID        		int			`gorm:"primaryKey"`

	Hotel     		Hotel 		`gorm:"foreignkey:HotelID"`
	HotelID   		int

	Cliente   		Cliente 	`gorm:"foreignkey:ClienteID"`
	ClienteID 		int

	FechaInicio     int 		`gorm:"type:varchar(10)"`
	FechaFinal      int 		`gorm:"type:varchar(10)"` 
	Dias     		int 		`gorm:"type:varchar(2)"`
}
	
type Reservas []Reserva