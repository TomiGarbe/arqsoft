import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/inicio.css';

const HomePage = () => {
  const [hotels, setHotels] = useState([]);
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const { isLoggedCliente } = useContext(AuthContext);
  const { isLoggedAdmin } = useContext(AuthContext);

  const getHotels = async () => {
    try {
      const request = await fetch("http://localhost:8090/cliente/hoteles");
      const response = await request.json();
      setHotels(response);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  };

  useEffect(() => {
    getHotels();
  }, []);

  const Verificacion = (hotelId) => {
    if (!isLoggedCliente) {
      window.location.href = '/login-cliente';
    }
    else
    {
      window.location.href = `/reservar/${hotelId}`;
    }
  };

  const handleStartDateChange = (event) => {
    setStartDate(event.target.value);
    const selectedStartDateObj = new Date(event.target.value);
    const endDateObj = new Date(endDate);
    if (selectedStartDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no validas");
    }
  };

  const handleEndDateChange = (event) => {
    setEndDate(event.target.value);
    const selectedStartDateObj = new Date(startDate);
    const endDateObj = new Date(event.target.value);
    if (selectedStartDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no validas");
    }
  };

  const filterHotels = async () => {
    if (startDate === '' || endDate === '') {
      window.location.href = '/';
    }
    else {
      getHotels();
      const startDateObj = new Date(startDate);
      const endDateObj = new Date(endDate);
      for (let i = 0; i < hotels.length; i++) {
        const request = await fetch(`http://localhost:8090/cliente/disponibilidad/${hotels[i].id}/${startDateObj.getFullYear()}/${startDateObj.getMonth() + 1}/${startDateObj.getDate() + 1}/${endDateObj.getFullYear()}/${endDateObj.getMonth() + 1}/${endDateObj.getDate() + 1}`);
        const response = await request.json();
        if (response === 0) {
          setHotels((prevHotels) => prevHotels.filter((hotel) => hotel.id !== hotels[i].id));
        }
      }
    }
  };

  const Admin = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
    else
    {
      window.location.href = '/admin';
    }
  }

  const Cuenta = () => {
    window.location.href = '/cuenta';
  }

  return (
    <body className= "bodyinicio">
      <div className="header-content">
        <div className="cuenta-button-container">
          <button className="cuenta-button" onClick={Cuenta}>
            Tu Cuenta
          </button>
        </div>
        <div className="admin-button-container">
          <button className="admin-button" onClick={Admin}>
            Admin
          </button>
        </div>
      </div>
        
        <div className="contdeFechas">
          <div className="date-pickerINI1">
            <label htmlFor="start-date" className="fecha">Entrada</label>
            <input type="date" id="start-date" value={startDate} onChange={handleStartDateChange} />
          </div>
          <div className="date-pickerINI1">
            <label htmlFor="end-date" className="fecha">Salida</label>
            <input type="date" id="end-date" value={endDate} onChange={handleEndDateChange} />
          </div>
            <button className="botbusquedaFec" onClick={filterHotels}>Buscar</button>
            </div>
      <div className="containerIni">
      <div className="hotels-container">
            {hotels.length ? 
              ( hotels.map((hotel) => (
                <div className='hotel-card' key={hotel.id}>
                  <img src={hotel.image} alt={hotel.nombre} className="hotel-image" />
                  <div className="hotel-info">
                    <h4>{hotel.nombre}</h4>
                    <p>{hotel.email} </p>
                    <button onClick={() => Verificacion(hotel.id)}>
                      Reservar
                    </button>
                  </div>
                </div>
              ))
              ) : (
                <p>No hay hoteles</p>
            )}
          </div>
      </div>
    </body> 
  );
};

export default HomePage;
