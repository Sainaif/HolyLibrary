import { createApp } from 'vue'
import App from './App.vue'

createApp(App).mount('#app')
import axios from 'axios';

// Set default backend URL
axios.defaults.baseURL = 'http://localhost:8080';
