// src/auth.js
let isAuthenticated = false;
let currentUser = '';

export const login = (username) => {
  isAuthenticated = true;
  currentUser = username; // Store the username
};

export const logout = () => {
  isAuthenticated = false;
  currentUser = ''; // Clear the username
};

export const checkAuth = () => {
  return isAuthenticated;
};

export const getCurrentUser = () => {
  return currentUser; // Return the current username
};
