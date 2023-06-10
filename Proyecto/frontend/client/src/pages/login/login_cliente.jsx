import React, { useEffect, useState } from 'react';
import '../estilo/login_cliente.css';

const ClienteLogin = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [clientData, setClientData] = useState({});

  const handleLoginCliente = () => {
    if (email === clientData.email && password === clientData.password) {
      window.location.href = 'http://localhost:3000/home';
    } else {
      alert('Credenciales incorrectas');
    }
  };

  const handleRegister = () => {
    window.location.href = 'http://localhost:3000/register';
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

  return (
    <div className="container">
      <div className="container2">
      <h1 className="title">Bienvenido Cliente</h1>
      <div className="form-container">
        <input
          type="text"
          placeholder="Correo electrónico"
          className="input"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="Contraseña"
          className="input"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <div className="button-container">
          <button className="button" onClick={handleLoginCliente}>
            Iniciar Sesión
          </button>
          <button className="button" onClick={handleRegister}>
            Registrarse
          </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ClienteLogin;
