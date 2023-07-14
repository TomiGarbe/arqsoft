package app

import (
	reservaController "backend/controllers/reserva"

	log "github.com/sirupsen/logrus"
)

func mapUrlsReserva() {

	router.POST("/reserva", reservaController.InsertReserva)
	router.GET("/reservas/:id", reservaController.GetReservasById)
	router.GET("/reserva/:id", reservaController.GetReservaById)
	router.GET("/disponibilidad/:id/:AnioInicio/:MesInicio/:DiaInicio/:AnioFinal/:MesFinal/:DiaFinal", reservaController.GetDisponibilidad)
	router.GET("/reservas-por-fecha/:AnioInicio/:MesInicio/:DiaInicio/:AnioFinal/:MesFinal/:DiaFinal", reservaController.GetReservasByDate)
	
	log.Info("Terminando configuraciones de mapeos")
}