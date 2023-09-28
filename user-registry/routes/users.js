var express = require('express');
var router = express.Router();
var auth = require('../controllers/auth');
/* GET users listing. */
router.post('/signUp', auth.signUp);
router.post('/login', auth.login);
module.exports = router;
