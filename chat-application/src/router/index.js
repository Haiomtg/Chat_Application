import { createRouter, createMemoryHistory } from 'vue-router';
// import App from '../App.vue';
import ChatComponent from '../components/ChatComponent.vue';
import LoginComponent from '../components/LoginComponent.vue';
import RegisterComponent from '../components/RegisterComponent.vue';
import HomeComponent from '../components/HomeComponent.vue';
import { checkAuth } from '../auth'; // Import the checkAuth function

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeComponent,
  },
  {
    path: '/chat',
    name: 'Chat',
    component: ChatComponent,
    beforeEnter: (to, from, next) => {
      if (checkAuth()) {
        next(); // Allow access to the chat component
      } else {
        next('/login'); // Redirect to login if not authenticated
      }
    },
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginComponent,
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterComponent,
  },
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

export default router;
