package adminController

import (
	"Proyecto/dto"
	service "Proyecto/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAdminById(c *gin.Context) {
	log.Debug("Admin id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var adminDto dto.AdminDto

	adminDto, err := service.AdminService.GetAdminById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, adminDto)
}

func GetAdmins(c *gin.Context) {
	var adminsDto dto.AdminsDto
	adminsDto, err := service.AdminService.GetAdmins()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, adminsDto)
}

func InsertAdmin(c *gin.Context) {
	var adminDto dto.AdminDto
	err := c.BindJSON(&adminDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	adminDto, er := service.AdminService.InsertAdmin(adminDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, adminDto)
}

func GetClienteById(c *gin.Context) {
	log.Debug("Cliente id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var clienteDto dto.ClienteDto

	clienteDto, err := service.AdminService.GetClienteById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, clienteDto)
}

func GetClientes(c *gin.Context) {
	var clientesDto dto.ClientesDto
	clientesDto, err := service.AdminService.GetClientes()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, clientesDto)
}

func GetReservas(c *gin.Context) {
	var reservasDto dto.ReservasDto
	reservasDto, err := service.AdminService.GetReservas()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func GetHotelById(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.AdminService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHoteles(c *gin.Context) {
	var hotelesDto dto.HotelesDto
	hotelesDto, err := service.AdminService.GetHoteles()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelesDto)
}

func InsertHotel(c *gin.Context) {
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := service.AdminService.InsertHotel(hotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func AddTelefono(c *gin.Context) {

	log.Debug("Adding Telephone to hotel: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var telefonoDto dto.TelefonoDto
	err := c.BindJSON(&telefonoDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	telefonoDto.HotelID = id

	var hotelDto dto.HotelDto

	hotelDto, er := service.AdminService.AddTelefono(telefonoDto)
	//Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}