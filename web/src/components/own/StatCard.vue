<script setup>
defineProps({
  title: String,
  value: [String, Number],
})
</script>

<template>
  <div class="bg-muted rounded-2xl p-4 shadow-md">
    <p class="text-sm text-muted-foreground">{{ title }}</p>
    <p class="text-xl font-semibold mt-1">{{ value }}</p>
  </div>
</template>
