import { loginUser, registerUser } from './service/authAPI';

let isAuthenticated = false;
let currentUser = '';
let userId = 0;
export const login = async (loginData) => {
  const response = await loginUser(loginData);
  // console.log(response.data.username);
  if (response.status === 200) {
    isAuthenticated = true;
    currentUser = response.data.username;
    userId = response.data.id;
    console.log(currentUser);
  } else {
    isAuthenticated = false;
  }
  return { isAuthenticated, response };
};

export const register = async (registerData) => {
  const response = await registerUser(registerData);
  return response;
};

export const logout = () => {
  isAuthenticated = false;
  currentUser = '';
  return isAuthenticated;
};

export const checkAuth = () => {
  const token = localStorage.getItem('token');
  return !!token;
};

export const getCurrentUser = () => {
  return currentUser;
};

export const getUserId = () => {
  return userId;
};
