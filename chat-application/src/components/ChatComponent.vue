<template>
  <div class="flex flex-col h-screen">
    <header class="bg-blue-500 text-white p-4">
      <h1 class="text-xl">Chat Application</h1>
    </header>
    <main class="flex-1 p-4 overflow-y-auto">
      <div class="bg-gray-100 p-4 rounded-lg shadow-md">
        <div v-for="message in messages" :key="message.id" class="mb-2">
          <span class="font-bold">{{ message.user }}:</span>
          <span>{{ message.text }}</span>
        </div>
      </div>
    </main>
    <footer class="p-4">
      <input
        v-model="newMessage"
        @keyup.enter="sendMessage"
        class="border rounded-lg p-2 w-full"
        placeholder="Type your message..."
      />
    </footer>
  </div>
</template>

<script>
export default {
  name: 'ChatComponent',
  data() {
    return {
      newMessage: '',
      messages: [],
    };
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
  },
};
</script>

<style scoped>
/* Add any additional styles here if needed */
</style>
