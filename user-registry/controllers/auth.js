// var express = require("express");
var User = require("../models/user.schema");
var bcrypt = require("bcrypt");

const signUp = async (req, res) => {
  const { name, email, password } = req.body;
  const salt = await bcrypt.genSalt(10);
  const hash = await bcrypt.hash(password, salt);
  const userExists = await User.findOne({ email: email });
  if (userExists) {
    res.status(401).send({
      message: "User already exists",
    });
    return;
  }
  const user = await new User({
    name: name,
    email: email,
    password: hash,
  });
  try {
    await user.save(user);
    res.status(200).send({
      message: "User created successfully!",
    });
  } catch (error) {
    console.log("Error whille creating user");
  }

  return;
};

const login = async (req, res) => {
  const { email, password } = req.body;
  const user = await User.findOne({ email: email });
  if (!user) {
    res.status(401).send({
      message: "User does not exist. Please check email address",
    });
    return;
  }

  try {
    const valid = await bcrypt.compare(password, user.password);
    console.log(valid)
    if(valid) {
      console.log('success');
      res.status(200).send({
        message: "User login successful!",
      });
      return;
    }
    else {
      res.status(400).send({
        message: 'Incorrect password'
      })
      return;
    }
    
  } catch (err) {
    res.status(401).send({
      message: "Unexpected error",
      error: err,
    });
  }
};

module.exports = {
  signUp,
  login,
};
