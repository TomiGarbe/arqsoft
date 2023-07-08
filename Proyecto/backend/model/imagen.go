package model

type Imagen struct {
	ID      int    `gorm:"primaryKey"`
	Url     string `gorm:"type:varchar(500);not null"`

	Hotel   Hotel  `gorm:"foreignkey:HotelID"`
	HotelID int
}

type Imagenes []Imagen