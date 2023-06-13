import React from 'react';
import { Link } from 'react-router-dom';
import './estilo/admin_clientes.css';

const AdminClientesPage = () => {
  return (
    <div className="container">
      <h1 className="titulo">Clientes</h1>
      <div className="botones-container">
      <Link to="/ver-reservas" className="boton">
          Ver Reservas
        </Link>
        <Link to="/ver-clientes" className="boton">
          Ver Clientes
        </Link>
      </div>
    </div>
  );
};

export default AdminClientesPage;
