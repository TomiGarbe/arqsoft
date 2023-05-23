package model

type Hotel struct {
	ID       int    `gorm:"primaryKey"`
	Nombre   string `gorm:"type:varchar(350);not null"`

	Telefono   		Telefono 	`gorm:"foreignkey:TelefonoId"`
	TelefonoID 		int

	Email    string `gorm:"type:varchar(150);not null"`
	Image    string `gorm:"type:varchar(255)"`
	Cant_Hab int	`gorm:"not null"`
}

type Hoteles []Hotel