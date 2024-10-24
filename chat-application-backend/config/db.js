// config/db.js
const mysql = require('mysql2/promise');

const connectDB = async () => {
  try {
    const connection = await mysql.createConnection({
      host: 'localhost',
      user: 'your_username', // replace with your MySQL username
      password: 'your_password', // replace with your MySQL password
      database: 'chat_app', // replace with your database name
    });
    console.log('MySQL connected');
    return connection;
  } catch (error) {
    console.error('MySQL connection error:', error);
    process.exit(1);
  }
};

module.exports = connectDB;
