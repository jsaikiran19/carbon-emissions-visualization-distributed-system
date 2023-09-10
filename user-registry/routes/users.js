var express = require('express');
var router = express.Router();
var auth = require('../controllers/auth');
/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('respond with a resource');
});
router.post('/signUp', auth.signUp);
router.post('/login', auth.login);
module.exports = router;
