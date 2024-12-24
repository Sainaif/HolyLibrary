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
<template>
  <div>
    <h1>Upload CBZ File</h1>
    <form @submit.prevent="uploadFile">
      <input type="file" @change="onFileChange" />
      <button type="submit">Upload</button>
    </form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      selectedFile: null,
    };
  },
  methods: {
    onFileChange(event) {
      this.selectedFile = event.target.files[0];
    },
    async uploadFile() {
      if (!this.selectedFile) return alert("No file selected!");

      let formData = new FormData();
      formData.append("file", this.selectedFile);

      try {
        const response = await axios.post("http://localhost:8080/upload", formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });
        alert(response.data);
      } catch (error) {
        alert("Upload failed!");
      }
    },
  },
};
</script>
