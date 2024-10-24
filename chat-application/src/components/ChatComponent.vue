<template>
  <div class="flex flex-col h-screen">
    <header class="bg-blue-500 text-white p-4 flex justify-between items-center">
      <h1 class="text-xl">Chat Application</h1>
      <div class="flex items-center">
        <span class="text-white">{{ currentUser }}</span>
        <button @click="logout" class="ml-4 bg-red-500 text-white p-2 rounded">Logout</button>
      </div>
    </header>
    <main class="flex-1 p-4 overflow-y-auto">
      <div class="bg-gray-100 p-4 rounded-lg shadow-md">
        <div v-for="message in messages" :key="message.id" class="mb-2 flex justify-end">
          <div class="bg-blue-500 text-white p-2 rounded-lg max-w-xs mt-2">
            <span class="font-bold">{{ currentUser }} :</span>
            <span>{{ message.text }}</span>
          </div>
        </div>
      </div>
    </main>
    <footer class="p-4 grid grid-cols-9">
      <input
        v-model="newMessage"
        @keyup.enter="sendMessage"
        class="border rounded-lg p-2 col-span-8"
        placeholder="Type your message..."
      />
      <button
        @click="sendMessage"
        class="border rounded-lg p-2 col-span-1 bg-blue-500 text-white"
      >
        Send
      </button>
    </footer>
  </div>
</template>

<script>
import { logout, getCurrentUser } from '../auth';
import { getMessages, sendMessage } from '../service/message';

export default {
  name: 'ChatComponent',
  data() {
    return {
      newMessage: '',
      messages: [],
      currentUser: getCurrentUser(),
    };
  },
  created() {
    this.fetchMessages();
  },
  
  methods: {
    async fetchMessages() {
      const token = localStorage.getItem('token');
      const response = await getMessages(token);
      this.messages = response.data; // Assuming response.data is an array of messages
    },
    async sendMessage() {
      const token = localStorage.getItem('token');
      await sendMessage(token, { text: this.newMessage });
      this.newMessage = '';
      this.fetchMessages();
    },
    logout() {
      logout();
      this.$router.push('/login');
    }
  },
};
</script>

<style scoped>
/* Add any additional styles here if needed */
</style>