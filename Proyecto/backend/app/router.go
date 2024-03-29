package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.Default())
	router.Static("/Imagenes", "./Imagenes")
}

func StartRoute() {
	mapUrlsCliente()
	mapUrlsAdmin()
	mapUrlsReserva()

	log.Info("Iniciando Servidor")
	router.Run(":8090")

}