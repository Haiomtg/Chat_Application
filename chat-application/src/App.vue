<template>
  <div>
    <nav>
      <RouterLink to="/chat" class="text-blue-500">Go to Chat</RouterLink>
    </nav>
    <RouterView/>
  </div>
</template>

<script>
import { RouterLink } from 'vue-router';

export default {
  name: 'App',
  components: {
    RouterLink
  },
  methods: {
    goToAbout() {
      this.$router.push('/chat')
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
