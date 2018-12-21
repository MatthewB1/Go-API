const MongoClient = require('mongodb').MongoClient;
const url = 'mongodb://localhost:27017';
const dbName = 'db';
const client = new MongoClient(url);


var express = require('express');
var router = express.Router();
var bodyParser = require('body-parser');
var multer = require('multer');
var upload = multer();

/* GET users listing. */
router.get('/users/getAll', function (req, res, next) {
  console.log("getAll hit!");
  client.connect(function (err) {
    const db = client.db(dbName);
    const collection = db.collection('users');
    collection.find({}).toArray(function (err, data) {
      if (err)
        return res.status(false).json(err);
      else
        return res.status(true).json(data);
    });
  });
});

router.post('/users/addUser', upload.array(), function (req, res) {
  let newUser = { username: req.body.username, password: req.body.password };

  client.connect(function (err) {
    const db = client.db(dbName);
    const collection = db.collection("users");
    collection.insertOne(newUser, function (err, result) {
      if (err != null) { console.log(err); return res.json({ success: false, error: err }); }
      return res.json({ success: true });
    });
  });
});

module.exports = router;
