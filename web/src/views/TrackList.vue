<script setup lang="ts">
import { shallowRef,ref, onMounted, watch } from "vue"
import TrackDataTable from "@/components/own/trackTable/TrackDataTable.vue"
import { trackColumns } from "@/components/own/trackTable/TrackColumns.ts"
import { checkAuthorization } from "@/lib/utils.js"
import axiosInstance from "@/axiosInstance.ts"
import type { TrackData } from "@/dto/TrackData.ts"

const selectedCombine = ref("track")
const trackData = shallowRef<TrackData[]>([])

async function fetchData() {
  const combine = selectedCombine.value
  if (!combine) return
  try {
    const res = await axiosInstance.get(`/track/list?combine=${combine}`)
    trackData.value = res.data as TrackData[]
  } catch (e) {
    console.error("Failed to fetch settings", e)
  }
}

onMounted(async () => {
  console.time("auth")
  await checkAuthorization("list")
  console.timeEnd("auth")
  console.time("fetch")
  await fetchData()
  console.timeEnd("fetch")
})

watch(selectedCombine, () => {
  fetchData()
})
</script>

<template>
  <div class="container py-10 mx-auto">
    <TrackDataTable
        :columns="trackColumns"
        :data="trackData"
        v-model="selectedCombine"
    />
  </div>
</template>

<style scoped>
</style>
