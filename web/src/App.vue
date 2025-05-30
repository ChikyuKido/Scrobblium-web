<script setup lang="ts">
import { SidebarProvider } from "@/components/ui/sidebar";
import AppSidebar from "@/components/own/AppSidebar.vue";
import { useRoute } from "vue-router";

const route = useRoute();
</script>

<template>
  <div v-if="route.path === '/login'">
    <router-view />
  </div>
  <SidebarProvider v-else>
    <div class="flex h-screen w-screen overflow-hidden">
      <AppSidebar class="w-64 shrink-0" />
      <main class="flex-1 overflow-auto">
        <router-view />
      </main>
    </div>
  </SidebarProvider>
</template>
