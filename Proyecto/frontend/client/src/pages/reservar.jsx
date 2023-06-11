import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import './estilo/reservar.css';

const ReservaPage = () => {
  const { hotelId } = useParams();
  const [hotelData, setHotelData] = useState('');
  const [fecha, setFecha] = useState('');
  const [cantidadPersonas, setCantidadPersonas] = useState('');
  const [commodities, setCommodities] = useState('');
  
  const handleReserva = (e) => {
    e.preventDefault();
    
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

  return (
    <div>
      {typeof hotelData === 'undefined' ? (<> CARGANDO... </>) :
      (
        <div className="container">
          <div className="reserva-form">
            <h2>Reserva de Hotel</h2>
            <h3>{hotelData["nombre"]}</h3>
            <img src="ruta/de/la/foto.jpg" alt="Foto del Hotel"/>
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
      )}
    </div>
  );
};

export default ReservaPage;

/*<p>{typeof clientData === 'undefined' ? (<> ACA VA LA INFO DEL CLIENTE... CARGANDO </>) :
                (
                    <>
                    <h4>Nombre: {clientData["name"]} </h4>
                    <h4>Last Name: {clientData["last_name"]}</h4>
                    <h4>PWD: {clientData["password"]}</h4>
                    </>
                ) }</p>*/
