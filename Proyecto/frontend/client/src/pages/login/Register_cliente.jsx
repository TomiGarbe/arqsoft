import React, { useState } from 'react';
import '../estilo/Register_cliente.css';

function RegistrationPage() {
  const [formData, setFormData] = useState({
    id: 1,
    name: "Tomi",
    last_name: "Garbe",
    username: "rdfghtbcjndxst",
    password: "q2qre2",
    email: "sdgf@123"
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const Register = () => {
    alert(JSON.stringify(formData));
    fetch(`http://localhost:5001/api/cliente`, {
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
            value={formData.username}
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
