import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import { Link } from 'react-router-dom';
import './estilo/inicio.css';

const HomePage = () => {
  const [hotels, setHotels] = useState([]);
  const { isLogged } = useContext(AuthContext);

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
    if (!isLogged) {
      window.location.href = '/login-cliente';
    }
    else
    {
      window.location.href = `/reservar/${hotelId}`;
    }
  };

  return (
    <body className= "bodyinicio">
    <div className="containerIni">
      <div className="admin-button-container">
        <Link to="/admin" className="admin-button">
          Admin
        </Link>
      </div>
      <div className="hotels-container">
        {hotels.length ? (
          hotels.map((hotel) => (
            <div className='hotel-card' key={hotel.id}>
              <img src={hotel.image} alt={hotel.nombre}></img>
              <h4>{hotel.nombre}</h4>
              <p>{hotel.email}</p>
                <button onClick={() => Verificacion(hotel.id)}>
                  Reservar
                </button>
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