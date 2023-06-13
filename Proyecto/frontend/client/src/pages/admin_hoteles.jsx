import React, { useContext } from 'react';
import { AuthContext } from './login/auth';
import { Link } from 'react-router-dom';
import './estilo/admin_hoteles.css';

const AdminHotelesPage = () => {
  const { isLoggedAdmin } = useContext(AuthContext);
  
  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  return (
    <div className="container" onLoad={Verificacion}>
      <h1 className="titulo">Hoteles🏨</h1>
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
