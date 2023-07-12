package cliente_controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	cliente_controller "backend/controllers"
	"backend/dto"
	"backend/services"
)

func TestGetClienteById(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.GET("/clientes/:id", cliente_controller.GetClienteById)

	// Configuración de los servicios de prueba
	mockClienteService := &services.MockClienteService{} // Suponiendo que tienes una implementación de servicio simulada para las pruebas
	cliente_controller.SetClienteService(mockClienteService) // Inyecta el servicio en el controlador

	// Datos de prueba
	clienteID := 1
	clienteDto := dto.ClienteDto{
		ID:   clienteID,
		Name: "John Doe",
	}

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("GetClienteById", clienteID).Return(clienteDto, nil)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/clientes/"+strconv.Itoa(clienteID), nil)
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.ClienteDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación de los datos del cliente devuelto
	assert.Equal(t, clienteDto.ID, responseDto.ID)
	assert.Equal(t, clienteDto.Name, responseDto.Name)

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "GetClienteById", clienteID)
}

func TestGetClienteByUsername(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.GET("/clientes/username/:username", cliente_controller.GetClienteByUsername)

	// Configuración del servicio simulado
	mockClienteService := &services.MockClienteService{}
	cliente_controller.SetClienteService(mockClienteService)

	// Datos de prueba
	username := "john_doe"
	clienteDto := dto.ClienteDto{
		ID:       1,
		Username: username,
	}

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("GetClienteByUsername", username).Return(clienteDto, nil)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/clientes/username/"+username, nil)
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.ClienteDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación de los datos del cliente devuelto
	assert.Equal(t, clienteDto.ID, responseDto.ID)
	assert.Equal(t, clienteDto.Username, responseDto.Username)

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "GetClienteByUsername", username)
}

func TestGetClienteByEmail(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.GET("/clientes/email/:email", cliente_controller.GetClienteByEmail)

	// Configuración del servicio simulado
	mockClienteService := &services.MockClienteService{}
	cliente_controller.SetClienteService(mockClienteService)

	// Datos de prueba
	email := "john.doe@example.com"
	clienteDto := dto.ClienteDto{
		ID:    1,
		Email: email,
	}

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("GetClienteByEmail", email).Return(clienteDto, nil)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/clientes/email/"+email, nil)
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.ClienteDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación de los datos del cliente devuelto
	assert.Equal(t, clienteDto.ID, responseDto.ID)
	assert.Equal(t, clienteDto.Email, responseDto.Email)

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "GetClienteByEmail", email)
}


func TestInsertCliente(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.POST("/clientes", cliente_controller.InsertCliente)

	// Configuración del servicio simulado
	mockClienteService := &services.MockClienteService{}
	cliente_controller.SetClienteService(mockClienteService)

	// Datos de prueba
	clienteDto := dto.ClienteDto{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		// Agrega otros campos necesarios para el DTO
	}

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("InsertCliente", clienteDto).Return(clienteDto, nil)

	// Conversión del objeto DTO a JSON
	requestBody, err := json.Marshal(clienteDto)
	assert.NoError(t, err)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/clientes", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusCreated, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.ClienteDto
	err = json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación de los datos del cliente devuelto
	assert.Equal(t, clienteDto.Name, responseDto.Name)
	assert.Equal(t, clienteDto.Email, responseDto.Email)

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "InsertCliente", clienteDto)
}

func TestGetReservaById(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.GET("/reservas/:id", cliente_controller.GetReservaById)

	// Configuración del servicio simulado
	mockClienteService := &services.MockClienteService{}
	cliente_controller.SetClienteService(mockClienteService)

	// Datos de prueba
	reservaID := 1
	reservaDto := dto.ReservaDto{
		ID:       reservaID,
		ClientID: 1,
	}

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("GetReservaById", reservaID).Return(reservaDto, nil)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/reservas/"+strconv.Itoa(reservaID), nil)
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.ReservaDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación de los datos de la reserva devuelta
	assert.Equal(t, reservaDto.ID, responseDto.ID)
	assert.Equal(t, reservaDto.ClientID, responseDto.ClientID)

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "GetReservaById", reservaID)
}


