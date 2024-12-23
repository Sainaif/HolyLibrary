<template>
  <nav class="header">
    <div class="logo">HolyLibrary</div>
    <input type="text" placeholder="Search..." v-model="searchQuery" />
    <button @click="$emit('search', searchQuery)">Search</button>
  </nav>
</template>

<script>
export default {
  data() {
    return {
      searchQuery: '',
    };
  },
};
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background: #2d2d2d;
  color: white;
}
input {
  padding: 5px;
  margin-right: 10px;
}
</style>
