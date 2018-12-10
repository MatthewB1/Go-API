var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Hopefully we\'re going to output some data here', });
});

//return res.json({success:true, data:data});

module.exports = router;
