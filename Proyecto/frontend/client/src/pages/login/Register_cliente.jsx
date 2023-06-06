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

  const handleSubmit = (e) => {
    e.preventDefault();
    // Aquí puedes agregar la lógica para enviar los datos del formulario al servidor
    console.log(formData);
  };

  return (
    <div className="registration-container">
      <h2>Registro</h2>
      <form onSubmit={handleSubmit} className="registration-form">
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
        <button type="submit">Registrarse</button>
      </form>
    </div>
  );
}

export default RegistrationPage;
