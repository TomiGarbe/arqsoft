package clients

import (
	"Proyecto/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetAdminById(id int) model.Admin {
	var admin model.Admin

	log.Debug("Admin: ", admin)

	return admin
}

func GetAdmins() model.Admins {
	var admins model.Admins
	Db.Find(&admins)

	log.Debug("Administradores: ", admins)

	return admins
}

func InsertAdmin(admin model.Admin) model.Admin {
	result := Db.Create(&admin)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Administrador Creado: ", admin.ID)
	return admin
}
