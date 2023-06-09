import OpLogin from './pages/opciones_login'
import LogCliente from './pages/login/login_cliente'
import LogAdmin from './pages/login/login_admin'
import Register from './pages/login/Register_cliente'
import Inicio from './pages/inicio'
import './App.css';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
//import { useEffect, useState } from 'react';

function App() {
    return (
        <div>
            <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css" />
            <Router>
                <Routes>
                    <Route path='/' element={<OpLogin />}></Route>
                    <Route path='/login_cliente/' element={<LogCliente />}></Route>
                    <Route path='/login_admin/' element={<LogAdmin />}></Route>
                    <Route path='/register/' element={<Register />}></Route>
                    <Route path='/home/' element={<Inicio />}></Route>
                </Routes>
            </Router>
        </div>
    )
}

export default App;


/*
<h1>Cliente</h1>

                <p>{typeof clientData === 'undefined' ? (<> ACA VA LA INFO DEL CLIENTE... CARGANDO </>) :
                (
                    <>
                    <h4>Nombre: {clientData["name"]} </h4>
                    <h4>Last Name: {clientData["last_name"]}</h4>
                    <h4>PWD: {clientData["password"]}</h4>
                    </>
                ) }</p>
*/