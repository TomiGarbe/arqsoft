package db

import (
	adminClient "Proyecto/clients/admin"
	telefonoClient "Proyecto/clients/telefono"
	clienteClient "Proyecto/clients/cliente"
	hotelClient "Proyecto/clients/hotel"
	reservaClient "Proyecto/clients/reserva"

	"Proyecto/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "arqsoft"
	DBUser := "root"
	DBPass := ""
	DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	
	adminClient.Db = db
	clienteClient.Db = db
	hotelClient.Db = db
	telefonoClient.Db = db
	reservaClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Cliente{})
	db.AutoMigrate(&model.Hotel{})
	db.AutoMigrate(&model.Reserva{})
	db.AutoMigrate(&model.Telefono{})

	log.Info("Finishing Migration Database Tables")
}
