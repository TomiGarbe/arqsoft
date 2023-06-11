import React, { useState, useEffect } from 'react'
import { isLoggedIn } from './auth';
import { Link } from 'react-router-dom';
import '../estilo/login_cliente.css';

const ClienteLogin = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [clientData, setClientData] = useState({});

  const handleLoginCliente = () => {
    if (email === clientData.email && password === clientData.password) {
      localStorage.setItem('Token', 'YOUR_TOKEN');
    } else {
      alert('Credenciales incorrectas');
    }
  };

  useEffect(() => {
    setClientData('');
  
    if (email) {
      fetch(`http://localhost:8090/cliente/email/${email}`)
        .then(response => response.json())
        .then(data => {
          setClientData(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del cliente:', error);
        });
    }
  }, [email]);

  if (isLoggedIn()) {
    return <Link to="/home" />;
  }

  return (
 <body className="bodylogclient">
    <div className="contLogClie1">
    <div className="contLogClien2">
      <h1 className="title">Bienvenido Cliente</h1>
       <div className="form-container">
        <input
          type="text"
          placeholder="Correo electrónico"
          className="inputLcli"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="Contraseña"
          className="inputLcli"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <div className="button-container">
          <button className="buttonClient" onClick={handleLoginCliente}>
            Iniciar Sesión
          </button>
          <Link to="/register" className="buttonClient">
          Registrarse
          </Link>
        </div>
      </div>
    </div>
    </div>
  </body>
  );
};

export default ClienteLogin;