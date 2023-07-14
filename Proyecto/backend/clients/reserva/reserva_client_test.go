package clients_test

import (
	"testing"

	"backend/model"

	"github.com/stretchr/testify/assert"
)

type MockReservaClient struct{}

func (m *MockReservaClient) InsertReserva(reserva model.Reserva) model.Reserva {
	// Simular la lógica de inserción en la base de datos
	// Se establece un ID para la reserva
	reserva.ID = 1 // Si modifico a cero, genera la alerta
	return reserva
}

// TEST PARA LA FUNCION GETBOOKINGBYIDyy

func TestInsertBooking(t *testing.T) {
	// Crear una instancia del mock del DAO de Booking
	mockClient := &MockReservaClient{}

	// Crear una nueva reserva
	newReserva := model.Reserva{
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

	// Insertar la reserva utilizando el mock del DAO
	inserted := mockClient.InsertReserva(newReserva)

	// Verificar que la reserva tenga un ID asignado
	assert.NotZero(t, inserted.ID, "La reserva no se pudo realizar")

	// Verificar otros atributos de la reserva
	assert.Equal(t, newReserva.ClienteID, inserted.ClienteID)
	assert.Equal(t, newReserva.HotelID, inserted.HotelID)
	assert.Equal(t, newReserva.AnioInicio, inserted.AnioInicio)
	assert.Equal(t, newReserva.MesInicio, inserted.MesInicio)
	assert.Equal(t, newReserva.DiaInicio, inserted.DiaInicio)
	assert.Equal(t, newReserva.AnioFinal, inserted.AnioFinal)
	assert.Equal(t, newReserva.MesFinal, inserted.MesFinal)
	assert.Equal(t, newReserva.DiaFinal, inserted.DiaFinal)
	assert.Equal(t, newReserva.Dias, inserted.Dias)
}

// TEST PARA LA FUNCION GETBOOKINGBYID

func (m *MockReservaClient) GetReservaById(id int) model.Reserva {
	// Simular la búsqueda en la base de datos
	reserva := model.Reserva{
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

	return reserva
}

func TestGetReservaById(t *testing.T) {
	// Crear una instancia del mock del DAO de Booking
	mockClient := &MockReservaClient{}

	// ID de reserva a buscar - Si la cambio deja de funcionar
	reservaId := 1

	// Obtener la reserva utilizando el mock del DAO
	reserva := mockClient.GetReservaById(reservaId)

	// Verificar que la reserva obtenida tenga el ID correcto
	assert.Equal(t, reservaId, reserva.ID, "El ID de la reserva no existe")

	// Verificar otros atributos de la reserva
	assert.Equal(t, 1, reserva.ClienteID)
	assert.Equal(t, 1, reserva.HotelID)
	assert.Equal(t, 2023, reserva.AnioInicio)
	assert.Equal(t, 1, reserva.MesInicio)
	assert.Equal(t, 10, reserva.DiaInicio)
	assert.Equal(t, 2023, reserva.AnioFinal)
	assert.Equal(t, 1, reserva.MesFinal)
	assert.Equal(t, 13, reserva.DiaFinal)
	assert.Equal(t, 3, reserva.Dias)
}