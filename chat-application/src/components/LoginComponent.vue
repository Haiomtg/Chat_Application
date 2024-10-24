<template>
  <div>
    <div class="flex flex-col items-center justify-center h-screen">
      <h2 class="text-2xl mb-4">Login</h2>
      <form @submit.prevent="login" class="w-1/3">
        <input
          type="text"
          v-model="loginData.username"
          placeholder="Username"
          class="border rounded-lg p-2 mb-2 w-full"
          required
        />
        <input
          type="password"
          v-model="loginData.password"
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
import FooterComponent from "./FooterComponent.vue";
import { login } from "../auth";

export default {
  name: "LoginComponent",
  data() {
    return {
      loginData: {
        username: "",
        password: "",
      },
    };
  },
  components: {
    FooterComponent,
  },
  methods: {
    async login() {
      try {
        const response = await login(this.loginData);
        if (response.isAuthenticated) {
          localStorage.setItem('token', response.response.data.token);
          this.$router.push("/chat");
          alert('Login successful!');
        } else {
          alert('Login failed.');
        }
      } catch (error) {
        console.error('Login error:', error);
        alert('Login failed.');
      }
    },
  },
};
</script>

<style scoped>
/* Add any additional styles here if needed */
</style>