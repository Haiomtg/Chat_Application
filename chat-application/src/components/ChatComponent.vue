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
            <span v-if="!isEditing[message.id]">
              <span class="font-bold">{{ currentUser }} :</span>
              <span>{{ message.text }}</span>
              <button @click="editMessage(message.id, message.text)" class="text-yellow-500 ml-2">Edit</button>
              <button @click="deleteMessage(message.id)" class="text-red-500 ml-2">Delete</button>
            </span>
            <div v-else>
              <input v-model="editedMessage[message.id]" :label="message.text" class="border text-black rounded-lg p-1" />
              <button @click="updateMessage(message.id, editedMessage[message.id])" class="text-green-500 ml-2">Done</button>
            </div>
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
import { logout, getCurrentUser, getUserId } from '../auth';
import { getMessages, sendMessage, deleteMessage, editMessage } from '../service/message';

export default {
  name: 'ChatComponent',
  data() {
    return {
      newMessage: '',
      messages: [],
      currentUser: getCurrentUser(),
      isEditing: {},
      editedMessage: {},
    };
  },
  created() {
    this.fetchMessages();
  },
  
  methods: {
    async fetchMessages() {
      const token = localStorage.getItem('token');
      const userId = getUserId();
      const response = await getMessages(token, userId);
      this.messages = response.data; // Assuming response.data is an array of messages
    },
    async sendMessage() {
      const token = localStorage.getItem('token');
      await sendMessage(token, { text: this.newMessage });
      this.newMessage = '';
      this.fetchMessages();
    },
    async deleteMessage(messageId) {
      console.log(messageId);
      const token = localStorage.getItem('token');
      await deleteMessage(token, messageId);
      this.fetchMessages(); // Refresh the message list after deletion
    },
    editMessage(messageId, text) {
      this.isEditing[messageId] = true;
      this.editedMessage[messageId] = text; // Store the current text for editing
    },
    async updateMessage(messageId, text) {
      const token = localStorage.getItem('token');
      await editMessage(token, messageId, text);
      this.isEditing[messageId] = false; // Exit editing mode
      this.fetchMessages(); // Refresh the message list after update
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
