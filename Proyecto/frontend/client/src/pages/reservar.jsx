import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import { useParams } from 'react-router-dom';
import './estilo/reservar.css';


const ReservaPage = () => {
  const { hotelId } = useParams();
  const [hotelData, setHotelData] = useState('');
  const [imagenHotel, setImagenHotel] = useState([]);
  const { isLoggedCliente } = useContext(AuthContext);
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const accountId = localStorage.getItem("id_cliente");
  const [Hoteles, setHoteles] = useState([]);
  const [Imagenes, setImagenes] = useState([]);
  const [confirmarReserva, setConfirmarReserva] = useState(false);

  const Verificacion = () => {
    if (!isLoggedCliente) {
      window.location.href = '/login-cliente';
    }
  };

  const handleReserva = () => {
    if (confirmarReserva) {
      const startDateObj = new Date(startDate);
      const endDateObj = new Date(endDate);
      const Dias = Math.round((endDateObj - startDateObj) / (1000 * 60 * 60 * 24));
      const formData = {
        hotel_id: parseInt(hotelId),
        cliente_id: parseInt(accountId),
        anio_inicio: startDateObj.getFullYear(),
        anio_final: endDateObj.getFullYear(),
        mes_inicio: startDateObj.getMonth() + 1, 
        mes_final: endDateObj.getMonth() + 1, 
        dia_inicio: startDateObj.getDate() + 1,
        dia_final: endDateObj.getDate() + 1,
        dias: Dias
      };

      fetch('http://localhost:8090/cliente/reserva', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
      })
      .then(response => response.json())
      .then(data => {
        console.log('Reserva exitosa:', data);
        alert('Reserva exitosa');
        handleVolver();
      })
      .catch(error => {
        console.error('Error en el registro:', error);
        alert('Error al reservar');
      });
    }
    else {
      alert("No hay habitaciones disponibles para esas fechas");
    }
  };

  useEffect(() => {
    setHotelData('');
    if (hotelId) {
      fetch(`http://localhost:8090/cliente/hotel/${hotelId}`)
        .then(response => response.json())
        .then(data => {
          setHotelData(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del hotel:', error);
        });
    }
  }, [hotelId]);

  const getImagen = useCallback(async () => {
    if (hotelId) {
      try {
        let imagen = [];
        const request = await fetch(`http://localhost:8090/cliente/imagenes/hotel/${hotelId}`);
        const response = await request.json();
        for (let i = 0; i < response.length; i++) {
          imagen.push(response[i].url);
        }
        setImagenHotel(imagen);
      } catch (error) {
        console.log("No se pudieron obtener los hoteles:", error);
      }
    }
  }, [hotelId]);

  useEffect(() => {
    getImagen();
  }, [getImagen]);

  const getHoteles = useCallback(async () => {
    try {
      let hotelesArray = [];
      const request = await fetch('http://localhost:8090/cliente/hoteles');
      const response = await request.json();
      hotelesArray = response.filter((hotel) => hotel.id !== parseInt(hotelId));
      setHoteles(hotelesArray);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  }, [hotelId]);

  const getImagenes = useCallback(async () => {
    try {
      let imagenesHotelesArray = [];
      for (let i = 0; i < Hoteles.length; i++) {
        const Hotel = Hoteles[i];
        const request = await fetch(`http://localhost:8090/cliente/imagenes/hotel/${Hotel.id}`);
        const response = await request.json();
        if (response.length > 0) {
          imagenesHotelesArray.push({url: response[0].url, hotel_id: response[0].hotel_id});
        }
      }
      imagenesHotelesArray = imagenesHotelesArray.filter((imagen) => imagen.hotel_id !== parseInt(hotelId));
      setImagenes(imagenesHotelesArray);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  }, [hotelId, Hoteles]);

  useEffect(() => {
    getImagenes();
  }, [getImagenes]);

  useEffect(() => {
    getHoteles();
  }, [getHoteles]);

  const filterHotel = useCallback(async () => {
    const startDateObj = new Date(startDate);
    const endDateObj = new Date(endDate);
    const request = await fetch(`http://localhost:8090/cliente/disponibilidad/${hotelId}/${startDateObj.getFullYear()}/${startDateObj.getMonth() + 1}/${startDateObj.getDate() + 1}/${endDateObj.getFullYear()}/${endDateObj.getMonth() + 1}/${endDateObj.getDate() + 1}`);
    const response = await request.json();
    if (response === 0) {
      setConfirmarReserva(false);
    }
    else {
      setConfirmarReserva(true);
    }
    for (let i = 0; i < Hoteles.length; i++) {
      const request = await fetch(`http://localhost:8090/cliente/disponibilidad/${Hoteles[i].id}/${startDateObj.getFullYear()}/${startDateObj.getMonth() + 1}/${startDateObj.getDate() + 1}/${endDateObj.getFullYear()}/${endDateObj.getMonth() + 1}/${endDateObj.getDate() + 1}`);
      const response = await request.json();
      if (response === 0) {
        setHoteles((prevHotels) => prevHotels.filter((hotel) => hotel.id !== Hoteles[i].id));
      }
    }
  }, [startDate, endDate, hotelId, Hoteles]);

  const handleStartDateChange = (event) => {
    setStartDate(event.target.value);
    const startDateObj = new Date(event.target.value);
    const endDateObj = new Date(endDate);
    if (startDateObj >= endDateObj) {
      alert("Fechas no válidas");
      setEndDate('');
    } else {
      filterHotel();
    }
  };

  const handleEndDateChange = (event) => {
    setEndDate(event.target.value);
    const startDateObj = new Date(startDate);
    const endDateObj = new Date(event.target.value);
    if (startDateObj >= endDateObj) {
      alert("Fechas no válidas");
      setEndDate('');
    } else {
      filterHotel();
    }
  };

  const handleVolver = () => {
    window.location.href = 'http://localhost:3000/';
  };

  const Reservar = (hotelId) => {
    window.location.href = `/reservar/${hotelId}`;
  };

  return (
    <div className="bodyReserva">
      <div>
        {typeof hotelData === 'undefined' ? (
          <>CARGANDO...</>
        ) : (
          <div className="container45" onLoad={Verificacion}>
            <div className="informacion">
              <div className="cuadroImag">
                {imagenHotel.map((imagen, index) => (
                  <img
                    key={index}
                    src={`http://localhost:8090/${imagen}`}
                    alt={hotelData.nombre}
                    className="tamanoImag"
                  />
                ))}
              </div>
              <div className="descripcion">{hotelData["descripcion"]}</div>
              <div className="amenities">
                <h6>Amenities:</h6>
                {Array.isArray(hotelData.amenities) && hotelData.amenities.length > 0 ? (
                  hotelData.amenities.map((amenity, index) => (
                    <span key={index} className="amenity">{amenity}</span>
                  ))
                ) : (
                  <span>No hay amenities disponibles</span>
                )}
              </div>
              <div className='other-hotels-title'><h6>Otras opciones:</h6></div>
              <div className="other-hotels">
                {Hoteles.map((hotel) => {
                  const imagen = Imagenes.find((imagen) => imagen.hotel_id === hotel.id);
                  return(
                  <div key={hotel.id} className="other-hotels">
                    {imagen ? (
                      <img src={`http://localhost:8090/${imagen.url}`} alt={hotel.nombre} className="other-hotel-image" />
                    ) : (
                      <div className="hotel-image-placeholder" />
                    )}
                    <div className="hotel-details">
                      <h6>{hotel.nombre}</h6>
                      <button className="reservar-button" onClick={() => Reservar(hotel.id)}>
                        Reservar
                      </button>
                    </div>
                  </div>
                )})}
              </div>
            </div>
            <div className="reserva-form">
              <h6>Realice reserva del Hotel</h6>
              <h6>{hotelData["nombre"]}</h6>
              <form onSubmit={handleReserva}>
                <div className="form-group">
                  <label htmlFor="fechaInicio">Fecha de inicio:</label>
                  <input
                    type="date"
                    id="fechaInicio"
                    value={startDate}
                    onChange={handleStartDateChange}
                    required
                  />
                </div>
                <div className="form-group">
                  <label htmlFor="fechaFin">Fecha de fin:</label>
                  <input
                    type="date"
                    id="fechaFin"
                    value={endDate}
                    onChange={handleEndDateChange}
                    required
                  />
                </div>
                <div>
                  <button type="submit" className="confReserva">Confirmar</button>
                  <button type="button" className="confReserva" onClick={handleVolver}>Volver</button>
                </div>
              </form>
            </div>
          </div>
        )}
      </div>
    </div>
  );

};

export default ReservaPage;
