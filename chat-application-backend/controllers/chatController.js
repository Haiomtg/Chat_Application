// controllers/chatController.js
const Message = require('../models/Message');

exports.sendMessage = async (req, res) => {
  const { user, text } = req.body;
  const message = new Message({ user, text });
  await message.save();
  res.status(201).json(message);
};

exports.getMessages = async (req, res) => {
  const messages = await Message.find().sort({ timestamp: 1 });
  res.json(messages);
};

