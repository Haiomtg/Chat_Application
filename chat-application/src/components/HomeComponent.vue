<template>
  <div>
    <div class="flex flex-col items-center justify-center h-screen">
      <h1 class="text-4xl mb-4">Welcome to the Chat Application</h1>
      <p class="text-lg text-center max-w-md">
        This chat application allows you to communicate with others in
        real-time. You can log in, register, and start chatting with your
        friends or colleagues. Join us and experience seamless communication!
      </p>
    </div>
    <FooterComponent />
  </div>
</template>
  
  <script>
import FooterComponent from "./FooterComponent.vue";

export default {
  name: "HomeComponent",
  components: {
    FooterComponent,
  },
};
</script>
  
  <style scoped>
/* Add any additional styles here if needed */
</style>