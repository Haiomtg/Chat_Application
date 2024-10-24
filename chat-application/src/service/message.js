import axios from 'axios';

const API_URL = 'http://localhost:5030/api';

export const sendMessage = async (token, messageData) => {
  return await axios.post(`${API_URL}/chat/sendMessage`, messageData, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};

export const getMessages = async (token, userId) => {
  return await axios.get(`${API_URL}/chat/getMessages`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
    body: {
      userId: userId,
    },
  });
};

export const deleteMessage = async (token, messageId) => {
  return await axios.delete(`${API_URL}/chat/deleteMessage/${messageId}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};
