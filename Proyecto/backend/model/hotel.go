package model

type Hotel struct {
	ID        int    	`gorm:"primaryKey;autoIncrement"`
	Nombre    string 	`gorm:"type:varchar(350);not null"`

	Telefonos Telefonos `gorm:"foreignkey:HotelID"`

	Email     string    `gorm:"type:varchar(150);not null"`
	Image     string 	`gorm:"type:varchar(255)"`
	Cant_Hab  int		`gorm:"not null"`
}

type Hoteles []Hotel