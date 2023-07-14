package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	// Configurar el assert
	assert := assert.New(t)

	// Crear valores de prueba para Cliente
	cliente := Cliente{
		ID:       1,
		Name:     "Facu",
		LastName: "Gazzera",
		UserName: "facuga",
		Password: "1234",
		Email:    "facugazzera@gmail.com",
	}

	// Crear valores de prueba para Hotel
	hotel := Hotel{
		ID:           	1,
		Nombre:         "Hotel 1",
		Descripcion:    "Este es un hotel de ejemplo",
		Email:        	"hotel1@gmail.com",
		Cant_Hab:       10,
	}

	// Crear una instancia de Reserva con valores de prueba
	reserva := Reserva{
		ID:       1,
		Cliente:     cliente,
		Hotel:    hotel,

		AnioInicio: 2023,
		AnioFinal:	2023,
		MesInicio:	1,
		MesFinal:	1,
		DiaInicio:	10,
		DiaFinal:	13,
		Dias:		3,
	}

	expectedAnioInicio := 2023
	assert.Equal(expectedAnioInicio, reserva.AnioInicio, "Se espero que el anio de inicio sea %v", expectedAnioInicio)

	expectedMesInicio := 1
	assert.Equal(expectedMesInicio, reserva.MesInicio, "Se espero que el anio de inicio sea %v", expectedMesInicio)

	expectedDiaInicio := 10
	assert.Equal(expectedDiaInicio, reserva.DiaInicio, "Se espero que el anio de inicio sea %v", expectedDiaInicio)

	expectedAnioFinal := 2023
	assert.Equal(expectedAnioFinal, reserva.AnioFinal, "Se espero que el anio de inicio sea %v", expectedAnioFinal)

	expectedMesFinal := 1
	assert.Equal(expectedMesFinal, reserva.MesFinal, "Se espero que el anio de inicio sea %v", expectedMesFinal)

	expectedDiaFinal := 13
	assert.Equal(expectedDiaFinal, reserva.DiaFinal, "Se espero que el anio de inicio sea %v", expectedDiaFinal)

	expectedDias := 3
	assert.Equal(expectedDias, reserva.Dias, "Se espero que la cantidad de dias sea %v", expectedDias)

	// Verificar que Id de Reserva
	assert.Equal(1, reserva.ID, "El ID de la reserva no coincide")

	// Verificar propiedades de Cliente
	assert.Equal("Facu", reserva.Cliente.Name, "El nombre no coincide")
	assert.Equal("facuga", reserva.Cliente.UserName, "El nombre de cliente no coincide")

	// Verificar propiedades de Hotel
	assert.Equal("Hotel 1", reserva.Hotel.Nombre, "El nombre del hotel coincide")
	assert.Equal(10, reserva.Hotel.Cant_Hab, "La cantidad de habitaciones no coincide")
}