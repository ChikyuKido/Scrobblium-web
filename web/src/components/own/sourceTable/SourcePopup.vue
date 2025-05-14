<script setup lang="ts">
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import {Button} from "@/components/ui/button";
import {Input} from "@/components/ui/input";
</script>

<template>
  <Dialog>
    <DialogTrigger>
      <Button variant="outline">Add Source</Button>
    </DialogTrigger>
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Add Source</DialogTitle>
        <Input class="max-w-sm" placeholder="Name"/>
        <Input class="max-w-sm" placeholder=""/>
      </DialogHeader>

      <DialogFooter>
        <Button variant="outline">Save changes</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>