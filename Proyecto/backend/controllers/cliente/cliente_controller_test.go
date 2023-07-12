package clienteController_test

import (
	"backend/controllers/clienteController"
	"backend/dto"
	"backend/service"
	"backend/utils/errors"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestClienteService struct {
	GetClienteByIdFunc     func(id int) (dto.ClienteDto, errors.APIError)
	GetClienteByUsernameFunc func(username string) (dto.ClienteDto, errors.APIError)
	// Implementa los demás métodos que necesites para tus pruebas
}

func (s *TestClienteService) GetClienteById(id int) (dto.ClienteDto, errors.APIError) {
	return s.GetClienteByIdFunc(id)
}

func (s *TestClienteService) GetClienteByUsername(username string) (dto.ClienteDto, errors.APIError) {
	return s.GetClienteByUsernameFunc(username)
}

// Implementa los demás métodos del servicio según sea necesario

func TestGetClienteById(t *testing.T) {
	// Configura los datos de prueba
	id := 1
	expectedClient := dto.ClienteDto{
		ID:       id,
		Username: "testuser",
		Email:    "test@example.com",
		// Completa los demás campos del cliente según tu estructura
	}

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		GetClienteByIdFunc: func(id int) (dto.ClienteDto, errors.APIError) {
			if id == expectedClient.ID {
				return expectedClient, nil
			}
			return dto.ClienteDto{}, errors.NewNotFoundAPIError("Cliente no encontrado")
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.GET("/cliente/:id", clienteController.GetClienteById)

	// Crea una solicitud HTTP GET al endpoint /cliente/{id}
	request, _ := http.NewRequest("GET", "/cliente/"+strconv.Itoa(id), nil)

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto ClienteDto
	var result dto.ClienteDto
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica que los datos del cliente sean correctos
	assert.Equal(t, expectedClient.ID, result.ID, "ID del cliente incorrecto")
	assert.Equal(t, expectedClient.Username, result.Username, "Nombre de usuario del cliente incorrecto")
	assert.Equal(t, expectedClient.Email, result.Email, "Correo electrónico del cliente incorrecto")
	// Verifica los demás campos del cliente según tu estructura
}

// Implementa los demás tests unitarios para los otros métodos del controlador

// Ejemplo de test para GetClienteByUsername
func TestGetClienteByUsername(t *testing.T) {
	// Configura los datos de prueba
	username := "testuser"
	expectedClient := dto.ClienteDto{
		ID:       1,
		Username: username,
		Email:    "test@example.com",
		// Completa los demás campos del cliente según tu estructura
	}

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		GetClienteByUsernameFunc: func(username string) (dto.ClienteDto, errors.APIError) {
			if username == expectedClient.Username {
				return expectedClient, nil
			}
			return dto.ClienteDto{}, errors.NewNotFoundAPIError("Cliente no encontrado")
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.GET("/cliente/username/:username", clienteController.GetClienteByUsername)

	// Crea una solicitud HTTP GET al endpoint /cliente/username/{username}
	request, _ := http.NewRequest("GET", "/cliente/username/"+username, nil)

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto ClienteDto
	var result dto.ClienteDto
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica que los datos del cliente sean correctos
	assert.Equal(t, expectedClient.ID, result.ID, "ID del cliente incorrecto")
	assert.Equal(t, expectedClient.Username, result.Username, "Nombre de usuario del cliente incorrecto")
	assert.Equal(t, expectedClient.Email, result.Email, "Correo electrónico del cliente incorrecto")
	// Verifica los demás campos del cliente según tu estructura
}

func TestGetClienteByEmail(t *testing.T) {
	// Configura los datos de prueba
	email := "test@example.com"
	expectedClient := dto.ClienteDto{
		ID:       1,
		Username: "testuser",
		Email:    email,
		// Completa los demás campos del cliente según tu estructura
	}

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		GetClienteByEmailFunc: func(email string) (dto.ClienteDto, errors.APIError) {
			if email == expectedClient.Email {
				return expectedClient, nil
			}
			return dto.ClienteDto{}, errors.NewNotFoundAPIError("Cliente no encontrado")
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.GET("/cliente/email/:email", clienteController.GetClienteByEmail)

	// Crea una solicitud HTTP GET al endpoint /cliente/email/{email}
	request, _ := http.NewRequest("GET", "/cliente/email/"+email, nil)

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto ClienteDto
	var result dto.ClienteDto
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica que los datos del cliente sean correctos
	assert.Equal(t, expectedClient.ID, result.ID, "ID del cliente incorrecto")
	assert.Equal(t, expectedClient.Username, result.Username, "Nombre de usuario del cliente incorrecto")
	assert.Equal(t, expectedClient.Email, result.Email, "Correo electrónico del cliente incorrecto")
	// Verifica los demás campos del cliente según tu estructura
}

func TestInsertCliente(t *testing.T) {
	// Configura los datos de prueba
	cliente := dto.ClienteDto{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		// Completa los demás campos del cliente según tu estructura
	}

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		InsertClienteFunc: func(clienteDto dto.ClienteDto) (dto.ClienteDto, errors.APIError) {
			// Simula la inserción exitosa del cliente
			return clienteDto, nil
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.POST("/cliente", clienteController.InsertCliente)

	// Convierte el objeto cliente en formato JSON
	jsonData, _ := json.Marshal(cliente)

	// Crea una solicitud HTTP POST al endpoint /cliente
	request, _ := http.NewRequest("POST", "/cliente", strings.NewReader(string(jsonData)))

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusCreated, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto ClienteDto
	var result dto.ClienteDto
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica que los datos del cliente sean correctos
	assert.Equal(t, cliente.ID, result.ID, "ID del cliente incorrecto")
	assert.Equal(t, cliente.Username, result.Username, "Nombre de usuario del cliente incorrecto")
	assert.Equal(t, cliente.Email, result.Email, "Correo electrónico del cliente incorrecto")
	// Verifica los demás campos del cliente según tu estructura
}

// Implementa los demás tests unitarios para las demás funciones del controlador

// Ejemplo de test para GetReservaById
func TestGetReservaById(t *testing.T) {
	// Configura los datos de prueba
	id := 1
	expectedReserva := dto.ReservaDto{
		ID:       id,
		ClientID: 1,
		// Completa los demás campos de la reserva según tu estructura
	}

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		GetReservaByIdFunc: func(id int) (dto.ReservaDto, errors.APIError) {
			if id == expectedReserva.ID {
				return expectedReserva, nil
			}
			return dto.ReservaDto{}, errors.NewNotFoundAPIError("Reserva no encontrada")
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.GET("/reserva/:id", clienteController.GetReservaById)

	// Crea una solicitud HTTP GET al endpoint /reserva/{id}
	request, _ := http.NewRequest("GET", "/reserva/"+strconv.Itoa(id), nil)

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto ReservaDto
	var result dto.ReservaDto
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica que los datos de la reserva sean correctos
	assert.Equal(t, expectedReserva.ID, result.ID, "ID de la reserva incorrecto")
	assert.Equal(t, expectedReserva.ClientID, result.ClientID, "ID de cliente de la reserva incorrecto")
}

func TestInsertReserva(t *testing.T) {
	// Configura los datos de prueba
	reserva := dto.ReservaDto{
		ID:       1,
		ClientID: 1,
		// Completa los demás campos de la reserva según tu estructura
	}

	// Configura el servicio de prueba
	mockService := &TestClienteService{
		InsertReservaFunc: func(reservaDto dto.ReservaDto) (dto.ReservaDto, errors.APIError) {
			// Simula la inserción exitosa de la reserva
			return reservaDto, nil
		},
	}

	// Configura el controlador con el servicio de prueba
	clienteController.ClienteService = mockService

	// Configura el enrutador de Gin para la prueba
	router := gin.Default()
	router.POST("/reserva", clienteController.InsertReserva)

	// Convierte el objeto reserva en formato JSON
	jsonData, _ := json.Marshal(reserva)

	// Crea una solicitud HTTP POST al endpoint /reserva
	request, _ := http.NewRequest("POST", "/reserva", strings.NewReader(string(jsonData)))

	// Crea un registrador de respuestas HTTP para capturar la respuesta
	response := httptest.NewRecorder()

	// Envía la solicitud HTTP al enrutador
	router.ServeHTTP(response, request)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusCreated, response.Code, "El código de respuesta no es el esperado")

	// Decodifica la respuesta JSON en un objeto ReservaDto
	var result dto.ReservaDto
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err, "Error al decodificar la respuesta JSON")

	// Verifica que los datos de la reserva sean correctos
	assert.Equal(t, reserva.ID, result.ID, "ID de la reserva incorrecto")
	assert.Equal(t, reserva.ClientID, result.ClientID, "ID de cliente de la reserva incorrecto")
	// Verifica los demás campos de la reserva según tu estructura
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