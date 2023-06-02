package services

import (
	clienteClient "backend/clients/cliente"
	hotelClient "backend/clients/hotel"
	reservaClient "backend/clients/reserva"

	"backend/dto"
	"backend/model"
	e "backend/utils/errors"
)

type clienteService struct{}

type clienteServiceInterface interface {
	GetClienteById(id int) (dto.ClienteDto, e.ApiError)
	GetClienteByUsername(username string) (dto.ClienteDto, e.ApiError)
	GetClienteByPassword(username string) (dto.ClienteDto, e.ApiError)
	GetClienteByEmail(email string) (dto.ClienteDto, e.ApiError)
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
		return clienteDto, e.NewBadRequestApiError("cliente not found")
	}

	clienteDto.Name = cliente.Name
	clienteDto.LastName = cliente.LastName
	clienteDto.UserName = cliente.UserName
	clienteDto.Password = cliente.Password
	clienteDto.Email = cliente.Email

	return clienteDto, nil
}

func (s *clienteService) GetClienteByUsename(username string) (dto.ClienteDto, e.ApiError) {
	var cliente model.Cliente = clienteClient.GetClienteByUsename(username)
	var clienteDto dto.ClienteDto

	if cliente.UserName == "" {
		return clienteDto, e.NewBadRequestApiError("cliente not found")
	}

	clienteDto.ID = cliente.ID
	clienteDto.Name = cliente.Name
	clienteDto.LastName = cliente.LastName
	clienteDto.Password = cliente.Password
	clienteDto.Email = cliente.Email

	return clienteDto, nil
}

func (s *clienteService) GetClienteByPassword(password string) (dto.ClienteDto, e.ApiError) {
	var cliente model.Cliente = clienteClient.GetClienteByPassword(password)
	var clienteDto dto.ClienteDto

	if cliente.Password == "" {
		return clienteDto, e.NewBadRequestApiError("cliente not found")
	}

	clienteDto.ID = cliente.ID
	clienteDto.Name = cliente.Name
	clienteDto.LastName = cliente.LastName
	clienteDto.UserName = cliente.UserName
	clienteDto.Email = cliente.Email

	return clienteDto, nil
}

func (s *clienteService) GetClienteByEmail(email string) (dto.ClienteDto, e.ApiError) {
	var cliente model.Cliente = clienteClient.GetClienteByPassword(password)
	var clienteDto dto.ClienteDto

	if cliente.Email == "" {
		return clienteDto, e.NewBadRequestApiError("cliente not found")
	}

	clienteDto.ID = cliente.ID
	clienteDto.Name = cliente.Name
	clienteDto.LastName = cliente.LastName
	clienteDto.UserName = cliente.UserName
	clienteDto.Password = cliente.Password

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

		for _, telefono := range hotel.Telefonos {
			var dtoTelefono dto.TelefonoDto
	
			dtoTelefono.Codigo = telefono.Codigo
			dtoTelefono.Numero = telefono.Numero
	
			hotelDto.TelefonosDto = append(hotelDto.TelefonosDto, dtoTelefono)
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

	reserva = reservaClient.InsertReserva(reserva)

	reservaDto.ID = reserva.ID

	return reservaDto, nil
}

func (s *clienteService) GetReservas() (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.GetReservas()
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto
		reservaDto.ID = reserva.ID
		reservaDto.Nombre = reserva.Hotel.Nombre
		reservaDto.Name = reserva.Cliente.Name
		reservaDto.LastName = reserva.Cliente.LastName
		reservaDto.FechaInicio = reserva.FechaInicio
		reservaDto.FechaFinal = reserva.FechaFinal
		reservaDto.Dias = reserva.Dias
		reservaDto.Disponibilidad = reserva.Disponibilidad

		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *clienteService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva = reservaClient.GetReservaById(id)
	var reservaDto dto.ReservaDto

	if reserva.ID == 0 {
		return reservaDto, e.NewBadRequestApiError("reserva not found")
	}

	reservaDto.Nombre = reserva.Hotel.Nombre
	reservaDto.Name = reserva.Cliente.Name
	reservaDto.LastName = reserva.Cliente.LastName
	reservaDto.FechaInicio = reserva.FechaInicio
	reservaDto.FechaFinal = reserva.FechaFinal
	reservaDto.Dias = reserva.Dias
	reservaDto.Disponibilidad = reserva.Disponibilidad

	return reservaDto, nil
}