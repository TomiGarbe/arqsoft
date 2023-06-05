import React, { createContext, useState } from "react";

export const AuthContext = createContext({});

const AuthContextProvider = ({ children }) => {
  const [isLogged, setIsLogged] = useState(false);
  const [user, setUser] = useState(null);
  
  const Base_Url = "http:localhost:8090"

  const handleLogin = async (email, password) => {
    try {
      const response = await fetch(`${Base_Url}/cliente/email/${email}`);
      const data = await response.json();

      if (data.email === email && data.password === password) {
        setUser(data);
        setIsLogged(true);
      }
      else
      {
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
    handleLogin,
    handleLogout,
  };

  return (
    <AuthContext.Provider value={propiedades}>{children}</AuthContext.Provider>
  );
};

export default AuthContextProvider;