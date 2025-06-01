<script setup>
import {computed, onMounted, ref, watch} from 'vue'
import StatCard from '@/components/own/StatCard.vue'
import ChartCard from '@/components/own/ChartCard.vue'
import {useColorMode} from '@vueuse/core'
import {checkAuthorization} from "@/lib/utils.js";
import {useRouter} from "vue-router";
import axiosInstance from "@/axiosInstance.js";

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
  totalScrobbles: 0,
  averageScrobbles: 0,
  ratio: 0,
  listeningTime: 0,
})
const overTimeSeries = ref([{ name: 'Scrobbles', data: [0] }])
const top5artistsSeries = ref([{
  name: 'Top 5 artists',
  data: [],
  categories: [],
}])
const top5albumsSeries = ref([{
  name: 'Top 5 albums',
  data: [],
  categories: [],
}])
const hourlySeries = ref([{
  name: 'Hourly',
  data: []
}])
const mode = useColorMode()
const chartKey = ref(0)
const chartOptions = computed(() => ({
  default: getChartOptions(mode.value),
  pie: getChartOptions(mode.value, 'pie')
}))

watch(mode, () => {
  chartKey.value++
})
function mapRange(rangeLabel) {
  switch (rangeLabel) {
    case 'Last 7d':
      return 'last7'
    case 'Last 30d':
      return 'last30'
    case 'Last Year':
      return 'lastY'
    case 'All Time':
      return 'all'
    default:
      return 'all'
  }
}
async function fetchDataForRange(range) {
  console.log('Fetching data for', range)
}
watch(selectedRange, (newRange) => {
  fetchDataForRange(newRange)
  fetchStats()
  fetchScrobbleOverTime()
  fetchTop5Artists()
  fetchTop5Albums()
  fetchHourly()
})

async function fetchStats() {
  const rangeParam = mapRange(selectedRange.value)
  try {
    const res = await axiosInstance.get(`/stats/card?range=${rangeParam}`)
    const data = res.data
    stats.value = {
      totalScrobbles: parseInt(data.totalScrobbles, 10) || 0,
      averageScrobbles: Math.floor(data.averageScrobbles) || 0,
      ratio: data.ratio ? data.ratio.toFixed(2) : '0.00',
      listeningTime: data.listeningTime
          ? `${(data.listeningTime / 3600).toFixed(0)}h`
          : '0h',
    }
  } catch (e) {
    console.error('Failed to fetch stats', e)
  }
}
async function fetchScrobbleOverTime() {
  const rangeParam = mapRange(selectedRange.value)
  try {
    const response = await axiosInstance.get(`/stats/scrobbleOverTime?range=${rangeParam}`)
    const data = response.data.data || []
    overTimeSeries.value = [{ name: 'Scrobbles', data }]
  } catch (error) {
    console.error('Failed to fetch scrobble over time', error)
  }
}
async function fetchTop5Artists() {
  const rangeParam = mapRange(selectedRange.value)
  try {
    const res = await axiosInstance.get(`/stats/top5artists?range=${rangeParam}`)
    const artists = res.data.topArtists || []

    top5artistsSeries.value[0].data = artists.map(a => a.Count || a.count || 0)
    top5artistsSeries.value[0].categories = artists.map(a => a.Name || a.name || '')
  } catch (e) {
    console.error('Failed to fetch top 5 artists', e)
  }
}
async function fetchTop5Albums() {
  const rangeParam = mapRange(selectedRange.value)
  try {
    const res = await axiosInstance.get(`/stats/top5albums?range=${rangeParam}`)
    const artists = res.data.topAlbums || []

    top5albumsSeries.value[0].data = artists.map(a => a.Count || a.count || 0)
    top5albumsSeries.value[0].categories = artists.map(a => a.Name || a.name || '')
  } catch (e) {
    console.error('Failed to fetch top 5 albums', e)
  }
}

async function fetchHourly() {
  const rangeParam = mapRange(selectedRange.value)
  try {
    const res = await axiosInstance.get(`/stats/hourly?range=${rangeParam}`)
    hourlySeries.value = [{
      name: 'Hourly',
      data: res.data.hourly.map((y, i) => ({ x: `${i}:00`, y }))
    }]
    console.log(res.data.hourly)
  } catch (e) {
    console.error('Failed to fetch top 5 albums', e)
  }
}
onMounted(async () => {
  const router = useRouter()
  const from = router.options.history.state.back || null

  if (from === '/login') {
    await new Promise(resolve => setTimeout(resolve, 1000))
  }

  await checkAuthorization('homepage')

  await fetchStats()
  await fetchScrobbleOverTime()
  await fetchTop5Artists()
  await fetchTop5Albums()
  await fetchHourly()
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
        <apexchart type="line" height="300" :key="chartKey" :options="chartOptions.default" :series="overTimeSeries" />
      </ChartCard>

      <ChartCard title="Top Artists">
        <apexchart type="bar" height="300" :key="chartKey" :options="{ ...chartOptions.default, xaxis: { categories: top5artistsSeries[0].categories } }" :series="[{ name: top5artistsSeries[0].name, data: top5artistsSeries[0].data }]" />
      </ChartCard>

      <ChartCard title="Top Albums">
        <apexchart type="bar" height="300" :key="chartKey" :options="{ ...chartOptions.default, xaxis: { categories: top5albumsSeries[0].categories } }" :series="[{ name: top5albumsSeries[0].name, data: top5albumsSeries[0].data }]" />
      </ChartCard>

      <ChartCard title="Hourly Listening Pattern">
        <apexchart type="heatmap" height="300" :key="chartKey" :options="chartOptions.default" :series="hourlySeries" />
      </ChartCard>

    </div>
  </div>
</template>

<style scoped>
</style>

