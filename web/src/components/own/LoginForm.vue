<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

const username = ref('')
const password = ref('')
const errorMessage = ref('')
const router = useRouter()

const login = async () => {
  errorMessage.value = ''
  try {
    const response = await axios.post('/api/v1/auth/login', {
      username: username.value,
      password: password.value
    })

    const token = response.data.jwt
    localStorage.setItem('jwt', token)
    router.push('/')
  } catch (error: any) {
    if (error.response?.data?.message) {
      errorMessage.value = error.response.data.message
    } else {
      errorMessage.value = 'Login failed. Please try again.'
    }
  }
}
</script>

<template>
  <form @submit.prevent="login">
    <Card class="w-full max-w-lg mx-auto">
      <CardHeader>
        <CardTitle class="text-2xl">Login</CardTitle>
        <CardDescription>
          Enter your username below to login to your account
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div class="grid gap-4">
          <div class="grid gap-2">
            <Label for="username">Username</Label>
            <Input
                id="username"
                type="text"
                placeholder="your username"
                required
                v-model="username"
            />
          </div>
          <div class="grid gap-2">
            <Label for="password">Password</Label>
            <Input
                id="password"
                type="password"
                required
                v-model="password"
            />
          </div>
          <div v-if="errorMessage" class="text-red-600 text-sm">
            {{ errorMessage }}
          </div>
          <Button type="submit" class="w-full">
            Login
          </Button>
        </div>
      </CardContent>
    </Card>
  </form>
</template>
