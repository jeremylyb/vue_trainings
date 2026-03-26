
// Similar to previous simplified version where we use cdn version of vue, where we vue.createApp
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

createApp(App).use(router).mount('#app')
