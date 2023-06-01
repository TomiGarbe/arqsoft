import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const Card = ({ name, username, id }) => {

  const navigate =  useNavigate();
  
  const favorites = JSON.parse(localStorage.getItem('favorites')) || [];

  const [isFavorite, setIsFavorite] = useState(!!favorites.find(fav => fav.id == id))

  const Fav = () => {
    setIsFavorite(true)
    favorites.push({ name, username, id })
    localStorage.setItem('favorites', JSON.stringify(favorites))
  }
  
  const selectDentist = () => {
    navigate(`/home/dentist/${id}`);
  };

  return (
    <div className="eachCard">
        <section onClick={selectDentist}>
          <img src="./../../public/images/doctor.jpg" alt="dentista"/>
        </section>
        <p>{id}</p>
        <p>{name}</p>
        <p>{username}</p>
        <button onClick={Fav} disabled={isFavorite} className="favButton">FAV</button>
    </div>
  );
};

export default Card;