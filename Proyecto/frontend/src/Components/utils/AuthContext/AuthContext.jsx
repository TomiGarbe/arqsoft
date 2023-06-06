import React, { createContext, useState } from "react";

export const AuthContext = createContext({});

const AuthContextProvider = ({ children }) => {

  const [isLogged, setIsLogged] = useState(false);
  const [user, setUser] = useState(null);
  const Base_Url = "http:localhost:8090"

  const handleLogin = async (email, password) => {
    const url = `${Base_Url}/cliente/email/${email}`;
    const options = {method: 'GET'};
    const response = await fetch(url,options);
    const data = await response.json();

    if (data.email === email && data.password === password) {
      setUser(data);
      setIsLogged(true);
    }
  };

  const propiedades = {
    isLogged,
    user,
    handleLogin,
  };

  return (
    <AuthContext.Provider value={propiedades}>{children}</AuthContext.Provider>
  );
};

export default AuthContextProvider;