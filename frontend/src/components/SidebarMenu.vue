<template>
    <aside class="sidebar">
      <ul>
        <li @click="$emit('navigate', 'home')">Home</li>
        <li @click="$emit('navigate', 'popular')">Popular</li>
        <li @click="$emit('navigate', 'latest')">Latest</li>
        <li @click="$emit('navigate', 'favorites')">Favorites</li>
      </ul>
    </aside>
  </template>

  <script>
  export default {
    emits: ['navigate']
  }
  </script>
  
  <style scoped>
  .sidebar {
    width: 200px;
    background: #1e1e2f;
    color: white;
    height: 100vh;
    padding: 10px;
  }
  ul {
    list-style: none;
    padding: 0;
  }
  li {
    padding: 10px;
    cursor: pointer;
  }
  li:hover {
    background: #44475a;
  }
  </style>