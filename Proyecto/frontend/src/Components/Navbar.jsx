import React from 'react'
import { Link } from 'react-router-dom'
import { useNavigate } from 'react-router-dom';

const Navbar = ({ toggleMode }) => {

  const navigate = useNavigate();

  return (
    <nav>
      <button onClick={ () => navigate(-1)}>Volver</button>
      <Link to="/home">Home</Link>
      <Link to="/favs">Favs</Link>
      <Link to="/contact">Contacto</Link>
      <button onClick={toggleMode}>Change theme</button>
    </nav>
  )
}

export default Navbar