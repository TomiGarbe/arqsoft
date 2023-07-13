package clienteController_test

import (
	clienteController "backend/controllers/cliente"
	"backend/dto"
	clienteService "backend/services"
	"backend/utils/errors"
	"fmt"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Implementa el servicio de prueba para InsertReserva, GetHoteles, GetImagenesByHotelId, GetHotelById, GetDisponibilidad, GetReservasByDate y generateToken según sea necesario

type TestReservas struct {
}

func (t *TestReservas) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {
	if reservaDto.ClienteID == 0{
		return dto.ReservaDto{}, errors.NewApiError("Error al insertar la reserva", "reserva_insert_error", http.StatusInternalServerError, nil)
	}
	
	return dto.ReservaDto{}, nil
}

func (t *TestReservas) GetReservasById(id int) (dto.ReservasDto, e.ApiError) {
	if id == 1 {
		return dto.ReservasDto{
			ID:       1,
			HotelID: 1,
			ClienteID: 1,
			AnioInicio: 2023,
			AnioFinal: 2023,
			MesInicio: 1,
			MesFinal:  1,
			DiaInicio: 10,
			DiaFinal: 13,
			Dias: 3,
		}, nil
	}
}		

func (t *TestReservas) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {
	if id == 1 {
		return dto.ReservaDto{
			ID:       1,
			HotelID: 1,
			ClienteID: 1,
			AnioInicio: 2023,
			AnioFinal:   2023,
			MesInicio:   1,
			MesFinal:  1,
			DiaInicio: 10,
			DiaFinal: 13,
			Dias: 3,
		}, nil
	}
}


// Si cambio esta funcion a false, puedo ver el error
func (t *TestReservas) GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (disponibilidad int) {
	return true
}

func (t *TestReservas) GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dto.ReservasDto, e.ApiError) {


}

