<template>
  <div>
    <div class="flex flex-col items-center justify-center h-screen">
      <h2 class="text-2xl mb-4">Login</h2>
      <form @submit.prevent="handleLogin" class="w-1/3">
        <input
          type="text"
          v-model="username"
          placeholder="Username"
          class="border rounded-lg p-2 mb-2 w-full"
          required
        />
        <input
          type="password"
          v-model="password"
          placeholder="Password"
          class="border rounded-lg p-2 mb-4 w-full"
          required
        />
        <button
          type="submit"
          class="bg-blue-500 text-white p-2 rounded-lg w-full"
        >
          Login
        </button>
      </form>
      <p class="mt-4">
        Don't have an account?
        <RouterLink to="/register" class="text-blue-500">Register</RouterLink>
      </p>
    </div>
    <FooterComponent />
  </div>
</template>
  
  <script>
import { login } from "../auth"; // Import the login function
import FooterComponent from "./FooterComponent.vue";

export default {
  name: "LoginComponent",
  data() {
    return {
      username: "",
      password: "",
    };
  },
  components: {
    FooterComponent,
  },
  methods: {
    handleLogin() {
      // Simulate a successful login
      login(this.username); // Pass the username to the login function
      this.$router.push("/chat"); // Redirect to chat after login
    },
  },
};
</script>
  
  <style scoped>
/* Add any additional styles here if needed */
</style>
