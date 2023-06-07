import React, { useEffect, useState } from 'react';
import './login_cliente.css';

const ClienteLogin = () => {
  const [email, setEmail] = useState([{}]);
  const [password, setPassword] = useState([{}]);
  const [clientData, setClientData] = useState([{}]);

  const handleLogin = () => {
    if (email === clientData?.email && password === clientData?.password) {
      alert('Inicio de sesión exitoso');
    } else {
      alert('Credenciales incorrectas');
    }
  };

  /*const handleLogin = () => {
    if (email === clientData["email"] && password === clientData["password"]) {
      alert('Inicio de sesión exitoso');
    } else {
      alert('Credenciales incorrectas');
    }
  };*/

  useEffect(() => {
    setClientData(undefined);
    if (email) {
      fetch(`http://localhost:5001/api/cliente/email/${email}`)
        .then(response => response.json())
        .then(data => {
          setClientData(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del cliente:', error);
        });
    }
  }, [email]);

  /*useEffect(() => {
    setClientData(undefined);
    fetch(`http://localhost:5001/api/cliente/email/${email}`).then(
      response => response.json()
    ).then(
      data => {
        setClientData(data)
      }
    )
  }, [])*/

  return (
    <div className="container">
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
          <button className="button" onClick={handleLogin}>
            Iniciar Sesión
          </button>
          <button className="button">Registrarse</button>
        </div>
      </div>
    </div>
  );
};

export default ClienteLogin;
