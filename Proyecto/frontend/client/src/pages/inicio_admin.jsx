import React from 'react';
import { Link } from 'react-router-dom';
import './estilo/opciones_admin.css';

const OpcionesAdminPage = () => {
  return (
    <div className="container">
      <h1 className="titulo">Opciones</h1>
      <div className="botones-container">
        <Link to="/administrar-hoteles" className="boton">
          Administrar Hoteles
        </Link>
        <Link to="/administrar-clientes" className="boton">
          Administrar Clientes
        </Link>
      </div>
    </div>
  );
};

export default OpcionesAdminPage;
