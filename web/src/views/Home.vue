<script setup>
import {ref, watch, computed, onMounted} from 'vue'
import ApexCharts from 'apexcharts'
import VueApexCharts from 'vue3-apexcharts'
import StatCard from '@/components/own/StatCard.vue'
import ChartCard from '@/components/own/ChartCard.vue'
import { useColorMode } from '@vueuse/core'

const palette = ['#4fedb9', '#50c9f2', '#6e6ef2', '#f26ec9', '#f2c94c', '#f2994a']

function getChartOptions(themeMode, type = 'default') {
  const base = {
    chart: {
      toolbar: { show: false },
      zoom: { enabled: false },
      selection: { enabled: false },
      background: 'transparent',
    },
    dataLabels: { enabled: false },
    theme: { mode: themeMode },
  }

  if (type === 'pie') {
    return {
      ...base,
      labels: ['Desktop', 'Mobile', 'Tablet', 'Web'],
      colors: palette,
    }
  }

  return {
    ...base,
    xaxis: { categories: [] },
    colors: palette,
  }
}

const timeRanges = ['Last 7d', 'Last 30d', 'Last Year', 'All Time']
const selectedRange = ref('Last 7d')

const stats = ref({
  totalScrobbles: 12560,
  averageScrobbles: 33,
  ratio: 0.95,
  listeningTime: '342h',
})

const series = ref({
  overTime: [{ name: 'Scrobbles', data: [100, 200, 180, 250, 220, 260, 300] }],
  topArtists: [{
    name: 'Scrobbles',
    data: [120, 110, 90, 85, 75],
    categories: ['Radiohead', 'Daft Punk', 'Björk', 'Kendrick Lamar', 'Sigur Rós'],
  }],
  topAlbums: [{
    name: 'Scrobbles',
    data: [90, 80, 75, 60, 55],
    categories: ['Kid A', 'Discovery', 'Vespertine', 'DAMN.', 'Ágætis byrjun'],
  }],
  hourlyPattern: [{
    name: 'Listening',
    data: Array.from({ length: 24 }, (_, i) => ({
      x: `${i}:00`, y: Math.floor(Math.random() * 100)
    }))
  }],

  sources: [
    { name: 'Spotify', data: [30, 45, 50, 60, 55, 70, 90] },
    { name: 'Apple Music', data: [20, 35, 40, 55, 50, 65, 80] },
    { name: 'YouTube', data: [10, 25, 30, 40, 45, 50, 70] }
  ],
  sourceDistribution: [3600, 3100, 2300, 1860] // matches pie labels
})
const mode = useColorMode()
const chartKey = ref(0)
const chartOptions = computed(() => ({
  default: getChartOptions(mode.value),
  pie: getChartOptions(mode.value, 'pie')
}))

watch(mode, () => {
  chartKey.value++
})

async function fetchDataForRange(range) {
  console.log('Fetching data for', range)
}
watch(selectedRange, (newRange) => {
  fetchDataForRange(newRange)
})

import {checkAuthorization} from "@/lib/utils.js";
import {useRouter} from "vue-router";

onMounted(async () => {
  const router = useRouter()
  const from = router.options.history.state.back || null

  if (from === '/login') {
    await new Promise(resolve => setTimeout(resolve, 1000))
  }

  await checkAuthorization('homepage')
})

</script>


<template>
  <div class="w-full h-full p-6 space-y-6 bg-background text-foreground">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold">Music Scrobbler Dashboard</h1>
      <div class="space-x-2">
        <button
            v-for="range in timeRanges"
            :key="range"
            @click="selectedRange = range"
            :class="['px-4 py-2 rounded-xl transition', selectedRange === range ? 'bg-primary text-primary-foreground' : 'bg-muted']"
        >
          {{ range }}
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <StatCard title="Total Scrobbles" :value="stats.totalScrobbles" />
      <StatCard title="Average Scrobbles" :value="stats.averageScrobbles" />
      <StatCard title="Ratio" :value="stats.ratio" />
      <StatCard title="Listening Time" :value="stats.listeningTime" />
    </div>

    <!-- Charts -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <ChartCard title="Scrobbles Over Time">
        <apexchart type="line" height="300" :key="chartKey" :options="chartOptions.default" :series="series.overTime" />
      </ChartCard>

      <ChartCard title="Top Artists">
        <apexchart type="bar" height="300" :key="chartKey" :options="{ ...chartOptions.default, xaxis: { categories: series.topArtists[0].categories } }" :series="[{ name: series.topArtists[0].name, data: series.topArtists[0].data }]" />
      </ChartCard>

      <ChartCard title="Top Albums">
        <apexchart type="bar" height="300" :key="chartKey" :options="{ ...chartOptions.default, xaxis: { categories: series.topAlbums[0].categories } }" :series="[{ name: series.topAlbums[0].name, data: series.topAlbums[0].data }]" />
      </ChartCard>

      <ChartCard title="Hourly Listening Pattern">
        <apexchart type="heatmap" height="300" :key="chartKey" :options="chartOptions.default" :series="series.hourlyPattern" />
      </ChartCard>

      <ChartCard title="Scrobbles by Source Over Time">
        <apexchart type="line" height="300" :key="chartKey" :options="chartOptions.default" :series="series.sources" />
      </ChartCard>

      <ChartCard title="Scrobble Distribution by Source">
        <apexchart type="pie" height="300" :key="chartKey" :options="chartOptions.pie" :series="series.sourceDistribution" />
      </ChartCard>
    </div>
  </div>
</template>

<style scoped>
</style>

