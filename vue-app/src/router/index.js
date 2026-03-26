import { createRouter, createWebHistory } from 'vue-router'

import Body from './../components/AppBody.vue'
import Login from './../components/AppLogin.vue'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Body,
    },
    {
        path:'/login',
        name: 'Login',
        component: Login,
    }
]

const router = createRouter({history: createWebHistory(), routes})
export default router