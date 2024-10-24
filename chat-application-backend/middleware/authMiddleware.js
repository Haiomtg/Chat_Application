// middleware/authMiddleware.js
const jwt = require('jsonwebtoken');

module.exports = (req, res, next) => {
  const token = req.headers['authorization'];
  if (!token) return res.status(403).send('Access denied.');

  jwt.verify(token, 'your_jwt_secret', (err, decoded) => {
    if (err) return res.status(403).send('Invalid token.');
    req.userId = decoded.id; // Set userId from the token
    next();
  });
};

