import { createPinia } from 'pinia'
import { useAppStore } from '@/pinia/modules/app'
import { useUserStore } from '@/pinia/modules/user'
import { useDictionaryStore } from '@/pinia/modules/dictionary'

const pinia = createPinia()

export { useAppStore, useUserStore, useDictionaryStore }

export default pinia