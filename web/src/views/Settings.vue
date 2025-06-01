<script setup>
import {onMounted, ref, watch} from 'vue'
import { Input } from '@/components/ui/input/index.js'
import { Button } from '@/components/ui/button'
import { Switch } from '@/components/ui/switch/index.js'

const username = ref('')

const guestAccess = ref({
  homepage: false,
  list: false
})

async function syncSettings() {
  console.log('Settings changed, syncing...')
  try {
    await axiosInstance({
      method: 'post',
      url: '/settings/save',
      data: {
        username: username.value,
        homepageAccess: guestAccess.value.homepage,
        listAccess: guestAccess.value.list
      }
    })
  } catch (e) {
    console.error('Failed to sync settings', e)
  }
}


function toggleHomepage(val) {
  guestAccess.value.homepage = val
}

function toggleList(val) {
  guestAccess.value.list = val
}

watch(
    [username, () => guestAccess.value.homepage, () => guestAccess.value.list],
    syncSettings,
    { deep: true }
)

function resetPassword() {
  alert('Reset password triggered')
}

async function getSettings() {
  try {
    const res = await axiosInstance.get('/settings/get')
    const data = res.data

    username.value = data.username || ''
    guestAccess.value = {
      homepage: data.homepageAccess || false,
      list: data.listAccess || false,
    }
  } catch (e) {
    console.error('Failed to fetch settings', e)
  }
}


import {checkAuthorization} from "@/lib/utils.js";
import axiosInstance from "@/axiosInstance.js";
import PasswordReset from "@/components/own/PasswordReset.vue";
onMounted(async () => {
  await checkAuthorization('settings')
  await getSettings()
})


</script>

<template>
  <div class="w-full mx-auto p-8 space-y-6">
    <h1 class="text-2xl font-bold">Settings</h1>

    <section class="space-y-4">
      <h2 class="text-xl font-semibold">Account</h2>

      <div>
        <label class="block text-sm font-medium mb-1">Rename User</label>
        <Input v-model="username" type="text" class="w-64" />
      </div>

      <PasswordReset/>
    </section>

    <section class="space-y-4">
      <h2 class="text-xl font-semibold">Guest User Access</h2>

      <div class="flex items-center gap-2">
        <span>Homepage</span>
        <Switch :model-value="guestAccess.homepage" @update:model-value="toggleHomepage"/>
      </div>

      <div class="flex items-center gap-2">
        <span>List</span>
        <Switch :model-value="guestAccess.list" @update:model-value="toggleList"  />
      </div>
    </section>
  </div>
</template>
