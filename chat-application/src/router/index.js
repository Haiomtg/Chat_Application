import { createRouter, createMemoryHistory } from 'vue-router';
import App from '../App.vue';
import ChatComponent from '../components/ChatComponent.vue';

const routes = [
  {
    path: '/',
    component: App,
  },
  {
    path: '/chat',
    name: 'Chat',
    component: ChatComponent,
  },
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

export default router;
