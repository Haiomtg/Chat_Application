<template>
  <div>
    <div class="flex flex-col items-center justify-center h-screen">
      <h2 class="text-2xl mb-4">Register</h2>
      <form @submit.prevent="register" class="w-1/3">
        <input
          type="text"
          v-model="registerData.username"
          placeholder="Username"
          class="border rounded-lg p-2 mb-2 w-full"
          required
        />
        <input
          type="password"
          v-model="registerData.password"
          placeholder="Password"
          class="border rounded-lg p-2 mb-2 w-full"
          required
        />
        <input
          type="password"
          v-model="registerData.confirmPassword"
          placeholder="Confirm Password"
          class="border rounded-lg p-2 mb-4 w-full"
          required
        />
        <input
          type="email"
          v-model="registerData.email"
          placeholder="Email"
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
import FooterComponent from "./FooterComponent.vue";
import { register } from "../auth";

export default {
  name: "RegisterComponent",
  data() {
    return {
      registerData: {
        username: "",
        password: "",
        confirmPassword: "",
        email: "",
      },
    };
  },
  components: {
    FooterComponent,
  },
  methods: {
    async register() {
      try {
        const response = await register(this.registerData);
        if (response) {
          alert('User registered successfully!');
          this.$router.push("/login");
        } else {
          alert('Registration failed.');
        }
      } catch (error) {
        console.error('Registration error:', error);
        alert('Registration failed.');
      }
    },
  },
};
</script>

<style scoped>
/* Add any additional styles here if needed */
</style>