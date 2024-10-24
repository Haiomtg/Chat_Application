// controllers/chatController.js
const db = require('../config/db'); // Import the database connection

exports.sendMessage = async (req, res) => {
  const { text } = req.body;
  const userId = req.userId; // Get userId from the auth middleware

  try {
    const [result] = await db.execute(
      'INSERT INTO messages (user_id, text) VALUES (?, ?)',
      [userId, text]
    );
    res.status(201).json({ id: result.insertId, userId, text });
  } catch (error) {
    console.error('Error saving message:', error);
    res.status(500).json({ message: 'Error saving message' });
  }
};

exports.getMessages = async (req, res) => {
  try {
    const [rows] = await db.execute('SELECT user_id, text, timestamp FROM messages ORDER BY timestamp ASC');
    res.json(rows);
  } catch (error) {
    console.error('Error retrieving messages:', error);
    res.status(500).json({ message: 'Error retrieving messages' });
  }
};
