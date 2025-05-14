<script setup>
import {ref, watch} from 'vue'
import ApexCharts from 'apexcharts'
import VueApexCharts from 'vue3-apexcharts'
import StatCard from '@/components/own/StatCard.vue'
import ChartCard from '@/components/own/ChartCard.vue'
import {useColorMode} from "@vueuse/core";

function getChartOptions(themeMode) {
  return {
    chart: {
      toolbar: {
        show: false,
      },
      zoom: { enabled: false },
      selection: { enabled: false },
      background: 'transparent',
    },
    dataLabels: { enabled: false },
    xaxis: { categories: [] },
    theme: { mode: themeMode },
    colors: ['#4fedb9'],
  }
}

const timeRanges = ['Last 7d', 'Last 30d', 'Last Year', 'All Time']
const selectedRange = ref('Last 7d')

// Placeholder data
const stats = {
  totalScrobbles: 12560,
  averageScrobbles: 33,
  ratio: 0.95,
  listeningTime: '342h',
}

const mode = useColorMode()


const chartOptions = ref({})
const chartKey = ref(0)

chartOptions.value = getChartOptions(mode.value)

// Watch for mode change
watch(mode, (newMode) => {
  chartOptions.value = getChartOptions(newMode)
  chartKey.value++
})
const series2 = ref([
  {
    name: 'Line 1', // First line
    data: [30, 40, 45, 50, 49, 60, 70], // Data for the first line
  },
  {
    name: 'Line 2', // Second line
    data: [20, 35, 40, 60, 65, 70, 80], // Data for the second line
  },
]);


const series = {
  overTime: [{ name: 'Scrobbles', data: [100, 200, 180, 250, 220, 260, 300] }],
  topArtists: [{ name: 'Scrobbles', data: [120, 110, 90, 85, 75], categories: ['Radiohead', 'Daft Punk', 'Björk', 'Kendrick Lamar', 'Sigur Rós'] }],
  genres: [40, 25, 15, 10, 10],
  hourlyPattern: [
    {
      name: 'Listening',
      data: Array.from({ length: 24 }, (_, i) => ({
        x: `${i}:00`, y: Math.floor(Math.random() * 100)
      }))
    }
  ],
  topAlbums: [{ name: 'Scrobbles', data: [90, 80, 75, 60, 55], categories: ['Kid A', 'Discovery', 'Vespertine', 'DAMN.', 'Ágætis byrjun'] }],
}
</script>
<template>
  <div class="w-full h-full p-6 space-y-6 bg-background text-foreground">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold">Music Scrobbler Dashboard</h1>
      <div class="space-x-2">
        <button v-for="range in timeRanges" :key="range" @click="selectedRange = range"
                :class="['px-4 py-2 rounded-xl transition', selectedRange === range ? 'bg-primary text-primary-foreground' : 'bg-muted']">
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
        <apexchart type="line" height="300" :key="chartKey" :options="chartOptions" :series="series.overTime" />
      </ChartCard>

      <ChartCard title="Top Artists">
        <apexchart type="bar" height="300"  :key="chartKey" :options="chartOptions" :series="series.topArtists" />
      </ChartCard>

      <ChartCard title="Genres Distribution">
        <apexchart type="pie" height="300"  :key="chartKey" :options="chartOptions" :series="series.genres" />
      </ChartCard>

      <ChartCard title="Hourly Listening Pattern">
        <apexchart type="heatmap" height="300"  :key="chartKey" :options="chartOptions" :series="series.hourlyPattern" />
      </ChartCard>

      <ChartCard title="Hourly Listening Pattern">
        <apexchart type="line" height="300" :options="chartOptions" :series="series2"  :key="chartKey"  />
      </ChartCard>
    </div>
  </div>
</template>

<style scoped>
</style>
