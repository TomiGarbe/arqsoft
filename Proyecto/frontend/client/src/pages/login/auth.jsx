import React, { createContext, useState } from 'react';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [token, setToken] = useState(localStorage.getItem('token') || '');
    const [isLogged, setIsLogged] = useState(localStorage.getItem('auth') ? true : false);

  const login = (newToken) => {
    setIsLogged(true);
    setToken(newToken);
    sessionStorage.setItem('token', token);
    localStorage.setItem('auth', true);
  };

  const logout = () => {
    setIsLogged(false);
    setToken('');
    sessionStorage.removeItem('token');
    localStorage.setItem('auth', false);
  };

  const propiedades = {
    isLogged,
    login,
    logout,
  };

  return (
    <AuthContext.Provider value={propiedades}>
      {children}
    </AuthContext.Provider>
  );
};