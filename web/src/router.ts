import { createRouter, createWebHistory } from 'vue-router'
import Home from "@/views/Home.vue";
import Settings from "@/views/Settings.vue";
import TrackList from "@/views/TrackList.vue";
import Source from "@/views/Source.vue";
import Login from "@/views/Login.vue";

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/settings', name: 'Settings', component: Settings },
    { path: '/list', name: 'List', component: TrackList },
    { path: '/sources', name: 'Source', component: Source },
    { path: '/login', name: 'Login', component: Login },
    // add more routes here
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
