<script setup lang="ts">
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

import { Input } from '@/components/ui/input'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'

const formSchema = toTypedSchema(z.object({
  name: z.string().min(6, 'Name must be at least 6 characters')
                  .max(50, 'Name must be less than 50 characters'),
}))

// No functionality
function onSubmit(values: { name: string; provider: string }) {
  console.log('Form Submitted:', values)
}
</script>

<template>
  <Form v-slot="{ handleSubmit }"
        :validation-schema="formSchema"
        :initial-values="{ name: '', provider: '0' }"
        keep-values>
    <Dialog>
      <DialogTrigger as-child>
        <Button variant="outline">
          Add Source
        </Button>
      </DialogTrigger>
      <DialogContent class="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Add Source</DialogTitle>
        </DialogHeader>

        <form
            id="dialogForm"
            @submit="handleSubmit($event, onSubmit)"
            class="space-y-4"
        >
          <FormField v-slot="{ componentField }" name="name">
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input type="text" placeholder="shadcn" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="provider">
            <FormItem>
              <FormLabel>Provider</FormLabel>
              <Select v-bind="componentField">
                <FormControl>
                  <SelectTrigger class="w-full">
                    <SelectValue placeholder="Select a Provider" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent class="w-full">
                  <SelectGroup>
                    <SelectItem value="0">Scrobblium</SelectItem>
                    <SelectItem value="1">LastFM</SelectItem>
                    <SelectItem value="2">MusicBrainz</SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
              <FormMessage />
            </FormItem>
          </FormField>
        </form>

        <DialogFooter>
          <Button type="submit" form="dialogForm">
            Save changes
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </Form>
</template>