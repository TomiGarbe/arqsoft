import React from 'react';
import '../estilo/login_admin.css';

const AdminLogin = () => {
  return (
    <div className="container">
      <h1 className="title">Bienvenido Administrador</h1>
      <div className="form-container">
        <input type="text" placeholder="Correo electrónico" className="input" />
        <input type="password" placeholder="Contraseña" className="input" />
        <div className="button-container">
          <button className="button">Iniciar Sesión</button>
        </div>
      </div>
    </div>
  );
};

export default AdminLogin;
