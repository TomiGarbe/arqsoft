import React, { createContext, useState } from "react";

export const AuthContext = createContext({});

const AuthContextProvider = ({ children }) => {

  const [isLogged, setIsLogged] = useState(false);
  const [user, setUser] = useState(null);
  const Base_Url = "http:localhost:8090"

  const handleLogin = async (email, password) => {
    const response = await fetch(`${Base_Url}/cliente/email/${email}`);
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


/*import { createContext, useState } from "react";

export const AuthContext = createContext();

const AuthContextProvider = ({ children }) => {
  const [user, setUser] = useState();
  const navigate = useNavigate();

  const getUser = async (userId) => {
    try {
      const userResponse = await fetch(`${BASE_URL}/user/${userId}`);
      if (userResponse.ok) {
        const userData = await userResponse.json();
        return userData; // Retornar los detalles del usuario directamente
      } else {
        throw new Error("Error al obtener los detalles del usuario");
      }
    } catch (error) {
      console.error(error);
      throw new Error("Error al obtener los detalles del usuario");
    }
  };

  const getUserBookings = async (userId) => {
    try {
      const response = await fetch(`${BASE_URL}/bookings/user/${userId}`);
      if (response.ok) {
        const bookingsData = await response.json();
        return bookingsData; // Retornar el arreglo de reservas directamente
      } else {
        throw new Error("Error al obtener los datos de reservas del usuario");
      }
    } catch (error) {
      console.error(error);
      throw new Error("Error al obtener los datos de reservas del usuario");
    }
  };

  const handleLogin = async (userName, password) => {
    const response = await fetch(`${BASE_URL}/user/user_name/${userName}`);
    const data = await response.json();

    if (data.user_name === userName && data.password === password) {
      setUser(data);
      return true;
    }
    return false;
  };

  const logOut = () => {
    setUser(undefined);
    navigate("/Home");
  };

  const handleRegister = async (
    userName,
    password,
    email,
    name,
    rol,
    lastName,
    state
  ) => {
    const response = await fetch(`${BASE_URL}/user/email/${email}`);
    const data = await response.json();

    if (data.email === email || data.userName === userName) {
      return false;
    }

    const newUser = {
      user_name: userName,
      password: password,
      email: email,
      name: name,
      rol: rol,
      last_name: lastName,
      state: state,
    };

    const createUserResponse = await fetch(`${BASE_URL}/user`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newUser),
    });

    if (!createUserResponse.ok) {
      throw new Error("Error al registrar el usuario");
    }

    const createdUser = await createUserResponse.json();
    setUser(createdUser);
    setUser(undefined);
    return true;
  };

  const propiedades = {
    user,
    handleLogin,
    logOut,
    handleRegister,
    getUser,
    getUserBookings,
  };

  return (
    <AuthContext.Provider value={propiedades}>{children}</AuthContext.Provider>
  );
};

AuthContextProvider.propTypes = {
  children: PropTypes.node.isRequired,
};

export default AuthContextProvider;*/