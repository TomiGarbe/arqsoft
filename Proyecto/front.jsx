import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

function Home() {
  return (
    <div>
      <h2>Welcome to Hotel Reservation</h2>
      <p>Please login to make a reservation.</p>
      <Link to="/login">Login</Link>
    </div>
  );
}

function HotelList() {
  const [hotels, setHotels] = useState([]);

  useEffect(() => {
    // Obtener la lista de hoteles desde el backend al cargar el componente
    fetch('/api/hotels')
      .then((response) => response.json())
      .then((data) => setHotels(data))
      .catch((error) => console.error(error));
  }, []);

  return (
    <div>
      <h2>Available Hotels</h2>
      <ul>
        {hotels.map((hotel) => (
          <li key={hotel.id}>
            <h3>{hotel.title}</h3>
            <p>{hotel.description}</p>
            <img
              src={`/uploads/${hotel.image}`}
              alt={hotel.title}
              style={{ width: '300px', height: '200px' }}
            />
          </li>
        ))}
      </ul>
    </div>
  );
}

function HotelForm() {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [image, setImage] = useState(null);

  const handleSubmit = (e) => {
    e.preventDefault();

    const formData = new FormData();
    formData.append('title', title);
    formData.append('description', description);
    formData.append('image', image);

    fetch('/api/hotels', {
      method: 'POST',
      body: formData,
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  return (
    <div>
      <h2>Add New Hotel</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Title:</label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
        </div>
        <div>
          <label>Description:</label>
          <textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
        </div>
        <div>
          <label>Image:</label>
          <input
            type="file"
            accept="image/jpeg,image/png"
            onChange={(e) => setImage(e.target.files[0])}
          />
        </div>
        <button type="submit">Save</button>
      </form>
    </div>
  );
}

function App() {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/hotels">Hotels</Link>
            </li>
            <li>
              <Link to="/add-hotel">Add Hotel</Link>
            </li>
          </ul>
        </nav>

        <Switch>
          <Route path="/hotels">
            <HotelList />
          </Route>
          <Route path="/add-hotel">
            <HotelForm />
          </Route>
          <Route path="/">
            <Home />
          </Route>
        </Switch>
        </div>
        </Router>
);
}

export default App;
