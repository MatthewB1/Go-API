const MongoClient = require('mongodb').MongoClient;
const url = 'mongodb://localhost:27017';
const dbName = 'db';
const client = new MongoClient(url, { useNewUrlParser: true });

var multer = require('multer');
var upload = multer();

module.exports = {

    userIsAdmin : function(userID){
        client.connect(function (err) {
            const db = client.db(dbName);
            const collection = db.collection("users");
            collection.findOne({"_id" : userID}, function (err, result) {
                if (err != null) { console.log(err); return false; }

                if (result.accessLevel == "admin"){
                    return true;
                }
                else{
                    return false;
                }
            });
        });


        return true;
    }
}