import React from 'react';
import './opciones_login.css';

const RoleSelection = () => {
  return (
    <div className="selection">
      <div className="buttons-container">
        <button className="button" variant="contained" size="large">
          <a className="link" href="http://localhost:3000/login_cliente">Cliente</a>
        </button>
        <button className="button" variant="contained" size="large">
          <a className="link" href="http://localhost:3000/login_admin">Administrador</a>
        </button>
      </div>
    </div>
  );
};

export default RoleSelection;
