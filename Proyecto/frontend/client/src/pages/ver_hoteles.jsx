import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/ver_hoteles.css';

const HomePage = () => {
  const [hotels, setHotels] = useState([]);
  const { isLoggedAdmin } = useContext(AuthContext);

  const getHotels = async () => {
    try {
      const request = await fetch("http://localhost:8090/admin/hoteles");
      const response = await request.json();
      setHotels(response);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  };

  useEffect(() => {
    getHotels();
  }, []);

  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  return (
    <body className="bodyinicio" onLoad={Verificacion}>
      <div className="containerIni">
        <div className="hotels-container">
          {hotels.length ? (
            hotels.map((hotel) => (
              <div className="hotel-card" key={hotel.id}>
                <img src={hotel.image} alt={hotel.nombre} className="hotel-image" />
                <div className="hotel-info">
                  <h4>{hotel.nombre}</h4>
                  <p>{hotel.email}</p>
                </div>
                <div className="hotel-description">
                    <label htmlFor={`description-${hotel.id}`}>Descripción:</label>
                    <p id={`description-${hotel.id}`}>{hotel.descripcion}</p>
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