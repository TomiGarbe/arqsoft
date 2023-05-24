package services

import (
	adminClient "Proyecto/clients/admin"
	telefonoClient "Proyecto/clients/telefono"
	clienteClient "Proyecto/clients/cliente"
	hotelClient "Proyecto/clients/hotel"
	reservaClient "Proyecto/clients/reserva"

	"Proyecto/dto"
	"Proyecto/model"
	e "Proyecto/utils/errors"
)

type clienteService struct{}

type clienteServiceInterface interface {
	GetClienteById(id int) (dto.ClienteDto, e.ApiError)
	InsertCliente(clienteDto dto.ClienteDto) (dto.ClienteDto, e.ApiError)
	GetHoteles() (dto.HotelesDto, e.ApiError)
	InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservas() (dto.ReservasDto, e.ApiError)
	GetReservaById(id int) (dto.ReservaDto, e.ApiError)
}

var (
	ClienteService clienteServiceInterface
)

func init() {
	ClienteService = &clienteService{}
}

func (s *clienteService) GetClienteById(id int) (dto.ClienteDto, e.ApiError) {

	var cliente model.Cliente = clienteClient.GetClienteById(id)
	var clienteDto dto.ClienteDto

	if cliente.ID == 0 {
		return clienteDto, e.NewBadRequestApiError("user not found")
	}

	clienteDto.Name = cliente.Name
	clienteDto.LastName = cliente.LastName
	clienteDto.UserName = cliente.UserName
	clienteDto.Password = cliente.Password
	clienteDto.Email = cliente.Email

	return clienteDto, nil
}

func (s *clienteService) InsertCliente(clienteDto dto.ClienteDto) (dto.ClienteDto, e.ApiError) {

	var cliente model.Cliente

	cliente.Name = clienteDto.Name
	cliente.LastName = clienteDto.LastName
	cliente.UserName = clienteDto.UserName
	cliente.Password = clienteDto.Password
	cliente.Email = clienteDto.Email

	cliente = clienteClient.InsertCliente(cliente)

	clienteDto.ID = cliente.ID

	return clienteDto, nil
}

func (s *clienteService) GetHoteles() (dto.HotelesDto, e.ApiError) {

	var hoteles model.Hoteles = hotelClient.GetHoteles()
	var hotelesDto dto.HotelesDto

	for _, hotel := range hoteles {
		var hotelDto dto.HotelDto
		hotelDto.ID = hotel.ID
		hotelDto.Nombre = hotel.Nombre
		hotelDto.Email = hotel.Email
		hotelDto.Image = hotel.Image
		hotelDto.Cant_Hab = hotel.Cant_Hab

		for _, telefono := range hotel.Telefono {
			var dtoTelefono dto.TelefonoDto
	
			dtoTelefono.Code = telefono.Code
			dtoTelefono.Number = telefono.Number
	
			hotelDto.TelefonoDto = append(hotelDto.TelefonoDto, dtoTelefono)
		}

		hotelesDto = append(hotelesDto, hotelDto)
	}

	return hotelesDto, nil
}

func (s *clienteService) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva
	var hotel model.Hotel
	var cliente model.Cliente

	reserva.FechaInicio = reservaDto.FechaInicio
	reserva.FechaFinal = reservaDto.FechaFinal
	reserva.Dias = reservaDto.Dias
	reserva.Disponibilidad = reservaDto.Disponibilidad

	hotel.Nombre = reservaDto.Nombre
	cliente.Name = reservaDto.Name
	cliente.LastName = reservaDto.LastName
	hotel = hotelClient.InsertHotel(hotel)
	cliente = clienteClient.InsertCliente(cliente)

	reserva.Hotel = hotel
	reserva.Cliente = cliente

	reserva = clienteClient.InsertCliente(reserva)

	reservaDto.ID = reserva.ID

	return reservaDto, nil
}

func (s *clienteService) GetReservas() (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.GetReservas()
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto
		reservaDto.ID = reserva.ID
		for _, hotel := range reserva.Hotel {
			var dtoHotel dto.HotelDto
	
			dtoHotel.Nombre = hotel.Nombre
	
			reservaDto.HotelDto = append(reservaDto.HotelDto, dtoHotel)
		}
		for _, cliente := range reserva.Cliente {
			var dtoCliente dto.ClienteDto
	
			dtoCliente.Name = cliente.Name
			dtoCliente.LastName = cliente.LastName
	
			reservaDto.ClienteDto = append(reservaDto.ClienteDto, dtoCliente)
		}
		reservaDto.FechaInicio = reserva.FechaInicio
		reservaDto.FechaFinal = reserva.FechaFinal
		reservaDto.Dias = reserva.Dias
		reservaDto.Disponibilidad = reserva.Disponibilidad

		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}