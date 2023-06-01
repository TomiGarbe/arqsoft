package model

type Telefono struct {
	ID      int    `gorm:"primaryKey"`
	Codigo  string `gorm:"type:varchar(10);not null"`
	Numero  string `gorm:"type:varchar(25);not null"`

	HotelID int
}

type Telefonos []Telefono