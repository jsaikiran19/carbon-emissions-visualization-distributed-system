const mongoose = require("mongoose");

const connect = () =>
  mongoose
    .connect(process.env.MONGODB_URI, {
      useNewUrlParser: true,
      useUnifiedTopology: true,
      // useFindAndModify: false,
    })
    .then(
      () => {
        console.log("Connected to MongoDB!");
      },
      (err) => console.log(`Something went wrong: ${err}`)
    );

module.exports.connect = connect;