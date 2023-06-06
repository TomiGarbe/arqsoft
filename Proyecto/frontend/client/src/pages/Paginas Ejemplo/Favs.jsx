import React from 'react';
import Card from "../Components/Card";

const Favs = () => {

  const dentists = JSON.parse(localStorage.getItem('favorites')) || [];

  return (
    <div className='favs'>
      <h1>Dentists Favs</h1>
      <div className='favs-grid'>
      <div className='card-grid'>
        {
          dentists.map(dentist => <Card key={dentist.id} name={dentist.name} username={dentist.username} id={dentist.id} />)
        }
      </div>
      </div>
    </div>
  );
};


export default Favs;