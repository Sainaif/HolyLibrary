<template>
    <div class="manga-card" @click="$emit('select', manga)">
      <img :src="manga.thumbnail" alt="Manga Thumbnail" />
      <h3>{{ manga.title }}</h3>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      manga: Object,
    },
  };
  </script>
  
  <style scoped>
  .manga-card {
    width: 150px;
    cursor: pointer;
    padding: 10px;
    margin: 10px;
    border: 1px solid #ddd;
  }
  .manga-card img {
    width: 100%;
    height: auto;
  }
  .manga-card:hover {
    background: #f1f1f1;
  }
  </style>
  