import React, { useState } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import './estilo/reserva.css';

const ReservaPage = () => {
  const { hotelId } = useParams();
  const history = useHistory();
  const [fecha, setFecha] = useState('');
  const [cantidadPersonas, setCantidadPersonas] = useState('');
  const [commodities, setCommodities] = useState('');

  const handleReserva = (e) => {
    e.preventDefault();
    // Realizar lógica de reserva aquí (puede ser una llamada a una API)

    // Redirigir a la página de inicio después de la reserva
    history.push('/');
  };

  return (
    <div className="container">
      <div className="reserva-form">
        <h2>Reserva de Hotel</h2>
        <h3>Nombre del Hotel</h3>
        <img src="ruta/de/la/foto.jpg" alt="Foto del Hotel" />
        <form onSubmit={handleReserva}>
          <div className="form-group">
            <label htmlFor="fecha">Fecha de reserva:</label>
            <input
              type="date"
              id="fecha"
              value={fecha}
              onChange={(e) => setFecha(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="cantidadPersonas">Cantidad de personas:</label>
            <input
              type="number"
              id="cantidadPersonas"
              value={cantidadPersonas}
              onChange={(e) => setCantidadPersonas(e.target.value)}
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
  );
};

export default ReservaPage;
