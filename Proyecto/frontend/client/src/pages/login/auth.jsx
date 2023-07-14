import React, { createContext, useState, useEffect, useCallback } from 'react';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [isLoggedCliente, setIsLoggedCliente] = useState(localStorage.getItem('auth') === 'true');
  const [isLoggedAdmin, setIsLoggedAdmin] = useState(localStorage.getItem('auth') === 'true');

  const loginAdmin = (newToken, id) => {
    setIsLoggedAdmin(true);
    localStorage.setItem('token', newToken);
    localStorage.setItem('id_admin', id);
    localStorage.setItem('auth', true);
  };

  const loginCliente = (newToken, id) => {
    setIsLoggedCliente(true);
    localStorage.setItem('token', newToken);
    localStorage.setItem('id_cliente', id);
    localStorage.setItem('auth', true);
  };

  const logout = useCallback(async () => {
    setIsLoggedCliente(false);
    setIsLoggedAdmin(false);
    localStorage.removeItem('token');
    localStorage.removeItem('id_admin');
    localStorage.removeItem('id_cliente');
    localStorage.setItem('auth', false);
  }, []);

  //localStorage.getItem('id_cliente') === null || localStorage.getItem('id_admin') === null

  useEffect(() => {
    if (localStorage.getItem('auth') === false) {
      logout();
    }
  }, [logout]);

  const propiedades = {
    isLoggedCliente,
    isLoggedAdmin,
    loginCliente,
    loginAdmin,
    logout,
  };

  return (
    <AuthContext.Provider value={propiedades}>
      {children}
    </AuthContext.Provider>
  );
};
