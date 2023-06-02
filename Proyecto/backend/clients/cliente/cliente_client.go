package clients

import (
	"backend/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetClienteById(id int) model.Cliente {
	var cliente model.Cliente

	Db.Where("id = ?", id).First(&cliente)
	log.Debug("Cliente: ", cliente)

	return cliente
}

func GetClienteByUserPass(username, password string) model.Cliente {
	var cliente model.Cliente

	Db.Where("username = ? AND password = ?", username, password).First(&cliente)
	log.Debug("Cliente: ", cliente)

	return cliente
}

/*func GetClienteByUserPass(usuario, contraseña string) model.Cliente {
	var cliente model.Cliente

	Db.Where("username = ? AND password = ?", usuario, contraseña).First(&cliente)
	log.Debug("Cliente: ", cliente)

	return cliente
}*/


func GetClientes() model.Clientes {
	var clientes model.Clientes

	Db.Find(&clientes)
	log.Debug("Clientes: ", clientes)

	return clientes
}

func InsertCliente(cliente model.Cliente) model.Cliente {
	result := Db.Create(&cliente)

	if result.Error != nil {
		log.Error("")
	}
	
	log.Debug("Cliente Creado: ", cliente.ID)
	return cliente
}