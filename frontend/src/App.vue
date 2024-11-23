<template>
  <div id="app">
    <h1>{{ message }}</h1>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      message: '', // Placeholder for backend response
    };
  },
  mounted() {
    // Fetch message from backend
    axios.get('/')
      .then(response => {
        this.message = response.data.message; // Assign backend response to `message`
      })
      .catch(error => {
        console.error('Error:', error); // Log errors
      });
  },
};
</script>
