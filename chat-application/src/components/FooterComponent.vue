<template>
  <footer class="bg-gray-800 text-white p-4 text-center">
    <p>Created by: Tom Truyen</p>
    <p>&copy; {{ currentYear }}</p>
  </footer>
</template>

<script>
export default {
  name: 'FooterComponent',
  data() {
    return {
      currentYear: new Date().getFullYear() // Get the current year
    };
  }
};
</script>

<style scoped>
footer {
  position: relative;
  bottom: 0;
  width: 100%;
}
</style>

