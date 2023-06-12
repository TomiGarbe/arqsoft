import LogCliente from './pages/login/login_cliente'
import LogAdmin from './pages/login/login_admin'
import Register from './pages/login/Register_cliente'
import Inicio from './pages/inicio'
import InicioAdmin from './pages/inicio_admin'
import AdminHoteles from './pages/admin_hoteles'
import AdminClientes from './pages/admin_clientes'
import Reservar from './pages/reservar'
import Cuenta from './pages/cuenta'
import ReservasCliente from './pages/reservas_cliente'
import './App.css';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import { AuthProvider } from './pages/login/auth';

function App() {
    return (
        <div>
            <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css" />
            <Router>
                <AuthProvider>
                    <Routes>
                        <Route path='/' element={<Inicio />}></Route>
                        <Route path='/admin' element={<InicioAdmin />}></Route>
                        <Route path='/login-cliente' element={<LogCliente />}></Route>
                        <Route path='/login-admin' element={<LogAdmin />}></Route>
                        <Route path='/register' element={<Register />}></Route>
                        <Route path='/administrar-hoteles' element={<AdminHoteles />}></Route>
                        <Route path='/administrar-clientes' element={<AdminClientes />}></Route>
                        <Route path='/reservar/:hotelId' element={<Reservar />}></Route>
                        <Route path='/cuenta' element={<Cuenta />}></Route>
                        <Route path='/reservas-cliente' element={<ReservasCliente />}></Route>
                    </Routes>
                </AuthProvider>
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