package clients

import (
	"Proyecto/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func AddTelefono(telefono model.Telefono) model.Telefono {
	result := Db.Create(&telefono)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Telefono Creado: ", telefono.ID)
	return telefono
}
