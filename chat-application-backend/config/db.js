// config/db.js
const mysql = require('mysql2/promise');
const dotenv = require('dotenv');

dotenv.config();

let connection;

const connectDB = async () => {
  try {
    connection = await mysql.createConnection({
      host: process.env.MYSQL_HOST,
      user: process.env.MYSQL_USER,
      password: process.env.MYSQL_PASSWORD,
      database: process.env.MYSQL_DATABASE,
    });
    console.log('MySQL connected');
  } catch (error) {
    console.error('MySQL connection error:', error);
    process.exit(1);
  }
};

const getConnection = () => {
  if (!connection) {
    throw new Error('Database connection not established');
  }
  return connection;
};

module.exports = { connectDB, getConnection };
