import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/ver_reservas.css';

const VerReservas = () => {
  const [reservas, setReservas] = useState([]);
  const [hoteles, setHoteles] = useState([]);
  const { isLoggedAdmin } = useContext(AuthContext);

  const getHoteles = useCallback(async () => {
    try {
      const hotelesArray = [];
      for (let i = 0; i < reservas.length; i++) {
        const reserva = reservas[i];
        const request = await fetch(`http://localhost:8090/admin/hotel/${reserva.hotel_id}`);
        const response = await request.json();
        hotelesArray.push(response);
      }
      setHoteles(hotelesArray);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  }, [reservas]);

  const getReservations = useCallback(async () => {
    if (isLoggedAdmin) {
      try {
        const request = await fetch(`http://localhost:8090/admin/reservas`);
        const response = await request.json();
        setReservas(response);
        console.log(reservas);
      } catch (error) {
        console.log("No se pudieron obtener las reservas:", error);
      }
      finally {
        getHoteles();
      }
    } else {
      window.location.href = '/';
    }
  }, [isLoggedAdmin, getHoteles, reservas]);

  useEffect(() => {
    getReservations();
  }, [getReservations]);

  return (
    <body className="bodyinicio">
      <div className="containerIni">
        <div className="reserva-container">
          <h4>Datos de tus reservas:</h4>
            {reservas.length ? (
              reservas.map((reserva) => {
                const hotel = hoteles.find((hotel) => hotel.id === reserva.hotel_id);
                return (
                  <div className="reservation-card" key={reserva.ID}>
                    <p>Hotel: {hotel ? hotel.nombre : 'Hotel desconocido'}</p>
                    <p>Fecha de llegada: {reserva.fecha_inicio}</p>
                    <p>Fecha de fin: {reserva.fecha_final}</p>
                    <p>Gracias por elegirnos!</p>
                  </div>
                );
              })
            ) : (
              <p>No hay reservas</p>
            )}
        </div>
      </div>
    </body>
  );
};

export default VerReservas;