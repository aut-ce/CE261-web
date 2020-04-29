const express = require('express');
const router = express.Router();

const Hotel = require('../models/hotels');


router.get('/:city', (req, res) => {
    Hotel.find({ city: req.params.city }, '_id name stars price images rate_number rate_desc', (err, hotels) => {
        if (err) return console.error(err);
        if (hotels) {
            const result = []
            hotels.forEach(h => result.push({
                '_id': h._id,
                'name': h.name,
                'stars': h.stars,
                'price': h.price,
                'image': h.images[0],
                'rate_number': h.rate_number,
                'rate_desc': h.rate_desc
            }))
            res.json(result);
        }
        else res.sendStatus(404);
    });
});

router.get('/hotel/:id', (req, res) => {
    Hotel.findOne({ _id: req.params.id }, '-__v -rate_number -rate_desc -price', (err, hotel) => {
        if (err) return console.error(err);
        if (hotel) res.send(hotel);
        else res.sendStatus(404);
    })
});

router.post('/import', (req, res) => {
    let hotel = new Hotel(req.body);
    hotel.save((err, h) => {
        if (err) {
            console.error(err);
            res.sendStatus(400);
        }
        else res.sendStatus(200);
    });
});

module.exports = router;
