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

type adminService struct{}

type adminServiceInterface interface {
	GetAdminById(id int) (dto.AdminDto, e.ApiError)
	GetAdmins() (dto.AdminsDto, e.ApiError)
	InsertAdmin(adminDto dto.AdminDto) (dto.AdminDto, e.ApiError)
	GetClienteById(id int) (dto.ClienteDto, e.ApiError)
	GetClientes() (dto.ClientesDto, e.ApiError)
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	GetHoteles() (dto.HotelesDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	GetReservas() (dto.ReservasDto, e.ApiError)
	AddTelefono(telefonoDto dto.TelefonoDto) (dto.HotelDto, e.ApiError)
}

var (
	AdminService adminServiceInterface
)

func init() {
	AdminService = &adminService{}
}

func (s *adminService) GetAdminById(id int) (dto.AdminDto, e.ApiError) {

	var admin model.Admin = adminClient.GetAdminById(id)
	var adminDto dto.AdminDto

	if admin.ID == 0 {
		return adminDto, e.NewBadRequestApiError("admin not found")
	}

	adminDto.Name = admin.Name
	adminDto.LastName = admin.LastName
	adminDto.UserName = admin.UserName
	adminDto.Password = admin.Password
	adminDto.Email = admin.Email

	return adminDto, nil
}

func (s *adminService) GetAdmins() (dto.AdminsDto, e.ApiError) {

	var admins model.Admins = adminClient.GetAdmins()
	var adminsDto dto.AdminsDto

	for _, admin := range admins {
		var adminDto dto.AdminDto
		adminDto.ID = admin.ID
		adminDto.Name = admin.Name
		adminDto.LastName = admin.LastName
		adminDto.UserName = admin.UserName
		adminDto.Password = admin.Password
		adminDto.Email = admin.Email

		adminsDto = append(adminsDto, adminDto)
	}

	return adminsDto, nil
}

func (s *adminService) InsertAdmin(adminDto dto.AdminDto) (dto.AdminDto, e.ApiError) {

	var admin model.Admin

	admin.Name = adminDto.Name
	admin.LastName = adminDto.LastName
	admin.UserName = adminDto.UserName
	admin.Password = adminDto.Password
	admin.Email = adminDto.Email

	admin = adminClient.InsertAdmin(admin)

	adminDto.ID = admin.ID

	return adminDto, nil
}

func (s *adminService) GetClienteById(id int) (dto.ClienteDto, e.ApiError) {

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

func (s *adminService) GetClientes() (dto.ClientesDto, e.ApiError) {

	var clientes model.Clientes = clienteClient.GetClientes()
	var clientesDto dto.ClientesDto

	for _, cliente := range clientes {
		var clienteDto dto.ClienteDto
		clienteDto.ID = cliente.ID
		clienteDto.Name = cliente.Name
		clienteDto.LastName = cliente.LastName
		clienteDto.UserName = cliente.UserName
		clienteDto.Password = cliente.Password
		clienteDto.Email = cliente.Email

		clientesDto = append(clientesDto, clienteDto)
	}

	return clientesDto, nil
}

func (s *adminService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.ID == 0 {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

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

	return hotelDto, nil
}

func (s *adminService) GetHoteles() (dto.HotelesDto, e.ApiError) {

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

func (s *adminService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Nombre = hotelDto.Nombre
	hotel.Email = hotelDto.Email
	hotel.Image = hotelDto.Image
	hotel.Cant_Hab = hotelDto.Cant_Hab

	hotel = hotelClient.InsertHotel(hotel)

	hotelDto.ID = hotel.ID

	return hotelDto, nil
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
