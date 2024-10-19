<template>
  <div>
    <div class="flex flex-col items-center justify-center h-screen">
      <h2 class="text-2xl mb-4">Register</h2>
      <form @submit.prevent="handleRegister" class="w-1/3">
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
          class="border rounded-lg p-2 mb-2 w-full"
          required
        />
        <input
          type="password"
          v-model="confirmPassword"
          placeholder="Confirm Password"
          class="border rounded-lg p-2 mb-4 w-full"
          required
        />
        <button
          type="submit"
          class="bg-blue-500 text-white p-2 rounded-lg w-full"
        >
          Register
        </button>
      </form>
      <p class="mt-4">
        Already have an account?
        <RouterLink to="/login" class="text-blue-500">Login</RouterLink>
      </p>
    </div>
    <FooterComponent />
  </div>
</template>
  
  <script>
import { RouterLink } from "vue-router";
import FooterComponent from "./FooterComponent.vue";
export default {
  name: "RegisterComponent",
  data() {
    return {
      username: "",
      password: "",
      confirmPassword: "",
    };
  },
  components: {
    RouterLink,
    FooterComponent,
  },
  methods: {
    handleRegister() {
      // Handle registration logic here
      console.log("Registering with", this.username, this.password);
    },
  },
};
</script>
  
  <style scoped>
/* Add any additional styles here if needed */
</style>