<script setup lang="ts">
import {Calendar, Home, HomeIcon, Inbox, LinkIcon, ListIcon, Search, Settings, SettingsIcon} from "lucide-vue-next"
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"
import ThemeToggler from "@/components/own/ThemeToggler.vue";

// Menu items.
const items = [
  {
    title: "Home",
    url: "/",
    icon: HomeIcon,
  },
  {
    title: "List",
    url: "/recentList",
    icon: ListIcon,
  },
  {
    title: "Source",
    url: "/sources",
    icon: LinkIcon,
  },
  {
    title: "Settings",
    url: "/settings",
    icon: SettingsIcon,
  },
];
</script>

<template>
  <Sidebar>
    <SidebarContent>
      <SidebarGroup>
        <SidebarGroupLabel className="flex justify-between items-center">
          <span>Scrobblium</span>
          <ThemeToggler />
        </SidebarGroupLabel>
        <SidebarGroupContent>
          <SidebarMenu>
            <SidebarMenuItem v-for="item in items" :key="item.title">
              <SidebarMenuButton asChild>
                <router-link :to="item.url">
                  <component :is="item.icon" />
                  <span>{{item.title}}</span>
                </router-link>
              </SidebarMenuButton>
            </SidebarMenuItem>
          </SidebarMenu>
        </SidebarGroupContent>
      </SidebarGroup>
    </SidebarContent>
  </Sidebar>
</template>