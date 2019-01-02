const MongoClient = require('mongodb').MongoClient;
const url = 'mongodb://localhost:27017';
const dbName = 'db';
const client = new MongoClient(url, { useNewUrlParser: true });


var express = require('express');
var router = express.Router();
var multer = require('multer');
var upload = multer();


var userUtils = require('../utils/userUtils');

router.get('/getAll', function (req, res, next) {
    console.log("get all teams hit!");
    client.connect(function (err) {
        const db = client.db(dbName);
        const collection = db.collection('teams');
        collection.find({}).toArray(function (err, data) {
            if (err)
                return res.json({ success: false, error: err });
            else
                return res.json({ success: true, data: data });
        });
    });
});

router.post('/add', upload.array(), function (req, res) {
    //only administrators can create teams
    if (!userUtils.userIsAdmin(req.session.userID)){
        return res.json({success: false, error : new Error("administrator access level required")});
    }
    //teamName will be provided by creator, teamLeader will be the id of a user, selected from a dropdown
    let newTeam = { teamName: req.body.teamName, teamLeader : req.body.teamLeader };

    client.connect(function (err) {
        const db = client.db(dbName);
        const collection = db.collection("teams");
        collection.insertOne(newTeam, function (err) {
            if (err != null) { console.log(err); return res.json({ success: false, error: err }); }
            return res.json({ success: true });
        });
    });
});


module.exports = router;
