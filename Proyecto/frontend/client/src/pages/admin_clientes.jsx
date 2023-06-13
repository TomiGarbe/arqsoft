import React, { useContext } from 'react';
import { AuthContext } from './login/auth';
import { Link } from 'react-router-dom';
import './estilo/admin_clientes.css';

const AdminClientesPage = () => {
  const { isLoggedAdmin } = useContext(AuthContext);
  
  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  return (
    <div className="container" onLoad={Verificacion}>
      <h1 className="titulo">Clientes👥</h1>
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
