<template>
  <header class="bg-blue-500 text-white p-4">
    <nav class="flex justify-between">
      <RouterLink to="/" class="text-white">Home</RouterLink>
      <RouterLink to="/chat" class="text-white">Chat</RouterLink>
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

