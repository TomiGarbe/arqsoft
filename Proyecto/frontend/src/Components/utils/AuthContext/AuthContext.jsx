import React, { createContext, useState } from "react";

export const AuthContext = createContext({});

const AuthContextProvider = ({ children }) => {
    
  const [isLogged, setIsLogged] = useState(
    localStorage.getItem("auth") ? true : false
  );

  const handleLogin = () => {
    setIsLogged(true);
    return {
        user: "marianapaulina@dh.com",
        password: "holamundo",
    }
  };

  const propiedades = {
    isLogged,
    handleLogin,
  };

  return (
    <AuthContext.Provider value={propiedades}>{children}</AuthContext.Provider>
  );
};

export default AuthContextProvider;



//Contectarlo a la base de datos
import React, { createContext, useState } from "react";

export const AuthContext = createContext({});

const AuthContextProvider = ({ children }) => {
  const [isLogged, setIsLogged] = useState(false);

  const handleLogin = async (username, password) => {
    try {
      // Aquí debes realizar la lógica para conectarte a la base de datos y verificar las credenciales del usuario
      // Puedes usar una librería como Axios o Fetch para realizar las peticiones a la API de tu base de datos

      // Ejemplo de una petición POST a una API
      const response = await fetch("https://api.example.com/login", {
        method: "POST",
        body: JSON.stringify({ username, password }),
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (response.ok) {
        // El inicio de sesión fue exitoso
        setIsLogged(true);
      } else {
        // El inicio de sesión falló, maneja el error según tu lógica de negocio
        console.log("Error en el inicio de sesión");
      }
    } catch (error) {
      // Maneja el error de conexión a la base de datos
      console.log("Error de conexión a la base de datos");
    }
  };

  const propiedades = {
    isLogged,
    handleLogin,
  };

  return (
    <AuthContext.Provider value={propiedades}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContextProvider;
