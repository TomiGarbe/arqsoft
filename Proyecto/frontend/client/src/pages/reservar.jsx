import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import { useParams } from 'react-router-dom';
import './estilo/reservar.css';

const ReservaPage = () => {
  const { hotelId } = useParams();
  const [hotelData, setHotelData] = useState('');
  const { isLoggedCliente } = useContext(AuthContext);
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const accountId = localStorage.getItem("id_cliente");

  const Verificacion = () => {
    if (!isLoggedCliente) {
      window.location.href = '/login-cliente';
    }
  };

  const handleReserva = () => {
    const startDateObj = new Date(startDate);
    const endDateObj = new Date(endDate);
    const [formData] = {
      hotel_id: hotelId,
      cliente_id: accountId,
      fecha_inicio: startDate,
      fecha_final: endDate,
      dias: endDateObj-startDateObj
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
        console.log('Registro exitoso:', data);
        window.location.href = '/';
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
    const startDateObj = new Date(event.target.value);
    const endDateObj = new Date(endDate);
    if (startDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no validas");
    }
    if (startDate && endDate) {
      filterHotels();
    }
  };

  const handleEndDateChange = (event) => {
    setEndDate(event.target.value);
    const startDateObj = new Date(startDate);
    const endDateObj = new Date(event.target.value);
    if (startDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no validas");
    }
    if (startDate && endDate) {
      filterHotels();
    }
  };

  const filterHotels = async () => {
    const request = await fetch(`http://localhost:8090/cliente/disponibilidad/${hotelId}/${startDate}/${endDate}`);
    const response = await request.json();
    if (response === 0) {
      setEndDate('');
      alert("No hay habitaciones disponibles para esas fechas");
    }
  };

  return (
    <body className="bodyReserva">
    <div>
      {typeof hotelData === 'undefined' ? (
        <>CARGANDO...</>
      ) : (
        
        <div className="container45" onLoad={Verificacion}>
          <div className= "descripcion">{hotelData["nombre"]}</div>
          <div className="reserva-form">
            <h6>Realice reserva del Hotel</h6>
            <h6>{hotelData["nombre"]}</h6>
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
              <button type="submit" className="confReserva">Confirmar</button>
            </form>
          </div>
        </div>
        
      )}
    </div>
    </body>
  );
};

export default ReservaPage;