func TestGetReservasById(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.GET("/reservas/cliente/:id", cliente_controller.GetReservasById)

	// Configuración del servicio simulado
	mockClienteService := &services.MockClienteService{}
	cliente_controller.SetClienteService(mockClienteService)

	// Datos de prueba
	clienteID := 1
	reservasDto := dto.ReservasDto{
		ClientID: clienteID,
		Reservas: []dto.ReservaDto{
			{
				ID:       1,
				ClientID: clienteID,
			},
			{
				ID:       2,
				ClientID: clienteID,
			},
		},
	}

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("GetReservasById", clienteID).Return(reservasDto, nil)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/reservas/cliente/"+strconv.Itoa(clienteID), nil)
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.ReservasDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación de los datos de las reservas devueltas
	assert.Equal(t, clienteID, responseDto.ClientID)
	assert.Len(t, responseDto.Reservas, len(reservasDto.Reservas))

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "GetReservasById", clienteID)
}

func TestInsertReserva(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.POST("/reservas", cliente_controller.InsertReserva)

	// Configuración del servicio simulado
	mockClienteService := &services.MockClienteService{}
	cliente_controller.SetClienteService(mockClienteService)

	// Datos de prueba
	reservaDto := dto.ReservaDto{
		ClientID: 1,
	}

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("InsertReserva", reservaDto).Return(reservaDto, nil)

	// Conversión del objeto DTO a JSON
	requestBody, err := json.Marshal(reservaDto)
	assert.NoError(t, err)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/reservas", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusCreated, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.ReservaDto
	err = json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación de los datos de la reserva devuelta
	assert.Equal(t, reservaDto.ClientID, responseDto.ClientID)
	// Realiza otras verificaciones necesarias para los campos relevantes

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "InsertReserva", reservaDto)
}

func TestGetHoteles(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.GET("/hoteles", cliente_controller.GetHoteles)

	// Configuración del servicio simulado
	mockClienteService := &services.MockClienteService{}
	cliente_controller.SetClienteService(mockClienteService)

	// Datos de prueba
	hotelesDto := dto.HotelesDto{
		Hoteles: []dto.HotelDto{
			{
				ID:   1,
				Name: "Hotel A",
			},
			{
				ID:   2,
				Name: "Hotel B",
			},
		},
	}

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("GetHoteles").Return(hotelesDto, nil)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hoteles", nil)
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.HotelesDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación de los datos de los hoteles devueltos
	assert.Len(t, responseDto.Hoteles, len(hotelesDto.Hoteles))
	// Realiza otras verificaciones necesarias para los campos relevantes

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "GetHoteles")
}

func TestGetImagenesByHotelId(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.GET("/Imagenes/hotel/:id", cliente_controller.GetImagenesByHotelId)

	// Configuración del servicio simulado
	mockClienteService := &services.MockClienteService{}
	cliente_controller.SetClienteService(mockClienteService)

	// Datos de prueba
	hotelID := 1
	imagenesDto := dto.ImagenesDto{
		Imagenes: []dto.ImagenDto{
			{
				ID:     1,
				URL:    "https://res.cloudinary.com/simpleview/image/upload/v1642787126/clients/loscabosmx/Copia_de_Copia_de_Esperanza_0010x_8dcb97e1-1c39-4cd8-8e36-326ec39d65b3.jpg",
			},
			{
				ID:     2,
				URL:    "https://res.cloudinary.com/simpleview/image/upload/v1642787126/clients/loscabosmx/Copia_de_Copia_de_Esperanza_0010x_8dcb97e1-1c39-4cd8-8e36-326ec39d65b3.jpg",
			},
		},
	}

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("GetImagenesByHotelId", hotelID).Return(imagenesDto, nil)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/imagenes/hotel/"+strconv.Itoa(hotelID), nil)
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.ImagenesDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación de los datos de las imágenes devueltas
	assert.Len(t, responseDto.Imagenes, len(imagenesDto.Imagenes))

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "GetImagenesByHotelId", hotelID)
}

func TestGetHotelById(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.GET("/hoteles/:id", cliente_controller.GetHotelById)

	// Configuración del servicio simulado
	mockAdminService := &services.MockAdminService{}
	cliente_controller.SetAdminService(mockAdminService)

	// Datos de prueba
	hotelID := 1
	hotelDto := dto.HotelDto{
		ID:   hotelID,
		Name: "Hotel A",
	}

	// Configuración del comportamiento simulado del servicio
	mockAdminService.On("GetHotelById", hotelID).Return(hotelDto, nil)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hoteles/"+strconv.Itoa(hotelID), nil)
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.HotelDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación de los datos del hotel devuelto
	assert.Equal(t, hotelDto.ID, responseDto.ID)
	assert.Equal(t, hotelDto.Name, responseDto.Name)

	// Verificación de que se llamó al método del servicio esperado
	mockAdminService.AssertCalled(t, "GetHotelById", hotelID)
}

