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
    const userId = req.userId;
    const db = getConnection(); // Get the database connection
    const [rows] = await db.execute('SELECT id, user_id, text, timestamp FROM messages WHERE user_id = ? ORDER BY timestamp ASC', [userId]);
    res.json(rows);
  } catch (error) {
    console.error('Error retrieving messages:', error);
    res.status(500).json({ message: 'Error retrieving messages' });
  }
};

exports.deleteMessage = async (req, res) => {
  const { id } = req.params;
  console.log('Attempting to delete message with ID:', id); // Log the ID for debugging
  const db = getConnection();
  try {
    // Ensure id is a number
    const messageId = Number(id);
    
    await db.execute('DELETE FROM messages WHERE id = ?', [messageId]);
    res.status(200).json({ message: 'Message deleted' });
  } catch (error) {
    console.error('Error deleting message:', error);
    res.status(500).json({ message: 'Error deleting message' });
  }
};
