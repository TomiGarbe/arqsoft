package reservaController_test

import (
	clienteController "backend/controllers/cliente"
	"backend/dto"
	"backend/services"
	"backend/utils/errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/service"
	"github.com/stretchr/testify/assert"
)

type TestBookings struct {
}

/*InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservasById(id int) (dto.ReservasDto, e.ApiError)
	GetReservaById(id int) (dto.ReservaDto, e.ApiError)
	GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (disponibilidad int)
	GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dto.ReservasDto, e.ApiError)*/

func (t *TestBookings) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {
	if reservaDto.ClienteID == 0 {
		return dto.ReservaDto{}, errors.NewApiError("Error al insertar la reserva", "reserva_insert_error", http.StatusInternalServerError, nil)
	}

	return dto.ReservaDto{}, nil
}

func (t *TestBookings) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {
	if id == 1 {
		return dto.ReservaDto{
			ID:       	1,
			ClienteID:  1,
			HotelID:  	1,
			
			AnioInicio: 2023,
			AnioFinal:	2023,
			MesInicio:	7,
			MesFinal:	7,
			DiaInicio:	12,
			DiaFinal:	14,
			Dias:		2,
		}, nil
	}

	return dto.ReservaDto{}, errors.NewApiError("Booking not found", "booking_not_found", http.StatusNotFound, nil)
}

func (t *TestBookings) GetBookings() (dto.BookingsDto, errors.ApiError) {
	return dto.BookingsDto{}, nil
}

// Si se cambia el valor a 0, no deja realizar la reserva
func (t *TestBookings) RoomsAvailable(bookingDto dto.BookingDto) (dto.RoomsAvailable, errors.ApiError) {
	return dto.RoomsAvailable{Rooms: 1}, nil
}

func (t *TestBookings) GetBookingsByUserId(id int) (dto.BookingsDto, errors.ApiError) {
	return dto.BookingsDto{}, nil
}

func (t *TestBookings) GetUnavailableDatesByHotel(hotelID int) ([]time.Time, error) {
	return []time.Time{}, nil
}

// Si cambio esta funcion a false, puedo ver el error
func (t *TestBookings) CheckAvailability(hotelID int, dateFrom, dateTo time.Time) bool {
	return true
}

func TestInsertReserva(t *testing.T) {
	service.ClienteService = &TestReservas{}
	router := gin.Default()

	router.POST("/reserva", clienteController.InsertReserva)

	// Solicitud HTTP POST - Si se cambia el User id a 0 se ve el error
	myJson := `{
		"hotel_id": 1,
		"cliente_id": 1,
		"date_from": "2023/05/30",
		"date_to": "2023/06/05"
		}`

	response := httptest.NewRecorder()

	availability := service.BookingService.CheckAvailability(2, time.Date(2023, 5, 30, 0, 0, 0, 0, time.UTC), time.Date(2023, 6, 5, 0, 0, 0, 0, time.UTC))
	if !availability {
		assert.Equal(t, http.StatusBadRequest, response.Code, "No se pudo lograr la disponibilidad para las fechas especificadas")
	} else {

		bodyJson := strings.NewReader(myJson)
		request, _ := http.NewRequest("POST", "/booking", bodyJson)

		router.ServeHTTP(response, request)

		fmt.Println(response.Body.String())

		// Verificar el código de estado de la respuesta
		assert.Equal(t, http.StatusCreated, response.Code, "El código de respuesta no es el esperado")
	}
}

func TestGetBookingById(t *testing.T) {
	service.BookingService = &TestBookings{}
	router := gin.Default()

	router.GET("/booking/:id", bookingController.GetBookingById)

	// Crear una solicitud HTTP de tipo GET al endpoint /booking/{id}

	// Si se cambia el id a otro numero se ve el error
	request, _ := http.NewRequest("GET", "/booking/1", nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	fmt.Println(response.Body.String())

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El ID buscado no existe")
}