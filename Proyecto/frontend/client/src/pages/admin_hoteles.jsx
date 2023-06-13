import React from 'react';
import { Link } from 'react-router-dom';
import './estilo/admin_hoteles.css';

const AdminHotelesPage = () => {
  return (
    <div className="container">
      <h1 className="titulo">Hoteles</h1>
      <div className="botones-container">
        <Link to="/agregar-hoteles" className="boton">
          Agregar Hoteles
        </Link>
        <Link to="/ver-hoteles" className="boton">
          Ver Hoteles
        </Link>
      </div>
    </div>
  );
};

export default AdminHotelesPage;
