const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const HotelSchema = Schema({
    _id: { type: Number, required: true },
    name: { type: String, required: true },
    city: { type: String, required: true },
    price: { type: Number, default: 0 },
    rate_number: { type: Number, default: 0 },
    rate_desc: { type: String, default: "" },
    stars: { type: Number, default: 0 },
    address: { type: String, default: "" },
    breadcrumbs: [String],
    loc: { x: Number, y: Number },
    images: [String],
    facilities: {
        general: [String],
        customer_services: [String],
        public_spaces: [String],
        shopping: [String],
        transportation: [String],
        sports: [String],
        housekeeping: [String],
        hotel_services: [String],
        activities: [String],
        misc: [String],
        entertainment: [String],
        foods_and_drinks: [String],
        business_meetings: [String]
    }
});

module.exports = mongoose.model('Hotel', HotelSchema, 'hotels');
