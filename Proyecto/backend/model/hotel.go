package model

type Hotel struct {
	ID        	int    		`gorm:"primaryKey"`
	Nombre    	string 		`gorm:"type:varchar(350);not null;unique"`
	Descripcion	string 		`gorm:"type:varchar(200);not null"`

	Telefonos 	Telefonos 	`gorm:"foreignkey:HotelID;unique"`

	Email     	string    	`gorm:"type:varchar(150);not null;unique"`
	Image     	string 		`gorm:"type:varchar(255)"`
	Cant_Hab  	int			`gorm:"not null"`
}

type Hoteles []Hotel