<script setup lang="ts">

import SourceDataTable from "@/components/own/sourceTable/SourceDataTable.vue";
import {sourceColumns} from "@/components/own/sourceTable/SourceColumn.ts";

const sourceData: SourceData[] = [
  {
    name: 'Spotify',
    provider: 1,
    apiKey: 'abc123xyz',
    scrobbles: 1500,
    lastScrobble: 1715674800, // Example UNIX timestamp
  },
  {
    name: 'Apple Music',
    provider: 2,
    apiKey: 'def456uvw',
    scrobbles: 980,
    lastScrobble: 1715600000,
  },
  {
    name: 'YouTube Music',
    provider: 3,
    apiKey: 'ghi789rst',
    scrobbles: 320,
    lastScrobble: 1715550000,
  },
  {
    name: 'Deezer',
    provider: 4,
    apiKey: 'jkl012mno',
    scrobbles: 450,
    lastScrobble: 1715520000,
  },
  {
    name: 'Last.fm',
    provider: 5,
    apiKey: 'mno345pqr',
    scrobbles: 2300,
    lastScrobble: 1715480000,
  },
]
import {checkAuthorization} from "@/lib/utils.js";
import {onMounted} from "vue";
onMounted(async () => {
  await checkAuthorization('sources')
})
</script>

<template>
  <div class="container py-10 mx-auto">
    <SourceDataTable :columns="sourceColumns" :data="sourceData" />
  </div>
</template>

<style scoped>

</style>