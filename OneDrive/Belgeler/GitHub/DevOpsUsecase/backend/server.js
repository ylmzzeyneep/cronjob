// backend/server.js
const express = require('express');
const cors = require('cors');
const app = express();
const port = 5000;

app.use(cors());
app.use(express.json());

app.get('/', (req, res) => {
    res.json({ message: 'Hello from the backend!' });
});

app.get('/data', (req, res) => {
    const data = {
        name: 'Zeynep',
        age: 23,
        location: 'TURKIYE',
    };
    res.json(data);
});

app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});
