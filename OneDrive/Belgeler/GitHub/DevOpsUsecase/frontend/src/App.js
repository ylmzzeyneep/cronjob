// frontend/src/App.js
import React, { useEffect, useState } from 'react';

function App() {
  const [message, setMessage] = useState('');
  const [userData, setUserData] = useState(null);

  // Fetch the welcome message from the backend
  useEffect(() => {
    fetch('http://localhost:5001')
      .then((response) => response.json())
      .then((data) => setMessage(data.message))
      .catch((error) => console.error('Error:', error));
  }, []);

  // Fetch additional user data from the backend
  useEffect(() => {
    fetch('http://localhost:5001/data')
      .then((response) => response.json())
      .then((data) => setUserData(data))
      .catch((error) => console.error('Error:', error));
  }, []);

  return (
    <div className="App">
      <h1>{message}</h1>
      {userData && (
        <div>
          <h2>User Info </h2>
          <p>Name: {userData.name}</p>
          <p>Age: {userData.age}</p>
          <p>Location: {userData.location}</p>
        </div>
      )}
    </div>
  );
}

export default App;
