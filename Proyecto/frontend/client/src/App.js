import Item from './pages/item'

import './App.css';

import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import { useEffect, useState } from 'react';

function App() {
    const [clientData, setClientData] = useState([{}]);

    useEffect(() => {
        setClientData(undefined);
        fetch('http://localhost:5001/api/clients/1').then(
            response => response.json()
        ).then(
            data => {
                setClientData(data)
            }
        )
    }, [])

    return (
        <div>
            <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css" />
            
            <h1>Cliente</h1>

            <p>{typeof clientData === 'undefined' ? (<> ACA VA LA INFO DEL CLIENTE... CARGANDO </>) :
            (
                <>
                <h4>Nombre: {clientData["name"]} </h4>
                <h4>Last Name: {clientData["last_name"]}</h4>
                <h4>PWD: {clientData["password"]}</h4>
                </>
            ) }</p>



        </div>
    )
}

export default App;