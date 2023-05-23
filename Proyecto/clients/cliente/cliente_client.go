package clients

import (
	"Proyecto/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetClienteById(id int) model.Cliente {
	var cliente model.Cliente

	log.Debug("Cliente: ", cliente)

	return cliente
}

func GetClientes() model.Clientes {
	var clientes model.Clientes

	log.Debug("Clientes: ", clientes)

	return clientes
}

func Insert(cliente model.Cliente) model.Cliente {
	result := Db.Create(&cliente)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Cliente Creado: ", cliente.ID)
	return cliente
}
