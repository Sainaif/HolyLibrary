<template>
    <div class="gallery-grid">
      <MangaCard v-for="manga in mangas" :key="manga.id" :manga="manga" @select="$emit('select', manga)" />
    </div>
  </template>
  
  <script>
  import MangaCard from './MangaCard.vue';
  
  export default {
    components: { MangaCard },
    props: {
      mangas: Array,
    },
  };
  </script>
  
  <style scoped>
  .gallery-grid {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
    padding: 10px;
  }
  </style>
  
  