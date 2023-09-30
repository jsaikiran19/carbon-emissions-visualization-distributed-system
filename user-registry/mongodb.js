const mongoose = require("mongoose");

const connect = () =>
  mongoose
    .connect(process.env.MONGODB_URI || 'mongodb+srv://saikiranjella19:7cpmVPqv5xxCrjdh@cluster0.8erhxpx.mongodb.net/user-data?retryWrites=true&w=majority', {
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