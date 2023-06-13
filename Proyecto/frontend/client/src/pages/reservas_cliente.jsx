import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/reservas_cliente.css';

const HomePage = () => {
  const [reservations, setReservations] = useState([]);
  const [hoteles, setHoteles] = useState([]);
  const { isLoggedCliente } = useContext(AuthContext);

  const getHoteles = useCallback(async () => {
    try {
      const hotelesArray = [];
      for (let i = 0; i < reservations.length; i++) {
        const reserva = reservations[i];
        const request = await fetch(`http://localhost:8090/cliente/hotel/${reserva.hotel_id}`);
        const response = await request.json();
        hotelesArray.push(response);
      }
      setHoteles(hotelesArray);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  }, [reservations]);

  const getReservations = useCallback(async () => {
    if (isLoggedCliente) {
      const accountId = localStorage.getItem("id_cliente");
      try {
        const request = await fetch(`http://localhost:8090/cliente/reservas/${accountId}`);
        const response = await request.json();
        setReservations(response);
      } catch (error) {
        console.log("No se pudieron obtener las reservas:", error);
      }
      finally {
        getHoteles();
      }
    } else {
      window.location.href = '/';
    }
  }, [isLoggedCliente, getHoteles]);

  useEffect(() => {
    getReservations();
  }, [getReservations]);

  return (
    <div className="reservations-container1">
      <div className="reservations-container2">
        <h4>Datos de tus reservas:</h4>
        {reservations.length ? (
          reservations.map((reservation) => {
            const hotel = hoteles.find((hotel) => hotel.id === reservation.hotel_id);
            return (
              <div className="reservation-card" key={reservation.ID}>
                <p>Hotel: {hotel ? hotel.nombre : 'Hotel desconocido'}</p>
                <p>Fecha de llegada: {reservation.fecha_inicio}</p>
                <p>Fecha de fin: {reservation.fecha_final}</p>
                <p>Gracias por elegirnos!</p>
              </div>
            );
          })
        ) : (
          <p>No tienes reservas</p>
        )}
      </div>
    </div>
  );
};

export default HomePage;
