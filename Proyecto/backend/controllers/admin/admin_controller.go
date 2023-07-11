package adminController

import (
	"backend/dto"
	service "backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAdminById(c *gin.Context) {
	log.Debug("ID de administrador para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var adminDto dto.AdminDto

	adminDto, err := service.AdminService.GetAdminById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, adminDto)
}

func GetAdminByUsername(c *gin.Context) {
	log.Debug("Admin a cargar: " + c.Param("username"))

	username := c.Param("username")
	var adminDto dto.AdminDto

	adminDto, err := service.AdminService.GetAdminByUsername(username)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, adminDto)
}

func GetAdminByEmail(c *gin.Context) {
	log.Debug("Admin a cargar: " + c.Param("email"))

	email := c.Param("email")
	var adminDto dto.AdminDto

	adminDto, err := service.AdminService.GetAdminByEmail(email)

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

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	adminDto, er := service.AdminService.InsertAdmin(adminDto)
	
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, adminDto)
}

func GetClienteById(c *gin.Context) {
	log.Debug("ID de Cliente para cargar: " + c.Param("id"))

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

func GetHotelByEmail(c *gin.Context) {
	log.Debug("Email de Hotel para cargar: " + c.Param("email"))

	email := c.Param("email")
	var hotelDto dto.HotelDto

	hotelDto, err := service.AdminService.GetHotelByEmail(email)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHotelByNombre(c *gin.Context) {
	log.Debug("Nombre de Hotel para cargar: " + c.Param("nombre"))

	nombre := c.Param("nombre")
	var hotelDto dto.HotelDto

	hotelDto, err := service.AdminService.GetHotelByNombre(nombre)

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

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := service.AdminService.InsertHotel(hotelDto)
	
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func InsertImagenByHotelId(c *gin.Context) {
	log.Debug("Hotel id to insert imagen: " + c.Param("id"))
	var imagenDto dto.ImagenDto
	hotelID, erint := strconv.Atoi(c.Param("id"))
	if erint != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erint.Error()})
		return
	}
	log.Debug("1")
	imagen, err := c.FormFile("imagen")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Debug("2")
	// Guardar la imagen y manejar la lógica de relación con el hotel
	imagenDto, er := service.AdminService.InsertImageByHotelId(hotelID, imagen)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	log.Debug("3")
	c.JSON(http.StatusCreated, imagenDto)
}

func GetImagenesByHotelId(c *gin.Context) {
	log.Debug("Hotel id to load imagenes: " + c.Param("id"))

	hotelID, _ := strconv.Atoi(c.Param("id"))
	var imagenesDto dto.ImagenesDto
	
	imagenesDto, err := service.AdminService.GetImagenesByHotelId(hotelID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imagenesDto)
}

func DeleteImagenById(c *gin.Context) {
	// Obtiene el ID de la imagen de los parámetros de la solicitud
	imagenId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amenitie ID"})
		return
	}

	// Llama al servicio para eliminar la imagen por su ID
	err = service.AdminService.DeleteImagenById(imagenId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
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

func GetReservasByDate(c *gin.Context) {
	log.Debug("Reservas para cargar: " + c.Param("AnioInicio") + c.Param("MesInicio") + c.Param("DiaInicio") + c.Param("AnioFinal") + c.Param("MesFinal") + c.Param("DiaFinal"))

	AnioInicio, _ := strconv.Atoi(c.Param("AnioInicio"))
	AnioFinal, _ := strconv.Atoi(c.Param("AnioFinal"))
	MesInicio, _ := strconv.Atoi(c.Param("MesInicio"))
	MesFinal, _ := strconv.Atoi(c.Param("MesFinal"))
	DiaInicio, _ := strconv.Atoi(c.Param("DiaInicio"))
	DiaFinal, _ := strconv.Atoi(c.Param("DiaFinal"))

	var reservasDto dto.ReservasDto

	reservasDto, err := service.AdminService.GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func UpdateHotel(c *gin.Context) {
	// Obtener el ID del hotel a editar
	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	// Obtener los datos actualizados del hotel desde el cuerpo de la solicitud
	var updatedHotelDto dto.HotelDto
	err = c.BindJSON(&updatedHotelDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Obtener el hotel existente desde la base de datos
	existingHotel, err := service.AdminService.GetHotelById(hotelID)
	if existingHotel.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	// Actualizar los atributos del hotel con los nuevos valores
	existingHotel.Nombre = updatedHotelDto.Nombre
	existingHotel.Descripcion = updatedHotelDto.Descripcion
	existingHotel.Email = updatedHotelDto.Email
	existingHotel.Cant_Hab = updatedHotelDto.Cant_Hab
	existingHotel.Amenities = updatedHotelDto.Amenities

	// Guardar los cambios en la base de datos
	updatedHotel, err := service.AdminService.UpdateHotel(hotelID, existingHotel)
	c.JSON(http.StatusOK, updatedHotel)
}


/*func InsertAmenidades(c *gin.Context) {
	// Obtener el ID del hotel desde los parámetros de la solicitud
	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	// Obtener los datos de las amenidades desde el cuerpo de la solicitud
	var hotelDto dto.HotelDto
	err = c.BindJSON(&hotelDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Llamar al servicio para actualizar las amenidades del hotel
	updatedHotelDto, err := service.AdminService.UpdateHotel(hotelID, HotelDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedHotelDto)
}

func GetAmenidades(c *gin.Context) {
	// Obtener el ID del hotel desde los parámetros de la solicitud
	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	// Llamar al servicio para obtener las amenidades del hotel
	hotelDto, err := service.AdminService.GetHotelById(hotelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	amenidades := hotelDto.Amenities

	c.JSON(http.StatusOK, gin.H{"amenidades": amenidades})
}

func DeleteAmenidades(c *gin.Context) {
	// Obtener el ID del hotel desde los parámetros de la solicitud
	hotelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	// Llamar al servicio para eliminar las amenidades del hotel
	updatedHotelDto, err := service.AdminService.DeleteAmenidades(hotelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedHotelDto)
}*/

