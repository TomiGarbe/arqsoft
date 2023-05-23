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
	InsertCliente(ClienteDto dto.ClienteDto) (dto.ClienteDto, e.ApiError)
	GetHoteles() (dto.HotelesDto, e.ApiError)
	InsertReserva(ReservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservas(dto.ReservasDto, e.ApiError)
	GetReservaById(id int) (dto.ReservaDto, e.ApiError)
}

var (
	clienteService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDetailDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)
	var userDetailDto dto.UserDetailDto

	if user.Id == 0 {
		return userDetailDto, e.NewBadRequestApiError("user not found")
	}

	userDetailDto.Name = user.Name
	userDetailDto.LastName = user.LastName
	userDetailDto.Street = user.Address.Street
	userDetailDto.Number = user.Address.Number
	for _, telephone := range user.Telephones {
		var dtoTelephone dto.TelephoneDto

		dtoTelephone.Code = telephone.Code
		dtoTelephone.Number = telephone.Number

		userDetailDto.TelephonesDto = append(userDetailDto.TelephonesDto, dtoTelephone)
	}

	return userDetailDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.Name
		userDto.Id = user.Id

		userDto.Street = user.Address.Street
		userDto.Number = user.Address.Number

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	var address model.Address

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.UserName = userDto.UserName
	user.Password = userDto.Password

	address.Number = userDto.Number
	address.Street = userDto.Street
	address = addressCliente.InsertAddress(address)

	user.Address = address
	user = userCliente.InsertUser(user)

	userDto.Id = user.Id

	return userDto, nil
}

func (s *userService) AddUserTelephone(telephoneDto dto.TelephoneDto) (dto.UserDetailDto, e.ApiError) {

	var telephone model.Telephone

	telephone.Code = telephoneDto.Code
	telephone.Number = telephoneDto.Number
	telephone.UserId = telephoneDto.UserId
	//Adding
	telephone = telephoneCliente.AddTelephone(telephone)

	// Find User
	var user model.User = userCliente.GetUserById(telephoneDto.UserId)
	var userDetailDto dto.UserDetailDto

	userDetailDto.Name = user.Name
	userDetailDto.LastName = user.LastName
	userDetailDto.Street = user.Address.Street
	userDetailDto.Number = user.Address.Number
	for _, telephone := range user.Telephones {
		var dtoTelephone dto.TelephoneDto

		dtoTelephone.Code = telephone.Code
		dtoTelephone.Number = telephone.Number

		userDetailDto.TelephonesDto = append(userDetailDto.TelephonesDto, dtoTelephone)
	}

	return userDetailDto, nil
}
