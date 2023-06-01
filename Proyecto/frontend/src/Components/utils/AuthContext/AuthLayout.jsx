import { Box, CardMedia } from '@mui/material'
import React from 'react'
import { Outlet } from 'react-router-dom'
import { AuthContainer, Glass } from '../../../styles/StyledMUI'

const AuthLayout = () => {
  return (
    <AuthContainer>
      <Glass>
        <Box sx={{ margin: "0 auto", mb: 5 }} component="div">
          <CardMedia
            component="img"
            image="https://media.istockphoto.com/id/943914224/es/vector/icono-de-dental.jpg?s=612x612&w=0&k=20&c=GsWUIF-hIr-sc3T4PlwSPvVhIV1f2iTkN-vC-IX-FxI="
            sx={{ width: "100px", borderRadius: "8px" }}
          />
        </Box>
        <Outlet />
      </Glass>
    </AuthContainer>
  )
}

export default AuthLayout