func TestGetDisponibilidad(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.GET("/disponibilidad/:id/:AnioInicio/:MesInicio/:DiaInicio/:AnioFinal/:MesFinal/:DiaFinal", cliente_controller.GetDisponibilidad)

	// Configuración del servicio simulado
	mockClienteService := &services.MockClienteService{}
	cliente_controller.SetClienteService(mockClienteService)

	// Datos de prueba
	hotelID := 1
	anioInicio := 2023
	mesInicio := 7
	diaInicio := 1
	anioFinal := 2023
	mesFinal := 7
	diaFinal := 31

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("GetDisponibilidad", hotelID, anioInicio, anioFinal, mesInicio, mesFinal, diaInicio, diaFinal).Return(true)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/disponibilidad/"+strconv.Itoa(hotelID)+"/"+strconv.Itoa(anioInicio)+"/"+strconv.Itoa(mesInicio)+"/"+strconv.Itoa(diaInicio)+"/"+strconv.Itoa(anioFinal)+"/"+strconv.Itoa(mesFinal)+"/"+strconv.Itoa(diaFinal), nil)
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto bool
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación del resultado de disponibilidad
	assert.True(t, responseDto)

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "GetDisponibilidad", hotelID, anioInicio, anioFinal, mesInicio, mesFinal, diaInicio, diaFinal)
}

func TestGetReservasByDate(t *testing.T) {
	// Configuración del router Gin
	router := gin.Default()

	// Configuración de la ruta y el controlador
	router.GET("/reservas/:AnioInicio/:MesInicio/:DiaInicio/:AnioFinal/:MesFinal/:DiaFinal", cliente_controller.GetReservasByDate)

	// Configuración del servicio simulado
	mockClienteService := &services.MockClienteService{}
	cliente_controller.SetClienteService(mockClienteService)

	// Datos de prueba
	anioInicio := 2023
	mesInicio := 7
	diaInicio := 1
	anioFinal := 2023
	mesFinal := 7
	diaFinal := 31

	// Reservas de ejemplo
	reservasDto := dto.ReservasDto{
		Reservas: []dto.ReservaDto{
			{
				ID:       1,
				ClientID: 1,
			},
			{
				ID:       2,
				ClientID: 2,
			},
		},
	}

	// Configuración del comportamiento simulado del servicio
	mockClienteService.On("GetReservasByDate", anioInicio, anioFinal, mesInicio, mesFinal, diaInicio, diaFinal).Return(reservasDto, nil)

	// Ejecución de la solicitud HTTP de prueba
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/reservas/"+strconv.Itoa(anioInicio)+"/"+strconv.Itoa(mesInicio)+"/"+strconv.Itoa(diaInicio)+"/"+strconv.Itoa(anioFinal)+"/"+strconv.Itoa(mesFinal)+"/"+strconv.Itoa(diaFinal), nil)
	router.ServeHTTP(w, req)

	// Verificación del estado de la respuesta HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificación del cuerpo de la respuesta JSON
	var responseDto dto.ReservasDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.NoError(t, err)

	// Verificación del número de reservas devueltas
	assert.Len(t, responseDto.Reservas, len(reservasDto.Reservas))
	// Realiza otras verificaciones necesarias para los campos relevantes

	// Verificación de que se llamó al método del servicio esperado
	mockClienteService.AssertCalled(t, "GetReservasByDate", anioInicio, anioFinal, mesInicio, mesFinal, diaInicio, diaFinal)
}

func TestGenerateToken(t *testing.T) {
	// Datos de prueba
	loginDto := dto.ClienteDto{
		ID: 123,
	}

	// Llamada a la función para generar el token
	token := cliente_controller.GenerateToken(loginDto)

	// Verificación de que se haya generado un token no vacío
	assert.NotEmpty(t, token)

	// Verificación de los claims del token
	parsedToken, err := cliente_controller.ParseToken(token)
	assert.NoError(t, err)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)

	// Verificación del ID en los claims
	id, ok := claims["id"].(float64)
	assert.True(t, ok)
	assert.Equal(t, float64(loginDto.ID), id)

	// Verificación de la expiración en los claims
	expiration, ok := claims["expiration"].(float64)
	assert.True(t, ok)
	expirationTime := time.Unix(int64(expiration), 0)
	assert.WithinDuration(t, time.Now().Add(time.Hour*24), expirationTime, time.Second)
}