import React, { useState } from 'react';
import './Register_cliente.css'

function RegistrationPage() {
  const [formData, setFormData] = useState({
    name: '',
    last_name: '',
    user_name: '',
    password: '',
    email: ''
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const Register = () => {
    fetch('http://localhost:5001/api/cliente', {
      method: 'POST',
      body: JSON.stringify(formData)
    })
      .then(response => response.json())
      .then(data => {
        console.log('Registro exitoso:', data);
        // window.location.href = 'http://localhost:3000/home';
      })
      .catch(error => {
        console.error('Error en el registro:', error);
        // alert('Credenciales incorrectas');
      });
  };

  /*casi funciona
  const Register = () => {
    fetch(`http://localhost:5001/api/cliente/${formData}`, {method: 'POST'})
    .then(response => response.json())
    .then(data => {
      console.log('Registro exitoso:', data);
      // window.location.href = 'http://localhost:3000/home';
    })
    .catch(error => {
      console.error('Error en el registro:', error);
      // alert('Credenciales incorrectas');
    });
  };*/

  /*const Register = () => {
    fetch(`http://localhost:5001/api/cliente/${formData}`)
    .then(response => response.json())
    .then(data => {
      console.log('Registro exitoso:', data);
      // window.location.href = 'http://localhost:3000/home';
    })
    .catch(error => {
      console.error('Error en el registro:', error);
      // alert('Credenciales incorrectas');
    });
  };*/

  return (
    <div className="registration-container">
      <h2>Registro</h2>
      <form onSubmit={Register} className="registration-form">
        <label>
          Nombre:
          <input
            type="text"
            name="name"
            value={formData.name}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Apellido:
          <input
            type="text"
            name="last_name"
            value={formData.last_name}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Nombre de usuario:
          <input
            type="text"
            name="user_name"
            value={formData.user_name}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Contraseña:
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Email:
          <input
            type="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <button type="submit" onClick={Register}>Registrarse</button>
      </form>
    </div>
  );
}

export default RegistrationPage;
