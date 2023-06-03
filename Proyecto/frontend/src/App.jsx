import React, { Suspense, useState } from 'react'
import AuthContextProvider from './Components/utils/AuthContext/AuthContext';
import { BrowserRouter, Routes, Route, useNavigate, Navigate } from "react-router-dom";
import { ProtectedRoutes } from './Routes/ProtectedRoutes';
import AuthLayout from './Components/utils/AuthContext/AuthLayout';
import Login from "./Pages/login/login_cliente";
import Register from './Pages/login/Register_cliente';
import { navigation } from "./Routes/navigation";
import Navbar from "./Components/Navbar"
import Footer from "./Components/Footer"

function App() {

  const [darkMode, setDarkMode] = useState(false)

  function toggleMode() {
    setDarkMode(!darkMode)
  }

  const [values, setValues] = useState(null);

  const handleFetchValues = (userInfomation) => {
    setValues(userInfomation);
  };

  return (
    <Suspense fallback={<h1>Hola, Cargando tu página...</h1>}>
      <AuthContextProvider>
        <BrowserRouter>
          <div className={ darkMode? "dark" : "light"}>
          <Routes>
              <Route path="/auth" element={<AuthLayout />}>
              <Route path="/auth/login_cliente" element={<Login />} />
              <Route path="/auth/login/register_cliente" element={<Register />} />
            </Route>
            <Route element={<ProtectedRoutes />}>
              {
                navigation.map(({ id, path, Element }) => (
                  <Route key={id} path={path} element={
                    <>
                      <Navbar toggleMode={ toggleMode} />
                      <Element />
                      <Footer />
                    </>
                  } />
                ))
              }
            </Route>
            <Route path="/" element={<Navigate to="/auth/login_cliente" />} />
          </Routes>
          </div>
        </BrowserRouter>
      </AuthContextProvider>
    </Suspense>
  )
}

export default App
