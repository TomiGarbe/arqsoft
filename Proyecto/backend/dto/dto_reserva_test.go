package dto_test

import (
	"backend/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReservaDto(t *testing.T) {
	// Crear una instancia del DTO de Booking, si modifico alguna y deja de ser igual, da la alerta
	reservaDto := dto.ReservaDto{
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
	}

	// Verificar los valores de los campos del DTO de Booking
	assert.Equal(t, 1, reservaDto.ID, "El ID de la reserva no coincide")
	assert.Equal(t, 1, reservaDto.HotelID, "El ID del usuario no coincide")
	assert.Equal(t, 1, reservaDto.ClienteID, "El ID del hotel no coincide")
	assert.Equal(t, 2023, reservaDto.AnioInicio, "La fecha de inicio no coincide")
	assert.Equal(t, 1, reservaDto.MesInicio, "La fecha de inicio no coincide")
	assert.Equal(t, 10, reservaDto.DiaInicio, "La fecha de inicio no coincide")
	assert.Equal(t, 2023, reservaDto.AnioFinal, "La fecha de fin no coincide")
	assert.Equal(t, 1, reservaDto.MesFinal, "La fecha de fin no coincide")
	assert.Equal(t, 13, reservaDto.DiaFinal, "La fecha de fin no coincide")
	assert.Equal(t, 3, reservaDto.Dias, "La cantidad de dias no coincide")
}