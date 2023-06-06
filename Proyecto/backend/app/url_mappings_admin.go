	package app

	import (
		adminController "backend/controllers/admin"

		log "github.com/sirupsen/logrus"
	)

	func mapUrlsAdmin() {

		router.GET("/admin/:id", adminController.GetAdminById)
		router.GET("/admins", adminController.GetAdmins)
		router.POST("/admin", adminController.InsertAdmin)
		router.GET("/admin/cliente/:id", adminController.GetClienteById)
		router.GET("/admin/clientes", adminController.GetClientes)
		router.GET("/admin/hotel/:id", adminController.GetHotelById)
		router.GET("/admin/hoteles", adminController.GetHoteles)
		router.POST("/admin/hotel", adminController.InsertHotel)
		router.POST("/admin/hotel/:id/telefono", adminController.AddTelefono)
		router.GET("/admin/reservas", adminController.GetReservas)

		log.Info("Terminando configuraciones de mapeos")
	}