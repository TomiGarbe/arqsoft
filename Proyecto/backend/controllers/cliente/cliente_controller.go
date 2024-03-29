package clienteController

import (
	"backend/dto"
	service "backend/services"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetClienteById(c *gin.Context) {
	log.Debug("ID de cliente para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var clienteDto dto.ClienteDto

	clienteDto, err := service.ClienteService.GetClienteById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, clienteDto)
}

func GetClienteByUsername(c *gin.Context) {
	log.Debug("Cliente a cargar: " + c.Param("username"))

	username := c.Param("username")
	var clienteDto dto.ClienteDto

	clienteDto, err := service.ClienteService.GetClienteByUsername(username)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, clienteDto)
}

func GetClienteByEmail(c *gin.Context) {
	log.Debug("Cliente a cargar: " + c.Param("email"))

	email := c.Param("email")
	var clienteDto dto.ClienteDto

	clienteDto, err := service.ClienteService.GetClienteByEmail(email)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	token := generateToken(clienteDto)
    if err != nil {
        log.Error(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    response := struct {
        Token string      `json:"token"`
        Cliente  dto.ClienteDto `json:"cliente"`
    }{
        Token: token,
        Cliente:  clienteDto,
    }

	c.JSON(http.StatusOK, response)
}

func InsertCliente(c *gin.Context) {
	var clienteDto dto.ClienteDto
	err := c.BindJSON(&clienteDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	clienteDto, er := service.ClienteService.InsertCliente(clienteDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, clienteDto)
}

func GetReservaById(c *gin.Context) {
	log.Debug("ID de reserva para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservaDto dto.ReservaDto

	reservaDto, err := service.ClienteService.GetReservaById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, reservaDto)
}

func GetReservasById(c *gin.Context) {
	log.Debug("ID de reserva para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservasDto dto.ReservasDto

	reservasDto, err := service.ClienteService.GetReservasById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func InsertReserva(c *gin.Context) {
	var reservaDto dto.ReservaDto
	err := c.BindJSON(&reservaDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	reservaDto, er := service.ClienteService.InsertReserva(reservaDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservaDto)
}

func GetHoteles(c *gin.Context) {
	var hotelesDto dto.HotelesDto
	hotelesDto, err := service.ClienteService.GetHoteles()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelesDto)
}

func GetImagenesByHotelId(c *gin.Context) {
	log.Debug("Hotel id to load imagenes: " + c.Param("id"))

	hotelID, _ := strconv.Atoi(c.Param("id"))
	var imagenesDto dto.ImagenesDto
	
	imagenesDto, err := service.ClienteService.GetImagenesByHotelId(hotelID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imagenesDto)
}

func GetHotelById(c *gin.Context) {
	log.Debug("ID de Hotel para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.AdminService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetDisponibilidad(c *gin.Context) {
	log.Debug("Disponibilidad de reservas para cargar: " + c.Param("id") + c.Param("AnioInicio") + c.Param("MesInicio") + c.Param("DiaInicio") + c.Param("AnioFinal") + c.Param("MesFinal") + c.Param("DiaFinal"))

	id, _ := strconv.Atoi(c.Param("id"))
	AnioInicio, _ := strconv.Atoi(c.Param("AnioInicio"))
	AnioFinal, _ := strconv.Atoi(c.Param("AnioFinal"))
	MesInicio, _ := strconv.Atoi(c.Param("MesInicio"))
	MesFinal, _ := strconv.Atoi(c.Param("MesFinal"))
	DiaInicio, _ := strconv.Atoi(c.Param("DiaInicio"))
	DiaFinal, _ := strconv.Atoi(c.Param("DiaFinal"))
	

	disponibilidad := service.ClienteService.GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal)

	c.JSON(http.StatusOK, disponibilidad)
}

func GetReservasByDate(c *gin.Context) {
	log.Debug("Reservas para cargar: " + c.Param("AnioInicio") + c.Param("MesInicio") + c.Param("DiaInicio") + c.Param("AnioFinal") + c.Param("MesFinal") + c.Param("DiaFinal"))

	AnioInicio, _ := strconv.Atoi(c.Param("AnioInicio"))
	AnioFinal, _ := strconv.Atoi(c.Param("AnioFinal"))
	MesInicio, _ := strconv.Atoi(c.Param("MesInicio"))
	MesFinal, _ := strconv.Atoi(c.Param("MesFinal"))
	DiaInicio, _ := strconv.Atoi(c.Param("DiaInicio"))
	DiaFinal, _ := strconv.Atoi(c.Param("DiaFinal"))

	var reservasDto dto.ReservasDto

	reservasDto, err := service.ClienteService.GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func generateToken(loginDto dto.ClienteDto) (string) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = loginDto.ID
	claims["expiration"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return ""
	}

	return tokenString
}