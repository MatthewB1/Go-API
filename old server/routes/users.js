const MongoClient = require('mongodb').MongoClient;
const url = 'mongodb://localhost:27017';
const dbName = 'db';
const client = new MongoClient(url, { useNewUrlParser: true });



var express = require('express');
var router = express.Router();
var multer = require('multer');
var upload = multer();

/* GET users listing. */
router.get('/getAll', function (req, res, next) {
  console.log("get all users hit!");
  client.connect(function (err) {
    const db = client.db(dbName);
    const collection = db.collection('users');
    collection.find({}).toArray(function (err, data) {
      if (err)
        return res.json({ success: false, error: err });
      else
        return res.json({ success: true, data: data });
    });
  });
});

router.post('/add', upload.array(), function (req, res) {
  let newUser = { username: req.body.username, password: req.body.password, accessLevel: req.body.accessLevel };

  //adding a user does not work properly

  client.connect(function (err) {
    const db = client.db(dbName);
    const collection = db.collection("users");
    collection.insertOne(newUser, function (err) {
      if (err != null) { console.log(err); return res.json({ success: false, error: err }); }
      return res.json({ success: true });
    });
  });
});

router.post('/login', upload.array(), function (req, res) {
  let user = { username: req.body.username, password: req.body.password};
  client.connect(function (err) {
    const db = client.db(dbName);
    const collection = db.collection("users");
    collection.findOne(user, function (err, result) {
      if (err != null) { console.log(err); return res.json({ success: false, error: err }); }
      
      req.session.userID = result._id;
      return res.json({ success: true, data : result, session : req.session });
    });
  });
});

module.exports = router;