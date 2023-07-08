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

/*func InsertHotel(c *gin.Context) {
	log.Info("Recibiendo solicitud POST para InsertHotel")
	var hotelDto dto.HotelDto
	
	log.Info("1")
	file, err := c.FormFile("image")
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Info("2")
	fileName := uuid.New().String()
	fileExt := filepath.Ext(file.Filename)

	filePath := "Imagenes" + "/" + fileName + fileExt

	hotelDto.Image = ""
	log.Info("3")
	if err := c.ShouldBindJSON(&hotelDto); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al enlazar los datos de la solicitud"})
		return
	}
	log.Info("4")

	hotelDto.Image = filePath
	hotelDto, e := service.AdminService.InsertHotel(hotelDto)

	//log.Infof("Datos recibidos: %+v", hotelDto)
	log.Info("5")
	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}*/

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
	var imagenDto dto.ImagenDto
	hotelID, erint := strconv.Atoi(c.Param("id"))
	err := c.BindJSON(&imagenDto)
	if erint != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erint.Error()})
		return
	}

	imagen, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Guardar la imagen y manejar la lógica de relación con el hotel
	imagenDto, er := service.AdminService.InsertImageByHotelId(hotelID, imagen)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, imagenDto)
}

func GetImagenesByHotelId(c *gin.Context) {
	log.Debug("Hotel id to load images: " + c.Param("id"))

	hotelID, _ := strconv.Atoi(c.Param("id"))
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

/*func InsertImageHotel(c *gin.Context) {
	log.Debug("Agregar Teléfono al hotel: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var hotelDto dto.HotelDto
	
	log.Info("1")
	file, err := c.FormFile("image")
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Info("2")
	fileName := uuid.New().String()
	fileExt := filepath.Ext(file.Filename)

	filePath := "Imagenes" + "/" + fileName + fileExt

	hotelDto.Image = filePath
	log.Info("3")
	if err := c.ShouldBind(&hotelDto); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al enlazar los datos de la solicitud"})
		return
	}
	log.Info("4")

	hotelDto, e := service.AdminService.InsertImageHotel(hotelDto, id)

	//log.Infof("Datos recibidos: %+v", hotelDto)
	log.Info("5")
	if e != nil {
		c.JSON(e.Status(), e)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}*/

func AddTelefono(c *gin.Context) {

	log.Debug("Agregar Teléfono al hotel: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var telefonoDto dto.TelefonoDto
	err := c.BindJSON(&telefonoDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	telefonoDto.HotelID = id

	var hotelDto dto.HotelDto

	hotelDto, er := service.AdminService.AddTelefono(telefonoDto)
	
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
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