<script setup>
defineProps({
  title: String,
})
</script>

<template>
  <div class="bg-muted rounded-2xl p-4 shadow-md">
    <p class="text-md font-semibold mb-2">{{ title }}</p>
    <slot />
  </div>
</template>