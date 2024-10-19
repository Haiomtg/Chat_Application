<template>
  <div>
    <HeaderComponent/>
    <RouterView/>
  </div>
</template>

<script>
import HeaderComponent from './components/HeaderComponent.vue';



export default {
  name: 'App',
  components: {
    HeaderComponent
  },
  methods: {
    // goToAbout() {
    //   this.$router.push('/chat')
    // }
  }
}
</script>

<style>

</style>
