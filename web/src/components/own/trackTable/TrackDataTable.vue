<script setup lang="ts" generic="TData, TValue">
import {ArrowLeft,ArrowRight} from "lucide-vue-next"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Input } from '@/components/ui/input'
import {Button} from '@/components/ui/button'

import {
  FlexRender,
  getCoreRowModel,
  useVueTable,
  type ColumnFiltersState,
  type SortingState,
  getPaginationRowModel,
  getSortedRowModel,
  getFilteredRowModel,
  type ColumnDef,
} from '@tanstack/vue-table'
import { valueUpdater } from '@/lib/utils'
import {ref, watch} from "vue";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

const props = defineProps<{
  columns: ColumnDef<TData, TValue>[]
  data: TData[]
  modelValue: string
}>()

const emit = defineEmits(['update:modelValue'])

const selected = ref(props.modelValue || '')

watch(selected, (val) => {
  emit('update:modelValue', val)
})

const sorting = ref<SortingState>([])
const columnFilters = ref<ColumnFiltersState>([])


const table = useVueTable({
  get data() { return props.data },
  get columns() {
    return props.columns
  },
  getCoreRowModel: getCoreRowModel(),
  getPaginationRowModel: getPaginationRowModel(),
  getSortedRowModel: getSortedRowModel(),
  onSortingChange: updaterOrValue => valueUpdater(updaterOrValue, sorting),
  onColumnFiltersChange: updaterOrValue => valueUpdater(updaterOrValue, columnFilters),
  getFilteredRowModel: getFilteredRowModel(),

  state: {
    get sorting() { return sorting.value },
    get columnFilters() { return columnFilters.value },
    pagination: {
      pageIndex: 0,
      pageSize: 18
    }
  },
})
</script>

<template>
  <div class="flex items-center justify-between py-4">
    <Input class="max-w-sm" placeholder="Filter tracks..."
           :model-value="table.getColumn('artist')?.getFilterValue() as string"
           @update:model-value="table.getColumn('artist')?.setFilterValue($event)" />

    <Select v-model="selected">
      <SelectTrigger>
        <SelectValue placeholder="Combined by" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectItem value="artist">
            Artist
          </SelectItem>
          <SelectItem value="album">
            Album
          </SelectItem>
          <SelectItem value="track">
            Track
          </SelectItem>
          <SelectItem value="nothing">
            Nothing
          </SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  </div>
  <div class="border rounded-md">
    <Table>
      <TableHeader>
        <TableRow v-for="headerGroup in table.getHeaderGroups()" :key="headerGroup.id">
          <TableHead v-for="header in headerGroup.headers" :key="header.id">
            <FlexRender
                v-if="!header.isPlaceholder" :render="header.column.columnDef.header"
                :props="header.getContext()"
            />
          </TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <template v-if="table.getRowModel().rows?.length">
          <TableRow
              v-for="row in table.getRowModel().rows" :key="row.id"
              :data-state="row.getIsSelected() ? 'selected' : undefined"
          >
            <TableCell v-for="cell in row.getVisibleCells()" :key="cell.id">
              <FlexRender :render="cell.column.columnDef.cell" :props="cell.getContext()" />
            </TableCell>
          </TableRow>
        </template>
        <template v-else>
          <TableRow>
            <TableCell :colspan="columns.length" class="h-24 text-center">
              No results.
            </TableCell>
          </TableRow>
        </template>
      </TableBody>
    </Table>
  </div>
  <div class="flex items-center justify-center py-4 space-x-2">
    <Button
        variant="outline"
        size="sm"
        :disabled="!table.getCanPreviousPage()"
        @click="table.previousPage()"
    >
      <component :is="ArrowLeft" />
    </Button>
    <div class="flex items-center justify-center text-sm font-medium">
      {{ table.getState().pagination.pageIndex + 1 }} of
      {{ table.getPageCount() }}
    </div>
    <Button
        variant="outline"
        size="sm"
        :disabled="!table.getCanNextPage()"
        @click="table.nextPage()"
    >
      <component :is="ArrowRight" />
    </Button>
  </div>
</template>