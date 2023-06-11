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
  <body className= "bodyopciones">
    <div className="selection">
      
      <div className="contOPcio">
        <h2 className="welcome-text">Bienvenido! Elija una opción</h2>
        <div className="buttons-container">
          <button className="buttoninicio" variant="contained" size="large" onClick={handleLoginCliente}>
            Cliente
          </button>
          <button className="buttoninicio" variant="contained" size="large" onClick={handleLoginAdmin}>
            Administrador
          </button>
        </div>
      </div>
      </div>
   </body>
    
  );
};

export default RoleSelection;
