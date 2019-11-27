var mongoose = require('mongoose');
mongoose.Promise = require('bluebird');
var passportLocalMongoose = require('passport-local-mongoose');

var Schema = mongoose.Schema;

var Users = new Schema({
	type : {
		enum : ['S', 'T'],
		default : 'S'
	}
},{
    timestamps: true
});

Users.plugin(passportLocalMongoose);

module.exports = mongoose.model('Users',Users);