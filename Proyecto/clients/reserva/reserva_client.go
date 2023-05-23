package clients

import (
	"Proyecto/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetReservaById(id int) model.Reserva {
	var reserva model.Reserva

	log.Debug("Reserva: ", reserva)

	return reserva
}

func GetReservas() model.Reservas {
	var reservas model.Reservas

	log.Debug("Reservas: ", reservas)

	return reservas
}

func InsertReserva(reserva model.Reserva) model.Reserva {
	result := Db.Create(&reserva)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Reserva Creada: ", reserva.ID)
	return reserva
}
