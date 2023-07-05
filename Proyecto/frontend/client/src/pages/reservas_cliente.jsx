import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/reservas_cliente.css';

const HomePage = () => {
  const [reservations, setReservations] = useState([]);
  const [hoteles, setHoteles] = useState([]);
  const { isLoggedCliente } = useContext(AuthContext);
  const [hotelFiltrado, setHotelFiltrado] = useState('');
  //const [startDateFilter, setStartDateFilter] = useState('');
  //const [endDateFilter, setEndDateFilter] = useState('');

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
    } else {
      window.location.href = '/';
    }
  }, [isLoggedCliente]);

  useEffect(() => {
    getReservations();
  }, [getReservations]);

  useEffect(() => {
    getHoteles();
  }, [getHoteles]);

  const handleHotelFilterChange = (event) => {
    setHotelFiltrado(event.target.value);
  };

  /*const handleStartDateFilterChange = (event) => {
    setStartDateFilter(event.target.value);
  };

  const handleEndDateFilterChange = (event) => {
    setEndDateFilter(event.target.value);
  };*/

  // Obtener una lista de nombres de hoteles únicos
  const uniqueHotelNames = Array.from(new Set(hoteles.map((hotel) => ({ id: hotel.id, nombre: hotel.nombre }))));
  const filteredReservations = reservations.filter((reservation) => {
    const hotel = hoteles.find((hotel) => hotel.id === reservation.hotel_id);
    const hotelName = hotel ? hotel.nombre : 'Hotel desconocido';
    //const startDate = new Date(reservation.anio_inicio, reservation.mes_inicio - 1, reservation.dia_inicio);
    //const endDate = new Date(reservation.anio_final, reservation.mes_final - 1, reservation.dia_final);

    // Aplicar filtros
    const hotelFilterMatch = hotelName === hotelFiltrado.selectedHotel;
    //const startDateFilterMatch = startDateFilter === '' || startDate >= new Date(startDateFilter);
    //const endDateFilterMatch = endDateFilter === '' || endDate <= new Date(endDateFilter);

    //return hotelFilterMatch && startDateFilterMatch && endDateFilterMatch;
    return hotelFilterMatch;
  });

  return (
    <div className="reservations-container1">
      <div className="reservations-container2">
        <div className="filters-container">
          <div>
            <label htmlFor="hotelFilter">Hotel:</label>
            <ul id="hotelFilter">
              <li value="" onClick={handleHotelFilterChange}>Todos los hoteles</li>
              {uniqueHotelNames.map((hotel) => (
                <li key={hotel.id} value={hotel.id} onClick={handleHotelFilterChange}>
                  {hotel.name}
                </li>
              ))}
            </ul>
          </div>
        </div>
        <h4>Datos de tus reservas:</h4>
        {filteredReservations.length ? (
          filteredReservations.map((reservation) => {
            const hotel = hoteles.find((hotel) => hotel.id === reservation.hotel_id);
            const fechaInicio = `${reservation.dia_inicio}/${reservation.mes_inicio}/${reservation.anio_inicio}`;
            const fechaFin = `${reservation.dia_final}/${reservation.mes_final}/${reservation.anio_final}`;
            return (
              <div className="reservation-card" key={reservation.ID}>
                <p>Hotel: {hotel ? hotel.nombre : 'Hotel desconocido'}</p>
                <p>Fecha de llegada: {fechaInicio}</p>
                <p>Fecha de fin: {fechaFin}</p>
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
/*
<div>
            <label htmlFor="startDateFilter">Fecha de inicio:</label>
            <input type="date" id="startDateFilter" value={startDateFilter} onChange={handleStartDateFilterChange} />
          </div>
          <div>
            <label htmlFor="endDateFilter">Fecha de fin:</label>
            <input type="date" id="endDateFilter" value={endDateFilter} onChange={handleEndDateFilterChange} />
          </div
*/
export default HomePage;










/*import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/reservas_cliente.css';

const HomePage = () => {
  const [reservations, setReservations] = useState([]);
  const [hotelesFiltrados, setHotelesFiltrados] = useState([]);
  const [hoteles, setHoteles] = useState([]);
  const { isLoggedCliente } = useContext(AuthContext);
  const [hotelFiltrado, setHotelFiltrado] = useState('');
  const [startDateFilter, setStartDateFilter] = useState('');
  const [endDateFilter, setEndDateFilter] = useState('');

  const getHoteles = useCallback(async () => {
    try {
      const hotelesArray = [];
      for (let i = 0; i < reservations.length; i++) {
        const reserva = reservations[i];
        const request = await fetch(`http://localhost:8090/cliente/hotel/${reserva.hotel_id}`);
        const response = await request.json();
        hotelesArray.push(response);
      }
      setHoteles(Array.from(new Set(hotelesArray.map((hotel) => hotel.nombre))));
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  }, [reservations]);

  const getHotelesFiltrados = useCallback(async () => {
    try {
      const hotelesArray = [];
      for (let i = 0; i < reservations.length; i++) {
        const reserva = reservations[i];
        if (reserva.hotel_id === hotelFiltrado) {
          const request = await fetch(`http://localhost:8090/cliente/hotel/${reserva.hotel_id}`);
          const response = await request.json();
          hotelesArray.push(response);
        }
      }
      setHotelesFiltrados(hotelesArray);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  }, [reservations, hotelFiltrado]);

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
    } else {
      window.location.href = '/';
    }
  }, [isLoggedCliente]);

  useEffect(() => {
    getReservations();
  }, [getReservations]);

  useEffect(() => {
    if (reservations.length > 0) {
      getHoteles();
    }
  }, [reservations, getHoteles]);
  
  useEffect(() => {
    if (reservations.length > 0) {
      getHotelesFiltrados();
    }
  }, [reservations, getHotelesFiltrados]);

  const handleHotelFiltradoChange = (event) => {
    setHotelFiltrado(event.target.value);
  };

  const handleStartDateFilterChange = (event) => {
    setStartDateFilter(event.target.value);
  };

  const handleEndDateFilterChange = (event) => {
    setEndDateFilter(event.target.value);
  };

  return (
    <div className="reservations-container1">
      <div className="reservations-container2">
        <div className="filters-container">
          <div>
            <label htmlFor="hotelFilter">Hotel:</label>
            <ul id="hotelFilter">
              <li value="" onClick={handleHotelFiltradoChange}>Todos los hoteles</li>
              {hoteles.map((hotel) => (
                <li key={hotel.id} value={hotel.id} onClick={handleHotelFiltradoChange}>
                  {hotel.nombre}
                </li>
              ))}
            </ul>
          </div>
          <div>
            <label htmlFor="startDateFilter">Fecha de inicio:</label>
            <input type="date" id="startDateFilter" value={startDateFilter} onChange={handleStartDateFilterChange} />
          </div>
          <div>
            <label htmlFor="endDateFilter">Fecha de fin:</label>
            <input type="date" id="endDateFilter" value={endDateFilter} onChange={handleEndDateFilterChange} />
          </div>
        </div>
        <h4>Datos de tus reservas:</h4>
        {reservations.length ? (
          reservations.map((reservation) => {
            const hotel = hotelesFiltrados.find((hotel) => hotel.id === reservation.hotel_id);
            const fechaInicio = `${reservation.dia_inicio}/${reservation.mes_inicio}/${reservation.anio_inicio}`;
            const fechaFin = `${reservation.dia_final}/${reservation.mes_final}/${reservation.anio_final}`;
            return (
              <div className="reservation-card" key={reservation.ID}>
                <p>Hotel: {hotel ? hotel.nombre : 'Hotel desconocido'}</p>
                <p>Fecha de llegada: {fechaInicio}</p>
                <p>Fecha de fin: {fechaFin}</p>
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

export default HomePage;*/





/*import React, { useContext, useEffect, useState, useCallback } from 'react';
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
    } else {
      window.location.href = '/';
    }
  }, [isLoggedCliente]);

  useEffect(() => {
    getReservations();
  }, [getReservations]);

  useEffect(() => {
    getHoteles();
  }, [getHoteles]); // Se agrega getHoteles como dependencia separada

  return (
    <div className="reservations-container1">
      <div className="reservations-container2">
        <h4>Datos de tus reservas:</h4>
        {reservations.length ? (
          reservations.map((reservation) => {
            const hotel = hoteles.find((hotel) => hotel.id === reservation.hotel_id);
            const fechaInicio = `${reservation.dia_inicio}/${reservation.mes_inicio}/${reservation.anio_inicio}`;
            const fechaFin = `${reservation.dia_final}/${reservation.mes_final}/${reservation.anio_final}`;
            return (
              <div className="reservation-card" key={reservation.ID}>
                <p>Hotel: {hotel ? hotel.nombre : 'Hotel desconocido'}</p>
                <p>Fecha de llegada: {fechaInicio}</p>
                <p>Fecha de fin: {fechaFin}</p>
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

export default HomePage;*/