func TestGetHoteles(t *testing.T) {
	// Configura los datos de prueba
	expectedHoteles := dto.HotelesDto{
		// Completa los datos de los hoteles según tu estructura
	}

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		GetHotelesFunc: func() (dto.HotelesDto, errors.APIError) {
			return expectedHoteles, nil
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.GET("/hoteles", clienteController.GetHoteles)

	// Crea una solicitud HTTP GET al endpoint /hoteles
	request, _ := http.NewRequest("GET", "/hoteles", nil)

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto HotelesDto
	var result dto.HotelesDto
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica los datos de los hoteles
	assert.Equal(t, expectedHoteles, result, "Los datos de los hoteles no son los esperados")
}

func TestGetImagenesByHotelId(t *testing.T) {
	// Configura los datos de prueba
	hotelID := 1
	expectedImagenes := dto.ImagenesDto{
		HotelID: hotelID,
		// Completa los demás campos de las imágenes según tu estructura
	}

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		GetImagenesByHotelIdFunc: func(id int) (dto.ImagenesDto, errors.APIError) {
			if id == hotelID {
				return expectedImagenes, nil
			}
			return dto.ImagenesDto{}, errors.NewNotFoundAPIError("Imágenes no encontradas")
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.GET("/imagenes/:id", clienteController.GetImagenesByHotelId)

	// Crea una solicitud HTTP GET al endpoint /imagenes/{id}
	request, _ := http.NewRequest("GET", "/imagenes/"+strconv.Itoa(hotelID), nil)

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto ImagenesDto
	var result dto.ImagenesDto
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica que los datos de las imágenes sean correctos
	assert.Equal(t, expectedImagenes.HotelID, result.HotelID, "ID de hotel de las imágenes incorrecto")
	// Verifica los demás campos de las imágenes según tu estructura
}

func TestGetHotelById(t *testing.T) {
	// Configura los datos de prueba
	hotelID := 1
	expectedHotel := dto.HotelDto{
		ID: hotelID,
		// Completa los demás campos del hotel según tu estructura
	}

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		GetHotelByIdFunc: func(id int) (dto.HotelDto, errors.APIError) {
			if id == hotelID {
				return expectedHotel, nil
			}
			return dto.HotelDto{}, errors.NewNotFoundAPIError("Hotel no encontrado")
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.GET("/hotel/:id", clienteController.GetHotelById)

	// Crea una solicitud HTTP GET al endpoint /hotel/{id}
	request, _ := http.NewRequest("GET", "/hotel/"+strconv.Itoa(hotelID), nil)

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto HotelDto
	var result dto.HotelDto
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica que los datos del hotel sean correctos
	assert.Equal(t, expectedHotel.ID, result.ID, "ID del hotel incorrecto")
	// Verifica los demás campos del hotel según tu estructura
}

func TestGetDisponibilidad(t *testing.T) {
	// Configura los datos de prueba
	hotelID := 1
	AnioInicio := 2023
	AnioFinal := 2023
	MesInicio := 6
	MesFinal := 6
	DiaInicio := 1
	DiaFinal := 5
	expectedDisponibilidad := true

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		GetDisponibilidadFunc: func(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) bool {
			if id == hotelID && AnioInicio == 2023 && AnioFinal == 2023 && MesInicio == 6 && MesFinal == 6 && DiaInicio == 1 && DiaFinal == 5 {
				return expectedDisponibilidad
			}
			return false
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.GET("/disponibilidad/:id/:AnioInicio/:AnioFinal/:MesInicio/:MesFinal/:DiaInicio/:DiaFinal", clienteController.GetDisponibilidad)

	// Crea una solicitud HTTP GET al endpoint /disponibilidad/{id}/{AnioInicio}/{AnioFinal}/{MesInicio}/{MesFinal}/{DiaInicio}/{DiaFinal}
	request, _ := http.NewRequest("GET", "/disponibilidad/"+strconv.Itoa(hotelID)+"/"+strconv.Itoa(AnioInicio)+"/"+strconv.Itoa(AnioFinal)+"/"+strconv.Itoa(MesInicio)+"/"+strconv.Itoa(MesFinal)+"/"+strconv.Itoa(DiaInicio)+"/"+strconv.Itoa(DiaFinal), nil)

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto bool
	var result bool
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica que la disponibilidad sea correcta
	assert.Equal(t, expectedDisponibilidad, result, "La disponibilidad es incorrecta")
}

func TestGetReservasByDate(t *testing.T) {
	// Configura los datos de prueba
	AnioInicio := 2023
	AnioFinal := 2023
	MesInicio := 6
	MesFinal := 6
	DiaInicio := 1
	DiaFinal := 5
	expectedReservas := dto.ReservasDto{
		// Completa los datos de las reservas según tu estructura
	}

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		GetReservasByDateFunc: func(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dto.ReservasDto, errors.APIError) {
			return expectedReservas, nil
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.GET("/reservas/:AnioInicio/:AnioFinal/:MesInicio/:MesFinal/:DiaInicio/:DiaFinal", clienteController.GetReservasByDate)

	// Crea una solicitud HTTP GET al endpoint /reservas/{AnioInicio}/{AnioFinal}/{MesInicio}/{MesFinal}/{DiaInicio}/{DiaFinal}
	request, _ := http.NewRequest("GET", "/reservas/"+strconv.Itoa(AnioInicio)+"/"+strconv.Itoa(AnioFinal)+"/"+strconv.Itoa(MesInicio)+"/"+strconv.Itoa(MesFinal)+"/"+strconv.Itoa(DiaInicio)+"/"+strconv.Itoa(DiaFinal), nil)

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto ReservasDto
	var result dto.ReservasDto
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica los datos de las reservas
	assert.Equal(t, expectedReservas, result, "Los datos de las reservas no son los esperados")
}

func TestGenerateToken(t *testing.T) {
	// Configura los datos de prueba
	loginDto := dto.ClienteDto{
		ID: 1,
		// Completa los demás campos del cliente según tu estructura
	}
	expectedToken := "test-token"

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		GenerateTokenFunc: func(loginDto dto.ClienteDto) string {
			return expectedToken
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.GET("/token", clienteController.GenerateToken)

	// Crea una solicitud HTTP GET al endpoint /token
	request, _ := http.NewRequest("GET", "/token", nil)

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto con el token
	var result struct {
		Token string `json:"token"`
	}
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica que el token sea correcto
	assert.Equal(t, expectedToken, result.Token, "El token generado no es el esperado")
}

// Implementa los demás tests unitarios para las demás funciones del controlador
