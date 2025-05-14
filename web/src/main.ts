import { createApp } from 'vue'
import './index.css'
import App from './App.vue'
import router from "@/router.ts";
import VueApexCharts from "vue3-apexcharts";

createApp(App).use(router).use(VueApexCharts).mount('#app')