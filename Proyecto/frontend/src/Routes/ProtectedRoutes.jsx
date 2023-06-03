import React, { useContext } from 'react'
import { AuthContext } from '../Components/utils/AuthContext/AuthContext';
import { Navigate, Outlet } from 'react-router'

export const ProtectedRoutes = () => {
  const { isLogged } = useContext(AuthContext);

  return isLogged ? <Outlet /> : <Navigate to="/auth/login_cliente" />;
};