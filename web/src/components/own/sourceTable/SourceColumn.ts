import { h } from 'vue'
import type {ColumnDef} from "@tanstack/vue-table";
import type {SourceData} from "@/dto/SourceData.ts";
import { ArrowUpDown } from 'lucide-vue-next'
import { Button } from '@/components/ui/button'

export const sourceColumns: ColumnDef<SourceData>[] = [
    {
        accessorKey: 'name',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['Name', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => h('div', {}, row.getValue('name')),
    },
    {
        accessorKey: 'provider',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['Provider', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => h('div', {}, row.getValue('provider')),
    },
    {
        accessorKey: 'apiKey',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['API Key', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => h('div', {}, row.getValue('apiKey')),
    },
    {
        accessorKey: 'scrobbles',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['Scrobbles', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => h('div', {}, row.getValue('scrobbles')),
    },
    {
        accessorKey: 'lastScrobble',
        header: ({ column }) =>
            h(Button, {
                variant: 'ghost',
                onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
            }, () => ['Last Scrobble', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })]),
        cell: ({ row }) => h('div', {}, row.getValue('lastScrobble')),
    },
]
