package app

import (
	clienteController "backend/controllers/cliente"

	log "github.com/sirupsen/logrus"
)

func mapUrlsCliente() {

	router.GET("/cliente/:id", clienteController.GetClienteById)
	router.GET("/cliente/username/:username", clienteController.GetClienteByUsername)
	router.GET("/cliente/email/:email", clienteController.GetClienteByEmail)
	router.POST("/cliente", clienteController.InsertCliente)
	router.GET("/cliente/hoteles", clienteController.GetHoteles)
	router.POST("/cliente/reserva", clienteController.InsertReserva)
	router.GET("/cliente/reservas", clienteController.GetReservas)
	router.GET("/cliente/reserva/:id", clienteController.GetReservaById)

	log.Info("Terminando configuraciones de mapeos")
}