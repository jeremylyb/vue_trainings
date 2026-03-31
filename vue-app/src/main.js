/*
1. Entry Point — main.js
This bootstraps the entire app. It creates a Vue app instance, plugs in the router (for navigation), 
and mounts everything onto a #app div in your HTML. Nothing renders without this.
*/



// Similar to previous simplified version where we use cdn version of vue, where we vue.createApp
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'



createApp(App).use(router).mount('#app')
