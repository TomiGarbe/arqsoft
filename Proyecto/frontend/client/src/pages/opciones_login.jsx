import React from 'react';
import './opciones_login.css';

const RoleSelection = () => {
  return (
    <div>    
       <div className="header">
            <h2 className="titulo-header">TuHotel.com</h2>
       </div>


    <div className="selection">
      
        <div className="seleccion-container">
         
            <h2 className="titulo">Seleccione según corresponda</h2>
            <div className="buttons-container">
              <button className="button cliente" variant="contained" size="large">
                <a className="link" href="http://localhost:3000/login_cliente">Cliente</a>
              </button>
              <button className="button administrador" variant="contained" size="large">
                <a className="link" href="http://localhost:3000/login_admin">Administrador</a>
              </button>
            </div>
          
        </div>
       </div>
       <div className="footer">
        fotter
       </div>
       </div> 
       
    
  );
};

export default RoleSelection;
