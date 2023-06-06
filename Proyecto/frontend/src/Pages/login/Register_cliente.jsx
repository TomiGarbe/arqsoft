import React, { useState, Fragment } from 'react'
import { Box, Button, TextField } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { Pinwheel } from '@uiball/loaders';
import { wait } from '../../Components/utils/helper';

const Register = () => {

    const navigate = useNavigate();
    const [loading, setLoading] = useState(false);
    const [values, setValues] = useState({ name:"" });

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);

        if (values.name) {
            await wait();
            Swal.fire({
                title: '😃',
                text: `Gracias por ingresar ${values.name}! El usuario asignado es marianapaulina@dh.com y la contraseña es holamundo`,
                icon: 'success',
                showClass: {
                    popup: 'animate__animated animate__fadeInDown'
                  },
                  hideClass: {
                    popup: 'animate__animated animate__fadeOutUp'
                  }
              })
            navigate(-1)
            setLoading(false);
        } else {
            Swal.fire({
                title: '👀',
                text: 'No has ingresado los valores',
                icon: 'error',
                showClass: {
                    popup: 'animate__animated animate__fadeInDown'
                  },
                  hideClass: {
                    popup: 'animate__animated animate__fadeOutUp'
                  }
              })
            setLoading(false);
        }
    }

    return (
        <Fragment>
        <Box
            sx={{ display: "flex", flexDirection: "column", gap: "15px" }}
            component="form"
            onSubmit={handleSubmit}
        >
            <TextField
            fullWidth
            name="name"
            size="small"
            type="text"
            label="Nombre"
            placeholder="Ingrese su nombre"
            value={values.name}
            onChange={(e) => setValues({ ...values, name: e.target.value })}
            />
            <Box sx={{ textAlign: "right" }}>
            <Button disabled={loading} type="submit" size="small" variant="contained">
            {
              loading ? (
                <Pinwheel 
                size={35}
                lineWeight={3.5}
                speed={1} 
                color="black" 
                />
              ) : "🦷"
            }
          </Button>
            </Box>
        </Box>
        </Fragment>
    )
}

export default Register

/*import React, { useState } from 'react';

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
    <div>
      <h2>Registro</h2>
      <form onSubmit={handleSubmit}>
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
*/