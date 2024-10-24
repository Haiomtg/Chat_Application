// server.js
const express = require('express');
const cors = require('cors');
const { connectDB } = require('./config/db'); // Import connectDB
const authRoutes = require('./routes/authRoutes');
const chatRoutes = require('./routes/chatRoutes');

const app = express();
connectDB(); // Establish the database connection

app.use(cors());
app.use(express.json());

app.use('/api/auth', authRoutes);
app.use('/api/chat', chatRoutes);

const PORT = 5030;
app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});
