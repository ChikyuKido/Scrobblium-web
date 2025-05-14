<script setup lang="ts">
import {SidebarProvider, SidebarTrigger} from "@/components/ui/sidebar";
import AppSidebar from "@/components/own/AppSidebar.vue";
</script>

<template>
  <SidebarProvider>
    <div class="flex h-screen w-screen overflow-hidden">
      <AppSidebar class="w-64 shrink-0" />
      <main class="flex-1 overflow-auto">
        <router-view />
      </main>
    </div>
  </SidebarProvider>
</template>

