// routes/chatRoutes.js
const express = require('express');
const { sendMessage, getMessages } = require('../controllers/chatController');
const authMiddleware = require('../middleware/authMiddleware');
const router = express.Router();

router.post('/sendMessage', authMiddleware, sendMessage);
router.get('/getMessages', authMiddleware, getMessages);

module.exports = router;

