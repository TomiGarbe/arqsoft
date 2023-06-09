import React from 'react';
import './estilo/opciones_login.css';

const RoleSelection = () => {
  const handleLoginAdmin = () => {
    window.location.href = 'http://localhost:3000/login_admin';
  };

  const handleLoginCliente = () => {
    window.location.href = 'http://localhost:3000/login_cliente';
  };

  return (
    <div className="selection">
      <div className="buttons-container">
        <button className="button" variant="contained" size="large" onClick={handleLoginCliente}>
          Cliente
        </button>
        <button className="button" variant="contained" size="large" onClick={handleLoginAdmin}>
          Administrador
        </button>
      </div>
    </div>
  );
};

export default RoleSelection;
