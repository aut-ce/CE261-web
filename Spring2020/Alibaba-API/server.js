const express = require('express')
const bodyParser = require('body-parser');
const mongoose = require('mongoose');

const hotelsRoute = require('./routes/hotels');

// connect to MongoDB
const configs = require('./configs');
mongoose.connect(configs.database, { useNewUrlParser: true, useUnifiedTopology: true }, err => {
    if (err) console.error('ERROR! ' + err);
    else console.log('connected to database');
});


const PORT = 8080;
const app = express()
app.use(bodyParser.json());

// routes
app.use('/hotels', hotelsRoute);

app.get('/', (req, res) => res.send('Hello World!'))

app.listen(PORT, () => console.log(`Server running on http://localhost:${PORT}`))
