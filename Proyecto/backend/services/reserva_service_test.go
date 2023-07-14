package services_test

import (
	"backend/dto"
	service "backend/services"
	e "backend/utils/errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestReservas struct {
}

func (t *TestReservas) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {
	if reservaDto.ClienteID == 0 {
		return dto.ReservaDto{}, e.NewApiError("Error al insertar la reserva", "reserva_insert_error", http.StatusInternalServerError, nil)
	}

	return reservaDto, nil
}

func (t *TestReservas) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {
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

	return dto.ReservaDto{}, e.NewNotFoundApiError("Booking not found")
}

func (t *TestReservas) GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (disponibilidad int) {
	disponibilidad = 10
	return disponibilidad;
}

func (t *TestReservas) GetReservasById(id int) (dto.ReservasDto, e.ApiError) {
	return dto.ReservasDto{}, nil
}

func (t *TestReservas) GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dto.ReservasDto, e.ApiError) {
	return dto.ReservasDto{}, nil
}

func TestInsertBooking(t *testing.T) {
	// Si cambio el valor de los id puedo ver los errores
	reserva := dto.ReservaDto{
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
	}

	service.ReservaService = &TestReservas{}

	disponibilidad := service.ReservaService.GetDisponibilidad(reserva.HotelID, reserva.AnioInicio, reserva.AnioFinal, reserva.MesInicio, reserva.MesFinal, reserva.DiaInicio, reserva.DiaFinal)

	if disponibilidad == 0 {
		assert.Fail(t, "No se puede realizar la reserva debido a la falta de disponibilidad")
	} else {
		createdReserva, err := service.ReservaService.InsertReserva(reserva)

		assert.Nil(t, err, "Error al insertar la reserva")
		assert.Equal(t, 1, createdReserva.ClienteID, "El ID de usuario no coincide")
		assert.Equal(t, 1, createdReserva.HotelID, "El ID de hotel no coincide")
		assert.Equal(t, 2023, createdReserva.AnioInicio, "La fecha de inicio no coincide")
		assert.Equal(t, 1, createdReserva.MesInicio, "La fecha de inicio no coincide")
		assert.Equal(t, 10, createdReserva.DiaInicio, "La fecha de inicio no coincide")
		assert.Equal(t, 2023, createdReserva.AnioFinal, "La fecha de fin no coincide")
		assert.Equal(t, 1, createdReserva.MesFinal, "La fecha de fin no coincide")
		assert.Equal(t, 13, createdReserva.DiaFinal, "La fecha de fin no coincide")
		assert.Equal(t, 3, createdReserva.Dias, "La cantidad de dias no coincide")
	}
}

func TestGetReservaById(t *testing.T) {
	service.ReservaService = &TestReservas{}

	// Si cambio los valores de aca puedo ver los errores
	expectedReserva := dto.ReservaDto{
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
	}

	catchedReserva, err := service.ReservaService.GetReservaById(expectedReserva.ID)

	assert.Nil(t, err)
	assert.Equal(t, expectedReserva.ID, catchedReserva.ID, "El ID de reserva no coincide")
	assert.Equal(t, expectedReserva.ClienteID, catchedReserva.ClienteID, "El ID de usuario no coincide")
	assert.Equal(t, expectedReserva.HotelID, catchedReserva.HotelID, "El ID de hotel no coincide")
	assert.Equal(t, expectedReserva.AnioInicio, catchedReserva.AnioInicio, "La fecha de inicio no coincide")
	assert.Equal(t, expectedReserva.MesInicio, catchedReserva.MesInicio, "La fecha de inicio no coincide")
	assert.Equal(t, expectedReserva.DiaInicio, catchedReserva.DiaInicio, "La fecha de inicio no coincide")
	assert.Equal(t, expectedReserva.AnioFinal, catchedReserva.AnioFinal, "La fecha de fin no coincide")
	assert.Equal(t, expectedReserva.MesFinal, catchedReserva.MesFinal, "La fecha de fin no coincide")
	assert.Equal(t, expectedReserva.DiaFinal, catchedReserva.DiaFinal, "La fecha de fin no coincide")
	assert.Equal(t, expectedReserva.Dias, catchedReserva.Dias, "La cantidad de dias no coincide")
}