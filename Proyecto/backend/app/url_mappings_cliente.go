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
	router.GET("/cliente/imagenes/hotel/:id", clienteController.GetImagenesByHotelId)
	router.GET("/cliente/hotel/:id", clienteController.GetHotelById)
	router.POST("/cliente/reserva", clienteController.InsertReserva)
	router.GET("/cliente/reservas/:id", clienteController.GetReservasById)
	router.GET("/cliente/reserva/:id", clienteController.GetReservaById)
	router.GET("/cliente/disponibilidad/:id/:AnioInicio/:MesInicio/:DiaInicio/:AnioFinal/:MesFinal/:DiaFinal", clienteController.GetDisponibilidad)
	router.GET("/cliente/reservas-por-fecha/:AnioInicio/:MesInicio/:DiaInicio/:AnioFinal/:MesFinal/:DiaFinal", clienteController.GetReservasByDate)
	
	log.Info("Terminando configuraciones de mapeos")
}