<template>
  <header class="bg-blue-500 text-white p-4">
    <nav class="flex justify-between">
      <RouterLink to="/" class="text-white">Home</RouterLink>
      <div class="flex space-x-4">
        <RouterLink to="/login" class="text-white">Login</RouterLink>
        <RouterLink to="/register" class="text-white">Register</RouterLink>
        <RouterLink to="/chat" class="text-white">Chat</RouterLink>
      </div>
    </nav>
  </header>
</template>

<script>
import { RouterLink } from 'vue-router';

export default {
  name: 'HeaderComponent',
  components: {
    RouterLink
  }
}
</script>

<style scoped>
nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
