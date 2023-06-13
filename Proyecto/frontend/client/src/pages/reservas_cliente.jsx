import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/reservas_cliente.css';

const HomePage = () => {
  const [reservations, setReservations] = useState([]);
  const { isLoggedCliente } = useContext(AuthContext);
  const [accountId, setAccountId] = useState('');

  useEffect(() => {
    const getUser = () => {
      if (isLoggedCliente) {
        setAccountId(localStorage.getItem("id_cliente"))
      }
      else {
        window.location.href = '/';
      }
    };

    getUser();
  }, [isLoggedCliente]);

  const getReservations = async () => {
    try {
      const request = await fetch(`http://localhost:8090/cliente/reservas/${accountId}`);
      const response = await request.json();
      setReservations(response);
    } catch (error) {
      console.log("No se pudieron obtener las reservas:", error);
    }
  };

  return (
      <div className="reservations-container1">
       
        <div className="reservations-container2" onLoad={getReservations}>
          <h4>Datos de tus reservas:</h4>
          {reservations.length ? (
            reservations.map((reservation) => (
              <div className="reservation-card" key={reservation.id}>
                <p>Hotel: {reservation.hotel}</p>
                <p>Fecha de llegada: {reservation.fechaInicio}</p>
                <p>Fecha de fin: {reservation.fechaFin}</p>
                <p>Gracias por elegirnos!</p>
              </div>
            ))
          ) : (
            <p>No tienes reservas</p>
          )}
          </div>
        </div>
  );
};

export default HomePage;
