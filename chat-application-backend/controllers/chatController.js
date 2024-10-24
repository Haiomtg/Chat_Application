// controllers/chatController.js
const { getConnection } = require('../config/db'); // Import the getConnection function

exports.sendMessage = async (req, res) => {
  const { text } = req.body;
  const userId = req.userId; // Get userId from the auth middleware

  try {
    const db = getConnection(); // Get the database connection
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
    const db = getConnection(); // Get the database connection
    const [rows] = await db.execute('SELECT user_id, text, timestamp FROM messages ORDER BY timestamp ASC');
    res.json(rows);
    console.log(rows);
  } catch (error) {
    console.error('Error retrieving messages:', error);
    res.status(500).json({ message: 'Error retrieving messages' });
  }
};
