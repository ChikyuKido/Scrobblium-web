<script setup lang="ts">

import TrackDataTable from "@/components/own/trackTable/TrackDataTable.vue";
import {trackColumns} from "@/components/own/trackTable/TrackColumns.ts";

const sampleTrackData: TrackData[] = [
  {
    artist: 'Radiohead',
    album: 'In Rainbows',
    track: 'Nude',
    listenedAt: 1689435600000,
    listenCount: 12,
    listenTime: 247,
  },
  {
    artist: 'Daft Punk',
    album: 'Discovery',
    track: 'Harder, Better, Faster, Stronger',
    listenedAt: 1689522000000,
    listenCount: 20,
    listenTime: 224,
  },
  {
    artist: 'Nirvana',
    album: 'Nevermind',
    track: 'Smells Like Teen Spirit',
    listenedAt: 1689608400000,
    listenCount: 18,
    listenTime: 301,
  },
  {
    artist: 'Kendrick Lamar',
    album: 'DAMN.',
    track: 'HUMBLE.',
    listenedAt: 1689694800000,
    listenCount: 25,
    listenTime: 177,
  },
  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },
  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },  {
    artist: 'Billie Eilish',
    album: 'WHEN WE ALL FALL ASLEEP, WHERE DO WE GO?',
    track: 'bad guy',
    listenedAt: 1689781200000,
    listenCount: 30,
    listenTime: 194,
  },
]

import {checkAuthorization} from "@/lib/utils.js";
import {onMounted} from "vue";
onMounted(async () => {
  await checkAuthorization('list')
})

</script>

<template>
  <div class="container py-10 mx-auto">
    <TrackDataTable :columns="trackColumns" :data="sampleTrackData" />
  </div>
</template>

<style scoped>

</style>