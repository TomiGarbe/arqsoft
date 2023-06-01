import React from 'react';
import { Box, Button } from '@mui/material';
import { Link } from 'react-router-dom';

const RoleSelection = () => {

  const handleRoleSelection = (role) => {
    if (role === 'cliente') {
      // Redireccionar a la página del cliente
    } else if (role === 'administrador') {
      // Redireccionar a la página del administrador
    }
  };

  return (
    <Box sx={{ display: "flex", flexDirection: "column", alignItems: "center", gap: "15px" }}>
      <Button onClick={() => handleRoleSelection("cliente")} variant="contained" size="large">
        Cliente
      </Button>
      <Button onClick={() => handleRoleSelection("administrador")} variant="contained" size="large">
        Administrador
      </Button>
      <Link to="/auth/register">¿No tienes cuenta aún?</Link>
    </Box>
  );
};

export default RoleSelection;
