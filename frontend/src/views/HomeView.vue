<template>
    <div>
      <Header @search="onSearch" />
      <GalleryGrid :mangas="filteredMangas" @select="viewDetails" />
    </div>
  </template>
  
  <script>
  import Header from '../components/HeaderMenu.vue';
  import GalleryGrid from '../components/GalleryGrid.vue';
  
  export default {
    components: { Header, GalleryGrid },
    data() {
      return {
        mangas: [],
        searchQuery: '',
      };
    },
    computed: {
      filteredMangas() {
        return this.mangas.filter(manga =>
          manga.title.toLowerCase().includes(this.searchQuery.toLowerCase())
        );
      },
    },
    methods: {
      async fetchMangas() {
        const response = await fetch(`${import.meta.env.VITE_API_URL}/api/files`);
        this.mangas = await response.json();
      },
      onSearch(query) {
        this.searchQuery = query;
      },
      viewDetails(manga) {
        this.$router.push(`/details/${manga.id}`);
      },
    },
    mounted() {
      this.fetchMangas();
    },
  };
  </script>
  