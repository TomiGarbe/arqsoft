package reservaController_test

import (
	reservaController "backend/controllers/reserva"
	"backend/dto"
	service "backend/services"
	"backend/utils/errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestReservas struct {
}

func (t *TestReservas) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, errors.ApiError) {
	if reservaDto.ClienteID == 0 {
		return dto.ReservaDto{}, errors.NewApiError("Error al insertar la reserva", "reserva_insert_error", http.StatusInternalServerError, nil)
	}

	return dto.ReservaDto{}, nil
}

func (t *TestReservas) GetReservaById(id int) (dto.ReservaDto, errors.ApiError) {
	if id == 1 {
		return dto.ReservaDto{
			ID:       	1,
			ClienteID:  1,
			HotelID:  	1,
			
			AnioInicio: 2023,
			AnioFinal:	2023,
			MesInicio:	1,
			MesFinal:	1,
			DiaInicio:	10,
			DiaFinal:	13,
			Dias:		3,
		}, nil
	}

	return dto.ReservaDto{}, errors.NewApiError("Reserva not found", "reserva_not_found", http.StatusNotFound, nil)
}

// Si se cambia el valor a 0, no deja realizar la reserva
func (t *TestReservas) GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (disponibilidad int) {
	disponibilidad = 10
	return disponibilidad;
}

func (t *TestReservas) GetReservasById(id int) (dto.ReservasDto, errors.ApiError) {
	return dto.ReservasDto{}, nil
}

func (t *TestReservas) GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dto.ReservasDto, errors.ApiError) {
	return dto.ReservasDto{}, nil
}

func TestInsertReserva(t *testing.T) {
	service.ReservaService = &TestReservas{}
	router := gin.Default()

	router.POST("/reserva", reservaController.InsertReserva)

	// Solicitud HTTP POST - Si se cambia el User id a 0 se ve el error
	
	myJson := `{
		"hotel_id": 1,
		"cliente_id": 1,
		"anio_inicio": 2023,
		"anio_final": 2023,
		"mes_inicio": 1,
		"mes_final": 1,
		"dia_inicio": 10,
		"dia_final": 13,
		"dias": 3
		}`

	response := httptest.NewRecorder()

	disponibilidad := service.ReservaService.GetDisponibilidad(1, 2023, 2023, 1, 1, 10, 13)
	if disponibilidad == 0 {
		assert.Equal(t, http.StatusBadRequest, response.Code, "No se pudo lograr la disponibilidad para las fechas especificadas")
	} else {

		bodyJson := strings.NewReader(myJson)
		request, _ := http.NewRequest("POST", "/reserva", bodyJson)

		router.ServeHTTP(response, request)

		fmt.Println(response.Body.String())

		// Verificar el código de estado de la respuesta
		assert.Equal(t, http.StatusCreated, response.Code, "El código de respuesta no es el esperado")
	}
}

func TestGetBookingById(t *testing.T) {
	service.ReservaService = &TestReservas{}
	router := gin.Default()

	router.GET("/reserva/:id", reservaController.GetReservaById)

	// Crear una solicitud HTTP de tipo GET al endpoint /booking/{id}

	// Si se cambia el id a otro numero se ve el error
	request, _ := http.NewRequest("GET", "/reserva/1", nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	fmt.Println(response.Body.String())

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El ID buscado no existe")
}