import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/ver_reservas.css';

const HomePage = () => {
  const [reservas, setReservas] = useState([]);
  const { isLoggedAdmin } = useContext(AuthContext);

  const getReservas = async () => {
    try {
      const request = await fetch("http://localhost:8090/admin/reservas");
      const response = await request.json();
      setReservas(response);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  };

  useEffect(() => {
    getReservas();
  }, []);

  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  return (
    <body className="bodyinicio" onLoad={Verificacion}>
      <div className="containerIni">
        <div className="reserva-container">
          {reservas.length ? (
            reservas.map((reserva) => (
              <div className="reserva-card" key={reserva.id}>
                <img src={reserva.nombre} alt={reserva.nombre} className="hotel-image" />
                <div className="reserva-info">
                  <h4>{reserva.nombre}</h4>
                  <p>{reserva.email}</p>
                </div>
                <div className="cliente-info">
                  <p>{reserva.name}</p>
                  <p>{reserva.last_name}</p>
                </div>
                <div className="fechas-info">
                  <p>{reserva.fecha_inicio}</p>
                  <p>{reserva.fecha_final}</p>
                  <p>{reserva.dias}</p>
                </div>
              </div>
            ))
          ) : (
            <p>No hay reservas</p>
          )}
        </div>
      </div>
    </body>
  );
};

export default HomePage;