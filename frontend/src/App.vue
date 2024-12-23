<template>
  <div id="app">
    <Sidebar @navigate="navigate" />
    <router-view />
  </div>
</template>

<script>
import Sidebar from './components/SidebarMenu.vue';

export default {
  components: { Sidebar },
  methods: {
    navigate(route) {
      this.$router.push(`/${route}`);
    },
  },
};
</script>

<style>
body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
  display: flex;
}
</style>
