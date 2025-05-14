import { h } from 'vue'
import type {ColumnDef} from "@tanstack/vue-table";
import type {TrackData} from "@/dto/TrackData.ts";
import { ArrowUpDown, ChevronDown } from 'lucide-vue-next'
import { Button } from '@/components/ui/button'

export const trackColumns: ColumnDef<TrackData>[] = [
    {
        accessorKey: 'artist',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['Artist', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => h('div', {}, row.getValue('artist')),
    },
    {
        accessorKey: 'album',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['Album', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => h('div', {}, row.getValue('album')),
    },
    {
        accessorKey: 'track',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['Track', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => h('div', {}, row.getValue('track')),
    },
    {
        accessorKey: 'listenedAt',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                class: 'text-right',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['Listened At', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => {
            const timestamp = row.getValue('listenedAt')
            const formatted = new Date(timestamp).toLocaleString()
            return h('div', { class: 'text-right' }, formatted)
        },
    },
    {
        accessorKey: 'listenCount',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                class: 'text-right',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['Listen Count', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => h('div', { class: 'text-right' }, row.getValue('listenCount')),
    },
    {
        accessorKey: 'listenTime',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                class: 'text-right',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['Listen Time (s)', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => h('div', { class: 'text-right' }, row.getValue('listenTime')),
    },
]
