package app

import (
	clienteController "backend/controllers/cliente"

	log "github.com/sirupsen/logrus"
)

func mapUrlsCliente() {

	router.GET("/cliente/:id", clienteController.GetClienteById)
	router.GET("/cliente/:username/:password", clienteController.GetClienteByUserPass)
	//router.GET("/cliente/login", clienteController.GetClienteByUserPass)
	router.POST("/cliente", clienteController.InsertCliente)
	router.GET("/cliente/hoteles", clienteController.GetHoteles)
	router.POST("/cliente/reserva", clienteController.InsertReserva)
	router.GET("/cliente/reservas", clienteController.GetReservas)
	router.GET("/cliente/reserva/:id", clienteController.GetReservaById)

	log.Info("Finishing mappings configurations")
}
/*
router.GET("/cliente/:username&:password", clienteController.GetClienteByUserPass)

func GetClienteByUserPass(username string, password string) model.Cliente {
	var cliente model.Cliente

	Db.Where("username = ?", username, "password = ?", password).First(&cliente)
	log.Debug("Cliente: ", cliente)

	return cliente
}

func GetClienteByUserPass(c *gin.Context) {
	log.Debug("Cliente  to load: " + c.Param("username") + c.Param("password"))

	username := c.Param("username")
	password := c.Param("password")
	var clienteDto dto.ClienteDto

	clienteDto, err := service.ClienteService.GetClienteByUserPass(username, password)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, clienteDto)
}

func (s *clienteService) GetClienteByUserPass(username string, password string) (dto.ClienteDto, e.ApiError) {

	var cliente model.Cliente = clienteClient.GetClienteByUserPass(username, password)
	var clienteDto dto.ClienteDto

	if cliente.UserName == "" || cliente.Password == "" {
		return clienteDto, e.NewBadRequestApiError("cliente not found")
	}

	clienteDto.Name = cliente.Name
	clienteDto.LastName = cliente.LastName
	clienteDto.Email = cliente.Email

	return clienteDto, nil
}/*