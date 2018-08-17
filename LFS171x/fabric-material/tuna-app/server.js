//SPDX-License-Identifier: Apache-2.0

// nodejs server setup 

// call the packages we need
var express       = require('express');        // call express
var session       = require('express-session');
var bodyParser    = require('body-parser');
var http          = require('http')
var fs            = require('fs');
var Fabric_Client = require('fabric-client');
var path          = require('path');
var util          = require('util');
var os            = require('os');

// instantiate the app
var app = express();

// Load all of our middleware
// configure app to use bodyParser()
// this will let us get the data from a POST
// app.use(express.static(__dirname + '/client'));
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

// this line requires and runs the code from our routes.js file and passes it app
require('./routes.js')(app);

// set up a static file server that points to the "client" directory
app.use(express.static(path.join(__dirname, './client')));

// initialize express-session to allow us track the logged-in user across sessions.
app.use(session({
  key: 'user_sid',
  secret: 'somerandonstuffs',
  resave: false,
  saveUninitialized: false,
  cookie: {
      expires: 600000
  }
}));

// middleware function to check for logged-in users
var sessionChecker = (req, res, next) => {
  if (req.session.user) {
      res.redirect('/admin');
  } else {
      next();
  }    
};

// route for Home-Page
// route for user's dashboard
app.get('/', sessionChecker, (req, res) => {
      res.redirect('/login');
});

// route for user's dashboard
app.get('/admin', (req, res) => {
  if (req.session.user) {
      res.sendFile(path.join(__dirname, './client/admin.html'));
  } else {
      res.redirect('/login');
  }
});

app.get('/customer', (req, res) => {
  if (req.session.user) {
    res.sendFile(path.join(__dirname, './client/customer.html'));
  } else {
      res.redirect('/login');
  }
});

// route for user Login
app.route('/login')
    .get(sessionChecker, (req, res) => {
        res.sendFile(path.join(__dirname, './client/login.html'));
    })
    .post((req, res) => {
        var username = req.body.username,
            password = req.body.password;
            if (username.toLowerCase() == "admin" && password.toLowerCase() == "admin") {
              req.session.user = true;
              res.redirect('/admin');
            }else if (username.toLowerCase() == "customer" && password.toLowerCase() == "customer") {
              req.session.user = true;
              res.redirect('/customer');
            }
    });

// route for user logout
app.get('/logout', (req, res) => {
  req.session.user = false;
  if (req.session.user) {
      res.redirect('/');
  } else {
      res.redirect('/login');
  }
});

// route for handling 404 requests(unavailable routes)
app.use(function (req, res, next) {
  res.status(404).send("Sorry can't find that!")
});

// Save our port
var port = process.env.PORT || 8000;

// Start the server and listen on port 
app.listen(port,function(){
  console.log("Live on port: " + port);
});

