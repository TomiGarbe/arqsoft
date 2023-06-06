import React from 'react';
import './login_cliente.css';
  
const ClienteLogin = () => {
  return (
    <div className="container">
      <h1 className="title">Bienvenido Cliente</h1>
      <div className="form-container">
        <input type="text" placeholder="Correo electrónico" className="input" />
        <input type="password" placeholder="Contraseña" className="input" />
        <div className="button-container">
          <button className="button">Iniciar Sesión</button>
          <button className="button">Registrarse</button>
        </div>
      </div>
    </div>
  );
};

export default ClienteLogin;