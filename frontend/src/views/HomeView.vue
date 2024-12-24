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
  <template>
    <div>
      <h1>Uploaded Files</h1>
      <ul>
        <li v-for="file in files" :key="file.id">
          {{ file.filename }}
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    data() {
      return {
        files: [],
      };
    },
    async created() {
      const response = await axios.get("http://localhost:8080/files");
      this.files = response.data;
    },
  };
  </script>
  