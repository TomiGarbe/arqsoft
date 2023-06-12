import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './auth';
import '../estilo/login_admin.css';

const AdminLogin = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [adminData, setAdminData] = useState({});
  const { login } = useContext(AuthContext);

  const handleLoginAdmin = () => {
    if (email === adminData.email && password === adminData.password) {
      const token = 'TOKEN_ADMIN';
      login(token);
      window.location.href = '/admin';
    } else {
      alert('Credenciales incorrectas');
    }
  };

  useEffect(() => {
    setAdminData('');
  
    if (email) {
      fetch(`http://localhost:8090/admin/email/${email}`)
        .then(response => response.json())
        .then(data => {
          setAdminData(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del cliente:', error);
        });
    }
  }, [email]);

  return (
    <div className="container">
      <div className="container2">
      <h1 className="title">Bienvenido Administrador</h1>
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
          <button className="button" onClick={handleLoginAdmin}>
            Iniciar Sesión
          </button>
        </div>
      </div>
      </div>
    </div>
  );
};

export default AdminLogin;
