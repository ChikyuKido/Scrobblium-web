import type { Updater } from '@tanstack/vue-table'
import type { ClassValue } from 'clsx'

import type { Ref } from 'vue'
import { clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'
import axiosInstance from "@/axiosInstance.ts";
import { useRouter } from 'vue-router'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function valueUpdater<T extends Updater<any>>(updaterOrValue: T, ref: Ref) {
  ref.value = typeof updaterOrValue === 'function'
      ? updaterOrValue(ref.value)
      : updaterOrValue
}


export async function checkAuthorization(url: string) {
  const router = useRouter();
  try {
    const response = await axiosInstance.get(`/auth/isAuthorized/${url}`);
    if (response.status !== 200) {
      await router.push('/login');
      return false;
    }
    return true;
  } catch (error) {
    await router.push('/login');
    return false;
  }
}