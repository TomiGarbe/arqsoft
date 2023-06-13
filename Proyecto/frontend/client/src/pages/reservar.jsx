import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import { useParams } from 'react-router-dom';
import './estilo/reservar.css';

const ReservaPage = () => {
  const { hotelId } = useParams();
  const [hotelData, setHotelData] = useState('');
  const [clientData, setClienteData] = useState('');
  const [cantidadPersonas, setCantidadPersonas] = useState('');
  const [commodities, setCommodities] = useState('');
  const { isLoggedCliente } = useContext(AuthContext);
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');

  const [formData, setFormData] = useState({
    nombre: hotelData.nombre,
    name: clientData.name,
    last_name: clientData.last_name,
    fecha_inicio: startDate,
    fecha_final: endDate
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const Verificacion = (hotelId) => {
    if (!isLoggedCliente) {
      window.location.href = '/login-cliente';
    }
    else
    {
      const accountId = localStorage.getItem("id_cliente");
      fetch(`http://localhost:8090/cliente/${accountId}`)
      .then(response => response.json())
      .then(data => {
        setClienteData(data);
      })
      .catch(error => {
        console.error('Error al obtener los datos del cliente:', error);
      });
    }
  };

  const handleReserva = () => {
    fetch('http://localhost:8090/cliente/reserva', {
      method: 'POST',
      headers: {
      'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData)
      })
      .then(response => response.json())
      .then(data => {
        console.log('Registro exitoso:', data);
        window.location.href = '/login-cliente';
      })
      .catch(error => {
        console.error('Error en el registro:', error);
        alert('Credenciales incorrectas');
      });
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
        console.error('Error al obtener los datos del cliente:', error);
      });
    }
  }, [hotelId]);

  const handleStartDateChange = (event) => {
    setStartDate(event.target.value);
    handleChange();
    const selectedStartDateObj = new Date(event.target.value);
    const endDateObj = new Date(endDate);
    if (selectedStartDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no validas");
    }
    if (startDate && endDate) {
      filterHotels();
    }
  };

  const handleEndDateChange = (event) => {
    setEndDate(event.target.value);
    handleChange();
    const selectedStartDateObj = new Date(startDate);
    const endDateObj = new Date(event.target.value);
    if (selectedStartDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no validas");
    }
    if (startDate && endDate) {
      filterHotels();
    }
  };

  const handleCantPersonasChange = (event) => {
    setCantidadPersonas(event.target.value)
    handleChange();
  };

  const filterHotels = async () => {
    const request = await fetch(`http://localhost:8090/cliente/disponibilidad/${hotelId}/${startDate}/${endDate}`);
    const response = await request.json();
    if (response === 0) {
      setEndDate('');
      alert("No hay habitaciones disponibles para esas fechas");
    }
    else if (response === cantidadPersonas) {
      setCantidadPersonas('');
      alert("No hay habitaciones disponibles para esa cantidad de personas");
    }
  };

  return (
    <div>
      {typeof hotelData === 'undefined' ? (
        <>CARGANDO...</>
      ) : (
        <div className="container" onLoad={Verificacion}>
          <div className="reserva-form">
            <h2>Reserva de Hotel</h2>
            <h3>{hotelData["nombre"]}</h3>
            <img src="ruta/de/la/foto.jpg" alt="Foto del Hotel" />
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
              <div className="form-group">
                <label htmlFor="cantidadPersonas">Cantidad de personas:</label>
                <input
                  type="number"
                  id="cantidadPersonas"
                  value={cantidadPersonas}
                  onChange={handleCantPersonasChange}
                  required
                />
              </div>
              <div className="form-group">
                <label htmlFor="commodities">Commodities:</label>
                <textarea
                  id="commodities"
                  value={commodities}
                  onChange={(e) => setCommodities(e.target.value)}
                  required
                ></textarea>
              </div>
              <button type="submit">Confirmar</button>
            </form>
          </div>
        </div>
      )}
    </div>
  );
};

export default ReservaPage;
