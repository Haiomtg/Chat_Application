// models/Message.js
class Message {
  constructor(user, text) {
    this.user = user;
    this.text = text;
    this.timestamp = new Date();
  }
}

module.exports = Message;
