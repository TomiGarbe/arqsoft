import React, { createContext, useState } from "react";

export const AuthContext = createContext({});

const AuthContextProvider = ({ children }) => {
  const [isLogged, setIsLogged] = useState(false);
  const [user, setUser] = useState(null);

  const handleLoginEmail = async (email) => {
    try {
      const response = await fetch("http://localhost:8090/cliente/email/:email", {
        method: "GET",
        body: JSON.stringify({ email }),
      });

      if (response.ok) {
        const data = await response.json();
        setIsLogged(true);
        setUser(data.user);
      } else {
        setIsLogged(false);
        setUser(null);
      }
    } catch (error) {
      setIsLogged(false);
      setUser(null);
      console.error("Error al iniciar sesión:", error);
    }
  };

  const handleLoginPassword = async (password) => {
    try {
      const response = await fetch("http://localhost:8090/cliente/password/:password", {
        method: "GET",
        body: JSON.stringify({ password }),
      });

      if (response.ok) {
        const data = await response.json();
        setIsLogged(true);
        setUser(data.user);
      } else {
        setIsLogged(false);
        setUser(null);
      }
    } catch (error) {
      setIsLogged(false);
      setUser(null);
      console.error("Error al iniciar sesión:", error);
    }
  };

  const handleLogout = () => {
    setIsLogged(false);
    setUser(null);
  };

  const propiedades = {
    isLogged,
    user,
    handleLoginEmail,
    handleLoginPassword,
    handleLogout,
  };

  return (
    <AuthContext.Provider value={propiedades}>{children}</AuthContext.Provider>
  );
};

export default AuthContextProvider;