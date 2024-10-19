// routes/chatRoutes.js
const express = require('express');
const { sendMessage, getMessages } = require('../controllers/chatController');
const authMiddleware = require('../middleware/authMiddleware');
const router = express.Router();

router.post('/messages', authMiddleware, sendMessage);
router.get('/messages', authMiddleware, getMessages);

module.exports = router;

