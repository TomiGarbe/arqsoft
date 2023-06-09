import React from 'react';
import { Link } from 'react-router-dom';
import './estilo/inicio.css';

const HomePage = () => {
  return (
    <div>
      <h1 style={{ textAlign: 'center' }}>¡Hace tu Reserva!</h1>
      <div style={{ display: 'flex', justifyContent: 'center' }}>
        <div style={{ width: '200px', margin: '10px', textAlign: 'center' }}>
          <img src="ruta-imagen-hotel-1.jpg" alt="Hotel 1" />
          <p>Nombre del Hotel 1</p>
          <p>Mail 1</p>
          <p>Teléfono 1</p>
          <Link to="/reservar">
            <button>Reservar</button>
          </Link>
        </div>
        <div style={{ width: '200px', margin: '10px', textAlign: 'center' }}>
          <img src="ruta-imagen-hotel-2.jpg" alt="Hotel 2" />
          <p>Nombre del Hotel 2</p>
          <p>Mail 2</p>
          <p>Teléfono 2</p>
          <Link to="/reservar">
            <button>Reservar</button>
          </Link>
        </div>
        <div style={{ width: '200px', margin: '10px', textAlign: 'center' }}>
          <img src="ruta-imagen-hotel-3.jpg" alt="Hotel 3" />
          <p>Nombre del Hotel 3</p>
          <p>Mail 3</p>
          <p>Teléfono 3</p>
          <Link to="/reservar">
            <button>Reservar</button>
          </Link>
        </div>
        <div style={{ width: '200px', margin: '10px', textAlign: 'center' }}>
          <img src="ruta-imagen-hotel-4.jpg" alt="Hotel 4" />
          <p>Nombre del Hotel 4</p>
          <p>Mail 4</p>
          <p>Teléfono 4</p>
          <Link to="/reservar">
            <button>Reservar</button>
          </Link>
        </div>
        <div style={{ width: '200px', margin: '10px', textAlign: 'center' }}>
          <img src="ruta-imagen-hotel-5.jpg" alt="Hotel 5" />
          <p>Nombre del Hotel 5</p>
          <p>Mail 5</p>
          <p>Teléfono 5</p>
          <Link to="/reservar">
            <button>Reservar</button>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default HomePage;