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
            <span class="font-bold">{{ message.user }}:</span>
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
import { getCurrentUser, logout } from '../auth'; // Import the logout function

export default {
  name: 'ChatComponent',
  data() {
    return {
      newMessage: '',
      messages: [],
    };
  },
  computed: {
    currentUser() {
      return getCurrentUser(); // Get the current username
    }
  },
  methods: {
    sendMessage() {
      if (this.newMessage.trim()) {
        this.messages.push({
          id: Date.now(),
          user: 'You',
          text: this.newMessage,
        });
        this.newMessage = '';
      }
    },
    logout() {
      logout(); // Call the logout function
      this.$router.push('/login'); // Redirect to login after logout
    }
  },
};
</script>

<style scoped>
/* Add any additional styles here if needed */
</style>
