import axios from 'axios';

const API_URL = 'http://localhost:5030/api';

export const sendMessage = async (token, messageData) => {
  return await axios.post(`${API_URL}/chat/sendMessage`, messageData, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};

export const getMessages = async (token) => {
  return await axios.get(`${API_URL}/chat/getMessages`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};