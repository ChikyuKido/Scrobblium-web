<script setup lang="ts">

import TrackDataTable from "@/components/own/trackTable/TrackDataTable.vue";
import {trackColumns} from "@/components/own/trackTable/TrackColumns.ts";

const trackData = ref({})



import {checkAuthorization} from "@/lib/utils.js";
import {onMounted, ref, watch} from "vue";
import type {TrackData} from "@/dto/TrackData.ts";
import axiosInstance from "@/axiosInstance.ts";

async function fetchData() {
  try {
    const res = await axiosInstance.get('/track/list?combine='+selectedCombine.value)
    const data = res.data
    trackData.value = data as TrackData[]
    console.log(trackData.value)
  } catch (e) {
    console.error('Failed to fetch settings', e)
  }
}
onMounted(async () => {
  await checkAuthorization('list')
  await fetchData()
})

const selectedCombine = ref('track')

watch(selectedCombine, (val) => {
  fetchData()
})

</script>

<template>
  <div class="container py-10 mx-auto">
    <TrackDataTable :columns="trackColumns" :data="trackData" v-model="selectedCombine" />
  </div>
</template>

<style scoped>

</style>