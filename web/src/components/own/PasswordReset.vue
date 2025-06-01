<script setup  lang="ts">
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'

import { Input } from '@/components/ui/input'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import axiosInstance from '@/axiosInstance.ts'

const formSchema = toTypedSchema(z.object({
  currentPassword: z.string().min(1, 'Current password is required'),
  newPassword: z.string().min(1, 'New password must be at least 6 characters'),
  confirmPassword: z.string().min(1, 'Confirm password is required'),
}).refine(data => data.newPassword === data.confirmPassword, {
  path: ['confirmPassword'],
  message: 'Passwords do not match',
}))

async function onSubmit(values: { currentPassword: string, newPassword: string, confirmPassword: string }) {
  try {
    await axiosInstance({
      method: 'post',
      url: '/settings/password',
      data: {
        currentPassword: values.currentPassword,
        newPassword: values.newPassword
      }
    })
    console.log('Password updated')
  } catch (e) {
    console.error('Failed to update password', e)
  }
}
</script>
<template>
  <Form v-slot="{ handleSubmit }"
        :validation-schema="formSchema"
        :initial-values="{ currentPassword: '', newPassword: '', confirmPassword: '' }"
        keep-values>
    <Dialog>
      <DialogTrigger as-child>
        <Button variant="outline">
          Change Password
        </Button>
      </DialogTrigger>
      <DialogContent class="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Change Password</DialogTitle>
        </DialogHeader>

        <form
            id="passwordForm"
            @submit="handleSubmit($event, onSubmit)"
            class="space-y-4"
        >
          <FormField v-slot="{ componentField }" name="currentPassword">
            <FormItem>
              <FormLabel>Current Password</FormLabel>
              <FormControl>
                <Input type="password" placeholder="Enter current password" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="newPassword">
            <FormItem>
              <FormLabel>New Password</FormLabel>
              <FormControl>
                <Input type="password" placeholder="Enter new password" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="confirmPassword">
            <FormItem>
              <FormLabel>Confirm New Password</FormLabel>
              <FormControl>
                <Input type="password" placeholder="Confirm new password" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
        </form>

        <DialogFooter>
          <Button type="submit" form="passwordForm">
            Update Password
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </Form>
</template>
