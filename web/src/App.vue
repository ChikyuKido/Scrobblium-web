<script setup lang="ts">
import {SidebarProvider, SidebarTrigger} from "@/components/ui/sidebar";
import AppSidebar from "@/components/own/AppSidebar.vue";
</script>

<template>
  <SidebarProvider>
    <AppSidebar />
    <main>
      <router-view />
    </main>
  </SidebarProvider>
</template>
