package model

type Hotel struct {
	ID        	int    		`gorm:"primaryKey"`
	Nombre    	string 		`gorm:"type:varchar(350);not null;unique"`
	Descripcion	string 		`gorm:"type:text"`

	Telefonos 	Telefonos 	`gorm:"foreignkey:HotelID;unique"`

	Email     	string    	`gorm:"type:varchar(150);not null;unique"`
	Imagenes    Imagenes    `gorm:"foreignKey:HotelID"`
	Cant_Hab  	int			`gorm:"not null"`

	Amenities 	string 		`gorm:"type:varchar(1000)"`

}

type Hoteles []Hotel