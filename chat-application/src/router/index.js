import { createRouter, createMemoryHistory } from 'vue-router';
import ChatComponent from '../components/ChatComponent.vue';
import LoginComponent from '../components/LoginComponent.vue';
import RegisterComponent from '../components/RegisterComponent.vue';
import HomeComponent from '../components/HomeComponent.vue';
import { checkAuth } from '../auth';

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
      checkAuth() ? next() : next('/login');
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