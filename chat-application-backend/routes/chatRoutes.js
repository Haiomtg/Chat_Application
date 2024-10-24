// routes/chatRoutes.js
const express = require('express');
const { sendMessage, getMessages, deleteMessage, editMessage } = require('../controllers/chatController');
const authMiddleware = require('../middleware/authMiddleware');
const router = express.Router();

router.post('/sendMessage', authMiddleware, sendMessage);
router.get('/getMessages', authMiddleware, getMessages);
router.delete('/deleteMessage/:id', authMiddleware, deleteMessage);
router.put('/editMessage', authMiddleware, editMessage);
module.exports = router;
