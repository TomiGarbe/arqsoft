import React, { useContext } from 'react';
import { AuthContext } from './login/auth';
import { Link } from 'react-router-dom';
import './estilo/opciones_admin.css';

const OpcionesAdminPage = () => {
  const { isLoggedAdmin } = useContext(AuthContext);
  const Verificacion = (hotelId) => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-cliente';
    }
    else
    {
      window.location.href = `/reservar/${hotelId}`;
    }
  };

  return (
    <div className="container" onLoad={Verificacion}>
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
