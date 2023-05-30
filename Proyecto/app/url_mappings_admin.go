	package app

	import (
		adminController "Proyecto/controllers/admin"

		log "github.com/sirupsen/logrus"
	)

	func mapUrlsAdmin() {

		// Users Mapping
		router.GET("/admin/:id", adminController.GetAdminById)
		router.GET("/admins", adminController.GetAdmins)
		router.POST("/admin", adminController.InsertAdmin)
		router.GET("/admin/cliente/:id", adminController.GetClienteById)
		router.GET("/admin/clientes", adminController.GetClientes)
		router.GET("/admin/hotel/:ID", adminController.GetHotelById)
		router.GET("/admin/hoteles", adminController.GetHoteles)
		router.POST("/admin/hotel", adminController.InsertHotel)
		router.POST("/admin/hotel/:id/telefono", adminController.AddTelefono)
		router.GET("/admin/reservas", adminController.GetReservas)

		log.Info("Finishing mappings configurations")
	}